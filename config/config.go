package config

import (
	"flag"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Config struct {
	Host               string
	Port               string
	GithubClientID     string
	GithubClientSecret string
	SessionSecret      string
	DatabasePath       string
}

func (c Config) BindAddr() string {
	return c.Host + ":" + c.Port
}

func (c Config) GithubOAuth2Config() oauth2.Config {
	return oauth2.Config{
		ClientID:     c.GithubClientID,
		ClientSecret: c.GithubClientSecret,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}

func Parse() (c Config) {
	flag.StringVar(&c.Host, "host", "localhost", "host")
	flag.StringVar(&c.Port, "port", "8080", "port")
	flag.StringVar(&c.GithubClientID, "github-client-id", os.Getenv("GITHUB_CLIENT_ID"), "github client id")
	flag.StringVar(&c.GithubClientSecret, "github-client-secret", os.Getenv("GITHUB_CLIENT_SECRET"), "github client secret")
	flag.StringVar(&c.SessionSecret, "session-secret", os.Getenv("SESSION_SECRET"), "session secret")
	flag.StringVar(&c.DatabasePath, "database-path", "./chat.db", "database path")
	flag.Parse()
	return c
}
