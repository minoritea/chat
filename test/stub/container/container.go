package container

import (
	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/database"
)

type QuerierContainer interface {
	Querier() database.Querier
}

type SessionStoreContainer interface {
	SessionStore() sessions.Store
}

type querierContainer struct{ q database.Querier }

func (c querierContainer) Querier() database.Querier { return c.q }

func NewQuerierContainer(q database.Querier) QuerierContainer {
	return querierContainer{q}
}

type sessionStoreContainer struct{ s sessions.Store }

func (c sessionStoreContainer) SessionStore() sessions.Store { return c.s }

func NewSessionStoreContainer(s sessions.Store) SessionStoreContainer {
	return sessionStoreContainer{s}
}
