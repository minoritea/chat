package main

import (
	"log"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/server"
)

func run() error {
	c, err := container.New()
	if err != nil {
		return err
	}

	return server.ListenAndServe("127.0.0.1:8080", c)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

type session struct{ userID string }
