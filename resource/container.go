package resource

import (
	"database/sql"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
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
	db, err := sql.Open("sqlite3", conf.DatabasePath)
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

func (c Container) Queries() *database.Queries   { return database.New(c.db) }
func (c Container) Renderer() *template.Renderer { return c.renderer }
func (c Container) SessionStore() sessions.Store { return c.sessionStore }
func (c Container) Config() config.Config        { return c.config }
