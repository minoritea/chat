// Package resource provides the application's resource container.
package resource

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/template"
)

// Container holds the application's resources and implements the container interfaces.
type Container struct {
	db           *sql.DB
	renderer     *template.Renderer
	sessionStore sessions.Store
	config       config.Config
}

// New returns a new Container.
func New(conf config.Config) (*Container, error) {
	db, err := sql.Open(conf.DatabaseDriver, conf.DatabasePath+"?_loc=UTC&_journal=WAL&_timeout=5000")
	if err != nil {
		return nil, err
	}
	renderer, err := template.NewRenderer()
	if err != nil {
		return nil, err
	}
	store := sessions.NewCookieStore([]byte(conf.SessionSecret))
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteLaxMode
	return &Container{config: conf, db: db, renderer: renderer, sessionStore: store}, nil
}

// Querier returns a database.Querier.
func (c Container) Querier() database.Querier { return database.New(c.db) }

// Renderer returns the template.Renderer.
func (c Container) Renderer() *template.Renderer { return c.renderer }

// SessionStore returns the sessions.Store.
func (c Container) SessionStore() sessions.Store { return c.sessionStore }

// Config returns the config.Config.
func (c Container) Config() config.Config { return c.config }

// DB returns the *sql.DB.
func (c Container) DB() *sql.DB { return c.db }
