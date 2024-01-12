package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
)

func requireSession(c Container) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Cache-Control", "private")
			s, err := session.Get(c, r)
			if err != nil {
				log.Println(err)
				http.Redirect(w, r, "/auth", http.StatusSeeOther)
				return
			}
			sessionUser, err := session.GetUserFromSession(r.Context(), c, s)
			if err != nil {
				log.Println(err)
				http.Redirect(w, r, "/auth", http.StatusSeeOther)
				return
			}
			ctx := user.SetToContext(r.Context(), sessionUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}

func logger(next http.Handler) http.Handler {
	return middleware.RequestLogger(
		&middleware.DefaultLogFormatter{
			Logger: log.New(log.Writer(), "", log.LstdFlags),
		},
	)(next)
}
