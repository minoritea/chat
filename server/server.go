package server

import (
	"github.com/minoritea/chat/container"
	"net/http"
)

type Container = *container.Container

func New(bind string, c Container) *http.Server {
	r := NewRouter(c)
	server := &http.Server{
		Addr:    bind,
		Handler: r,
	}
	return server
}

func ListenAndServe(bind string, c Container) error {
	s := New(bind, c)
	defer s.Close()
	return s.ListenAndServe()
}
