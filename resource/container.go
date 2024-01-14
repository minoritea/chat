package resource

import (
	"database/sql"

	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/template"
)

type Container struct {
	db           *sql.DB
	renderer     *template.Renderer
	sessionStore sessions.Store
	config       config.Config
}

func New(conf config.Config) (*Container, error) {
	db, err := sql.Open(conf.DatabaseDriver, conf.DatabasePath)
	if err != nil {
		return nil, err
	}
	renderer, err := template.NewRenderer()
	if err != nil {
		return nil, err
	}
	store := sessions.NewCookieStore([]byte(conf.SessionSecret))
	return &Container{config: conf, db: db, renderer: renderer, sessionStore: store}, nil
}

func (c Container) Querier() database.Querier    { return database.New(c.db) }
func (c Container) Renderer() *template.Renderer { return c.renderer }
func (c Container) SessionStore() sessions.Store { return c.sessionStore }
func (c Container) Config() config.Config        { return c.config }
func (c Container) DB() *sql.DB                  { return c.db }
