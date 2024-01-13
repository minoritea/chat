package config

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

func (c Config) BindAddr() string {
	return c.Host + ":" + c.Port
}

func (c Config) AssetPath() string {
	return "/asset-" + c.Version
}
