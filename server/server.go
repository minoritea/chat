package server

import (
	"github.com/minoritea/chat/resource"
	"net/http"
)

type Container = resource.Container

func New(c Container) *http.Server {
	r := NewRouter(c)
	server := &http.Server{
		Addr:    c.Config().BindAddr(),
		Handler: r,
	}
	return server
}

func ListenAndServe(c Container) error {
	s := New(c)
	defer s.Close()
	return s.ListenAndServe()
}
