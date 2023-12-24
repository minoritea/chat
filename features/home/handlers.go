package home

import (
	"log"
	"net/http"
	"time"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/user"
)

type Container = container.Container

type Data struct {
	Error      string
	InputError string
	Messages   []database.ListMessagesRow
}

func (d Data) FormatCreatedAt(t time.Time) string {
	return t.Format(time.DateTime)
}

func (d Data) ReversedMessages() []database.ListMessagesRow {
	l := len(d.Messages)
	reversed := make([]database.ListMessagesRow, l)
	for i := range d.Messages {
		reversed[l-i-1] = d.Messages[i]
	}
	return reversed
}

func GetHandler(c *Container) http.HandlerFunc {
	renderer := c.GetTemplateRenderer()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data Data
		messages, err := c.GetQueries().ListMessages(r.Context(), 20)
		if err != nil {
			log.Println(err)
			data.Error = "Internal Server Error"
			renderer.RenderHTML(w, "home", data, http.StatusInternalServerError)
			return
		}
		data.Messages = messages
		renderer.RenderOkHTML(w, "home", data)
	})
}

func PostHandler(c *Container) http.HandlerFunc {
	renderer := c.GetTemplateRenderer()
	return func(w http.ResponseWriter, r *http.Request) {
		var data Data
		message := r.PostFormValue("message")
		if message == "" {
			log.Println("message is empty")
			data.InputError = "message is empty"
			renderer.RenderHTML(w, "home", data, http.StatusBadRequest)
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
			data.Error = "Internal Server Error"
			renderer.RenderHTML(w, "home", data, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
