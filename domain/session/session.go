package session

import (
	"errors"
	"net/http"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/database"
)

type Container = container.Container
type User = database.User

func StoreNewSession(c *Container, w http.ResponseWriter, r *http.Request, userID string) error {
	q := c.GetQuerier()
	sessionID, err := q.CreateSession(userID)
	if err != nil {
		return err
	}
	store, err := c.GetSessionStore().New(r, "session")
	if err != nil {
		return err
	}
	store.Values["session_id"] = sessionID
	return store.Save(r, w)
}

var SessionNotFound = errors.New("session not found")

func GetUserFromSession(c *Container, r *http.Request) (*User, error) {
	store, err := c.GetSessionStore().Get(r, "session")
	if err != nil {
		return nil, errors.Join(SessionNotFound, err)
	}
	sessionID, ok := store.Values["session_id"].(string)
	if !ok {
		return nil, SessionNotFound
	}
	q := c.GetQuerier()
	user, err := q.GetUserBySessionID(sessionID)
	if err != nil && database.IsSqliteNotFound(err) {
		return nil, errors.Join(SessionNotFound, err)
	}
	return user, err
}
