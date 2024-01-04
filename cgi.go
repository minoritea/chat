//go:build cgi

package main

import (
	"log"
	"net/http/cgi"

	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/resource"
	"github.com/minoritea/chat/router"
	_ "modernc.org/sqlite"
)

var (
	GithubClientID     = ""
	GithubClientSecret = ""
	SessionSecret      = ""
	DatabasePath       = ""
)

func run() error {
	var conf config.Config
	conf.GithubClientID = GithubClientID
	conf.GithubClientSecret = GithubClientSecret
	conf.SessionSecret = SessionSecret
	conf.DatabasePath = DatabasePath
	conf.DatabaseDriver = "sqlite"

	c, err := resource.New(conf)
	if err != nil {
		return err
	}

	return cgi.Serve(router.New(*c))
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
