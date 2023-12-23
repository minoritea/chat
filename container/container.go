package container

import (
	"database/sql"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/templates"
)

type Container struct {
	queries          *database.Queries
	templateRenderer *templates.Renderer
	sessionStore     sessions.Store
}

func New() (*Container, error) {
	db, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		return nil, err
	}
	queries := database.New(db)
	renderer, err := templates.NewRenderer()
	if err != nil {
		return nil, err
	}
	store := sessions.NewCookieStore([]byte("secret"))
	return &Container{queries: queries, templateRenderer: renderer, sessionStore: store}, nil
}

func sessionSecretFromEnv() string {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "development secret"
	}
	return secret
}

func (c *Container) GetQueries() *database.Queries            { return c.queries }
func (c *Container) GetTemplateRenderer() *templates.Renderer { return c.templateRenderer }
func (c *Container) GetSessionStore() sessions.Store          { return c.sessionStore }
