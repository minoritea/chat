package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
)

func withMiddlewares(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return chi.Chain(middlewares...).Handler(next).ServeHTTP
	}
}

func requireSession(c Container) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionUser, err := session.GetUserFromSession(c, r)
			if err == nil {
				ctx := user.SetToContext(r.Context(), sessionUser)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			log.Println(err)
			if errors.Is(err, session.SessionNotFound) {
				http.Redirect(w, r, "/signin", http.StatusFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		})
	}
}
