package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/minoritea/chat/asset"
	"github.com/minoritea/chat/endpoint/auth"
	"github.com/minoritea/chat/endpoint/home"
	"github.com/minoritea/chat/endpoint/message"
	"github.com/minoritea/chat/resource"
)

const requestBodyLimit int64 = 1024 * 1024 // Accept request bodies at most 1M

// Container is an alias for resource.Container.
type Container = resource.Container

// New returns the application's HTTP handler with all routes registered.
func New(c Container) http.Handler {
	r := chi.NewRouter()
	r.Use(
		maxBytes(requestBodyLimit),
		middleware.ClientIPFromHeader("CF-Connecting-IP"),
		logger,
		middleware.Recoverer,
		http.NewCrossOriginProtection().Handler,
	)

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.NoCache,
		)

		// routes that require a session
		r.Group(func(r chi.Router) {
			r.Use(requireSession(c))
			r.Get("/", home.GetHandler(c))
			r.Get("/messages", message.GetHandler(c))
			r.Get("/messages/more", message.GetMoreHandler(c))
			r.Post("/messages", message.PostHandler(c))
		})

		// routes that don't require a session
		r.Get("/auth", auth.GetHandler(c))
		r.Post("/auth", auth.PostHandler(c))
		r.Get("/auth/callback", auth.GetCallbackHandler(c))
	})

	// static assets
	r.Route(c.Config().AssetPath(), func(r chi.Router) {
		r.Use(
			sourceMap,
			middleware.PathRewrite(c.Config().AssetPath(), ""),
			middleware.SetHeader("Cache-Control", "immutable; max-age=31536000"),
		)
		r.Get("/js/*", http.FileServer(http.FS(asset.FS)).ServeHTTP)
		r.Get("/css/*", http.FileServer(http.FS(asset.FS)).ServeHTTP)
	})
	r.Get("/favicon.ico", http.FileServer(http.FS(asset.FS)).ServeHTTP)
	return r
}
