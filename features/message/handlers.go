package message

import (
	"log"
	"net/http"
	"time"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
)

type Container = container.Container

func PostHandler(c *Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		if message == "" {
			log.Println("message is empty")
			session.MustAddFlash(c, w, r, session.NewErrorFlash("Message is empty"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		_, err := c.GetQueries().CreateMessage(r.Context(), database.CreateMessageParams{
			ID:        database.NewID(),
			UserID:    user.FromContext(r.Context()).ID,
			Message:   message,
			CreatedAt: time.Now(),
		})
		if err != nil {
			log.Println(err)
			session.MustAddFlash(c, w, r, session.NewErrorFlash("Internal Server Error"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
