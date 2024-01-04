//go:build !cgi

package main

import (
	"log"

	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/resource"
	"github.com/minoritea/chat/server"
)

func run() error {
	conf := config.Parse()
	c, err := resource.New(conf)
	if err != nil {
		return err
	}

	return server.ListenAndServe(*c)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
