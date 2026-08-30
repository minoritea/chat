package auth

import (
	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/template"
)

// Container is a service locator.
type Container interface {
	Querier() database.Querier
	Renderer() *template.Renderer
	SessionStore() sessions.Store
	Config() config.Config
}
