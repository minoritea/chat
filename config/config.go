package config

import (
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
	DatabaseDriver     string
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
