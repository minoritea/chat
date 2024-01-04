package session

import (
	"context"
	"encoding/gob"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/database"
)

func init() { gob.Register(Flash{}) }

type User = database.User

const SessionName = "session"

type SessionStoreContainer interface {
	SessionStore() sessions.Store
}

type QuerierContainer interface {
	Querier() database.Querier
}

type StoreNewSessionContainer interface {
	SessionStoreContainer
	QuerierContainer
}

func StoreNewSession(ctx context.Context, c StoreNewSessionContainer, w http.ResponseWriter, r *http.Request, userID string) error {
	q := c.Querier()
	session, err := q.CreateSession(ctx, database.CreateSessionParams{
		ID:     database.NewID(),
		UserID: userID,
	})
	if err != nil {
		return err
	}
	store, err := c.SessionStore().New(r, SessionName)
	if err != nil {
		return err
	}
	store.Values["session_id"] = session.ID
	return store.Save(r, w)
}

var SessionNotFound = errors.New("session not found")

type GetUserFromSessionContainer interface {
	SessionStoreContainer
	QuerierContainer
}

func GetUserFromSession(ctx context.Context, c GetUserFromSessionContainer, r *http.Request) (*User, error) {
	store, err := c.SessionStore().Get(r, SessionName)
	if err != nil {
		return nil, errors.Join(SessionNotFound, err)
	}
	sessionID, ok := store.Values["session_id"].(string)
	if !ok {
		return nil, SessionNotFound
	}
	q := c.Querier()
	user, err := q.GetUserBySessionID(ctx, sessionID)
	if err != nil && database.IsRecordNotFound(err) {
		return nil, errors.Join(SessionNotFound, err)
	}
	return &user, err
}

func MustGet(c SessionStoreContainer, r *http.Request) *sessions.Session {
	session, err := c.SessionStore().Get(r, SessionName)
	if err != nil {
		log.Panic(err)
	}
	return session
}

type Flash struct {
	Message string
	Type    string
}

func NewErrorFlash(message string) Flash {
	return Flash{Message: message, Type: "error"}
}

func AddFlash(c SessionStoreContainer, w http.ResponseWriter, r *http.Request, flash Flash) error {
	session, err := c.SessionStore().Get(r, SessionName)
	if err != nil {
		return err
	}
	session.AddFlash(flash)
	return session.Save(r, w)
}

func MustAddFlash(c SessionStoreContainer, w http.ResponseWriter, r *http.Request, flash Flash) {
	err := AddFlash(c, w, r, flash)
	if err != nil {
		log.Panic(err)
	}
}

func GetFlashes(c SessionStoreContainer, w http.ResponseWriter, r *http.Request) ([]Flash, error) {
	session, err := c.SessionStore().Get(r, SessionName)
	if err != nil {
		return nil, err
	}
	flashes := session.Flashes()
	if len(flashes) == 0 {
		return nil, nil
	}
	err = session.Save(r, w)
	if err != nil {
		return nil, err
	}
	messages := make([]Flash, len(flashes))
	for i, f := range flashes {
		messages[i] = f.(Flash)
	}
	return messages, nil
}

func MustGetFlashes(c SessionStoreContainer, w http.ResponseWriter, r *http.Request) []Flash {
	flashes, err := GetFlashes(c, w, r)
	if err != nil {
		log.Panic(err)
	}
	return flashes
}

type FlashData struct{ Flashes []Flash }

func RedirectWithErrorFlash(c SessionStoreContainer, w http.ResponseWriter, r *http.Request, url, message string) {
	MustAddFlash(c, w, r, NewErrorFlash(message))
	http.Redirect(w, r, url, http.StatusSeeOther)
}
