// Package config provides the application configuration.
package config

// Config holds the application configuration.
type Config struct {
	Host               string
	Port               string
	GithubClientID     string
	GithubClientSecret string
	SessionSecret      string
	DatabasePath       string
	DatabaseDriver     string
	Version            string
}

// BindAddr returns the host:port address the server listens on.
func (c Config) BindAddr() string {
	return c.Host + ":" + c.Port
}

// AssetPath returns the versioned URL prefix of static assets.
func (c Config) AssetPath() string {
	return "/asset-" + c.Version
}
