package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/resource"
	"github.com/minoritea/chat/router"
)

var version = "0.0.0"

func run() error {
	var conf config.Config
	conf.DatabaseDriver = "sqlite3"
	conf.Version = version
	flag.StringVar(&conf.Host, "host", "localhost", "host")
	flag.StringVar(&conf.Port, "port", "8080", "port")
	flag.StringVar(&conf.GithubClientID, "github-client-id", os.Getenv("GITHUB_CLIENT_ID"), "github client id")
	flag.StringVar(&conf.GithubClientSecret, "github-client-secret", os.Getenv("GITHUB_CLIENT_SECRET"), "github client secret")
	flag.StringVar(&conf.SessionSecret, "session-secret", os.Getenv("SESSION_SECRET"), "session secret")
	flag.StringVar(&conf.DatabasePath, "database-path", "./chat.db", "database path")
	flag.Parse()

	c, err := resource.New(conf)
	if err != nil {
		return err
	}

	r := router.New(*c)
	return http.ListenAndServe(conf.BindAddr(), r)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
