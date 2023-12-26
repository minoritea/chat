package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/minoritea/chat/features/auth"
	"github.com/minoritea/chat/features/home"
	"github.com/minoritea/chat/features/message"
)

func NewRouter(c Container) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", withMiddlewares(requireSession(c))(home.GetHandler(c)))
	r.Post("/messages", withMiddlewares(requireSession(c))(message.PostHandler(c)))
	r.Get("/auth", auth.GetHandler(c))
	r.Post("/auth", auth.PostHandler(c))
	r.Get("/auth/callback", auth.GetCallbackHandler(c))
	return r
}
