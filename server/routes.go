package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/minoritea/chat/asset"
	"github.com/minoritea/chat/endpoint/auth"
	"github.com/minoritea/chat/endpoint/home"
	"github.com/minoritea/chat/endpoint/message"
)

func NewRouter(c Container) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", withMiddlewares(requireSession(c))(home.GetHandler(c)))
	r.Get("/messages", withMiddlewares(requireSession(c))(message.GetHandler(c)))
	r.Get("/messages/more", withMiddlewares(requireSession(c))(message.GetMoreHandler(c)))
	r.Post("/messages", withMiddlewares(requireSession(c))(message.PostHandler(c)))
	r.Get("/auth", auth.GetHandler(c))
	r.Post("/auth", auth.PostHandler(c))
	r.Get("/auth/callback", auth.GetCallbackHandler(c))
	r.Get("/js/*", http.FileServer(http.FS(asset.FS)).ServeHTTP)
	return r
}
