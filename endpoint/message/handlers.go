package message

import (
	"log"
	"net/http"
	"time"

	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
	"github.com/minoritea/chat/resource"
	"github.com/samber/lo"
)

type Container = resource.Container

func PostHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		if message == "" {
			log.Println("message is empty")
			session.MustAddFlash(c, w, r, session.NewErrorFlash("Message is empty"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		_, err := c.Queries().CreateMessage(r.Context(), database.CreateMessageParams{
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

type Data struct {
	ReachedStart bool
	BeforeID     string
	Messages     []database.ListMessagesBeforeIDRow
}

func GetHandler(c Container) http.HandlerFunc {
	renderer := c.Renderer()
	return func(w http.ResponseWriter, r *http.Request) {
		beforeID := r.URL.Query().Get("before_id")
		if beforeID == "" {
			log.Println("message_id is empty")
			session.MustAddFlash(c, w, r, session.NewErrorFlash("message_id is empty"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		messages, err := c.Queries().ListMessagesBeforeID(r.Context(), database.ListMessagesBeforeIDParams{
			ID:    beforeID,
			Limit: 20,
		})
		if err != nil {
			log.Println(err)
			session.MustAddFlash(c, w, r, session.NewErrorFlash("Internal Server Error"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		var data Data
		data.BeforeID = beforeID
		data.Messages = lo.Reverse(messages)
		if len(messages) < 20 {
			data.ReachedStart = true
		}
		renderer.RenderStream(w, "message", data, http.StatusOK)
	}
}
