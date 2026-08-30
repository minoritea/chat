// Package session provides utilities for user sessions.
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

// User is an alias for database.User.
type User = database.User

// SessionName is the name of the session cookie.
const SessionName = "session"

// StoreContainer provides a sessions.Store.
type StoreContainer interface {
	SessionStore() sessions.Store
}

// QuerierContainer provides a database.Querier.
type QuerierContainer interface {
	Querier() database.Querier
}

// PerpetuateSession creates a session record for the user that expires in one year.
func PerpetuateSession(ctx context.Context, c QuerierContainer, userID string) (database.Session, error) {
	return c.Querier().CreateSession(ctx, database.CreateSessionParams{
		ID:        database.NewID(),
		UserID:    userID,
		ExpiresAt: time.Now().AddDate(0, 0, 30),
	})
}

// GetSessionID returns the session ID stored in the session.
func GetSessionID(session *sessions.Session) (string, bool) {
	sessionID, ok := session.Values["session_id"].(string)
	return sessionID, ok
}

// SetSessionID stores the session ID in the session.
func SetSessionID(session *sessions.Session, sessionID string) {
	session.Values["session_id"] = sessionID
}

// ErrSessionNotFound indicates the session was not found.
var ErrSessionNotFound = errors.New("session not found")

// GetUserFromSessionContainer provides dependencies for GetUserFromSession.
type GetUserFromSessionContainer interface {
	StoreContainer
	QuerierContainer
}

// GetUserFromSession returns the user associated with the session ID in the session.
func GetUserFromSession(ctx context.Context, c QuerierContainer, session *sessions.Session) (*User, error) {
	sessionID, ok := GetSessionID(session)
	if !ok {
		return nil, errors.Join(ErrSessionNotFound, errors.New("session id not found"))
	}
	q := c.Querier()
	user, err := q.GetUserBySessionID(ctx, sessionID)
	if err != nil && database.IsRecordNotFound(err) {
		return nil, errors.Join(ErrSessionNotFound, err)
	}
	return &user, err
}

// Get returns the session for the request.
func Get(c StoreContainer, r *http.Request) (*sessions.Session, error) {
	session, err := c.SessionStore().Get(r, SessionName)
	if err != nil {
		return nil, errors.Join(ErrSessionNotFound, err)
	}
	return session, nil
}

// MustGet is like Get but panics on error.
func MustGet(c StoreContainer, r *http.Request) *sessions.Session {
	session, err := Get(c, r)
	if err != nil {
		log.Panic(err)
	}
	return session
}

// New returns a session for the request.
func New(c StoreContainer, r *http.Request) (*sessions.Session, error) {
	return c.SessionStore().New(r, SessionName)
}

// MustNew is like New but panics on error.
func MustNew(c StoreContainer, r *http.Request) *sessions.Session {
	session, err := New(c, r)
	if err != nil {
		log.Panic(err)
	}
	return session
}

// Flash represents a flash message.
type Flash struct {
	Message string
	Type    string
}

// NewErrorFlash creates an error flash message.
func NewErrorFlash(message string) Flash {
	return Flash{Message: message, Type: "error"}
}

// GetFlashes returns the flash messages in the session.
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

// FlashData holds flash messages for template rendering.
type FlashData struct{ Flashes []Flash }

// RedirectWithErrorFlash adds an error flash message to the session and redirects to url.
func RedirectWithErrorFlash(w http.ResponseWriter, r *http.Request, session *sessions.Session, url, message string) {
	session.AddFlash(NewErrorFlash(message))
	MustSave(session, r, w)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// MustSave saves the session to the response. It panics on error.
func MustSave(session *sessions.Session, r *http.Request, w http.ResponseWriter) {
	err := session.Save(r, w)
	if err != nil {
		panic(err)
	}
}
