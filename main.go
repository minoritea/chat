// Package main is the root package of the chat application.
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/resource"
	"github.com/minoritea/chat/router"
	"golang.org/x/sync/errgroup"
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
	srv := &http.Server{
		Addr:              conf.BindAddr(),
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	g := new(errgroup.Group)
	g.Go(srv.ListenAndServe)
	g.Go(func() error {
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		ctx, cancel = signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL)
		defer cancel()
		return srv.Shutdown(ctx)
	})
	return g.Wait()
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
