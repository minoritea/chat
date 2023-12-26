package resource

import (
	"database/sql"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/template"
)

type Container struct {
	queries      *database.Queries
	renderer     *template.Renderer
	sessionStore sessions.Store
}

func New() (*Container, error) {
	db, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		return nil, err
	}
	queries := database.New(db)
	renderer, err := template.NewRenderer()
	if err != nil {
		return nil, err
	}
	store := sessions.NewCookieStore(sessionSecretFromEnv())
	return &Container{queries: queries, renderer: renderer, sessionStore: store}, nil
}

func sessionSecretFromEnv() []byte {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "development secret"
	}
	return []byte(secret)
}

func (c Container) Queries() *database.Queries   { return c.queries }
func (c Container) Renderer() *template.Renderer { return c.renderer }
func (c Container) SessionStore() sessions.Store { return c.sessionStore }
