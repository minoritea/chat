package session

import (
	"context"
	"encoding/gob"
	"errors"
	"log"
	"net/http"
	"time"

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

func PerpetuateSession(ctx context.Context, c QuerierContainer, userID string) (database.Session, error) {
	return c.Querier().CreateSession(ctx, database.CreateSessionParams{
		ID:        database.NewID(),
		UserID:    userID,
		ExpiresAt: time.Now().AddDate(1, 0, 0),
	})
}

func GetSessionID(session *sessions.Session) (string, bool) {
	sessionID, ok := session.Values["session_id"].(string)
	return sessionID, ok
}

func SetSessionID(session *sessions.Session, sessionID string) {
	session.Values["session_id"] = sessionID
}

var SessionNotFound = errors.New("session not found")

type GetUserFromSessionContainer interface {
	SessionStoreContainer
	QuerierContainer
}

func GetUserFromSession(ctx context.Context, c QuerierContainer, session *sessions.Session) (*User, error) {
	sessionID, ok := GetSessionID(session)
	if !ok {
		return nil, errors.Join(SessionNotFound, errors.New("session id not found"))
	}
	q := c.Querier()
	user, err := q.GetUserBySessionID(ctx, sessionID)
	if err != nil && database.IsRecordNotFound(err) {
		return nil, errors.Join(SessionNotFound, err)
	}
	return &user, err
}

func Get(c SessionStoreContainer, r *http.Request) (*sessions.Session, error) {
	session, err := c.SessionStore().Get(r, SessionName)
	if err != nil {
		return nil, errors.Join(SessionNotFound, err)
	}
	return session, nil
}

func MustGet(c SessionStoreContainer, r *http.Request) *sessions.Session {
	session, err := Get(c, r)
	if err != nil {
		log.Panic(err)
	}
	return session
}

func New(c SessionStoreContainer, r *http.Request) (*sessions.Session, error) {
	return c.SessionStore().New(r, SessionName)
}

func MustNew(c SessionStoreContainer, r *http.Request) *sessions.Session {
	session, err := New(c, r)
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

func GetFlashes(session *sessions.Session) []Flash {
	flashes := session.Flashes()
	if len(flashes) == 0 {
		return nil
	}
	messages := make([]Flash, len(flashes))
	for i, f := range flashes {
		messages[i] = f.(Flash)
	}
	return messages
}

type FlashData struct{ Flashes []Flash }

func RedirectWithErrorFlash(w http.ResponseWriter, r *http.Request, session *sessions.Session, url, message string) {
	session.AddFlash(NewErrorFlash(message))
	MustSave(session, r, w)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func MustSave(session *sessions.Session, r *http.Request, w http.ResponseWriter) {
	err := session.Save(r, w)
	if err != nil {
		panic(err)
	}
}
