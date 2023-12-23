package container

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/templates"
)

type Container struct {
	querier          *database.Querier
	templateRenderer *templates.Renderer
	sessionStore     sessions.Store
}

func New() (*Container, error) {
	querier, err := database.New("./chat.db")
	if err != nil {
		return nil, err
	}
	renderer, err := templates.NewRenderer()
	if err != nil {
		return nil, err
	}
	store := sessions.NewCookieStore([]byte("secret"))
	return &Container{querier: querier, templateRenderer: renderer, sessionStore: store}, nil
}

func sessionSecretFromEnv() string {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "development secret"
	}
	return secret
}

func (c *Container) GetQuerier() *database.Querier            { return c.querier }
func (c *Container) GetTemplateRenderer() *templates.Renderer { return c.templateRenderer }
func (c *Container) GetSessionStore() sessions.Store          { return c.sessionStore }
