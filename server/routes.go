package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/minoritea/chat/features/home"
	"github.com/minoritea/chat/features/message"
	"github.com/minoritea/chat/features/signin"
	"github.com/minoritea/chat/features/signup"
)

func NewRouter(c Container) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", withMiddlewares(requireSession(c))(home.GetHandler(c)))
	r.Post("/messages", withMiddlewares(requireSession(c))(message.PostHandler(c)))
	r.Get("/signin", signin.GetHandler(c))
	r.Post("/signin", signin.PostHandler(c))
	r.Get("/signup", signup.GetHandler(c))
	r.Post("/signup", signup.PostHandler(c))
	return r
}
