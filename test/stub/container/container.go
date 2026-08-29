// Package container provides stub implementations of container interfaces.
package container

import (
	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/database"
)

// QuerierContainer provides a database.Querier.
type QuerierContainer interface {
	Querier() database.Querier
}

// SessionStoreContainer provides a sessions.Store.
type SessionStoreContainer interface {
	SessionStore() sessions.Store
}

type querierContainer struct{ q database.Querier }

func (c querierContainer) Querier() database.Querier { return c.q }

// NewQuerierContainer returns a QuerierContainer stub wrapping q.
func NewQuerierContainer(q database.Querier) QuerierContainer {
	return querierContainer{q}
}

type sessionStoreContainer struct{ s sessions.Store }

func (c sessionStoreContainer) SessionStore() sessions.Store { return c.s }

// NewSessionStoreContainer returns a SessionStoreContainer stub wrapping s.
func NewSessionStoreContainer(s sessions.Store) SessionStoreContainer {
	return sessionStoreContainer{s}
}
