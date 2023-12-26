package home

import (
	"net/http"
	"time"

	"github.com/minoritea/chat/resource"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/session"
)

type Container = resource.Container

type Data struct {
	Messages []database.ListMessagesRow
	Flashes  []session.Flash
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

func GetHandler(c Container) http.HandlerFunc {
	renderer := c.Renderer()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data Data
		data.Flashes = session.MustGetFlashes(c, w, r)
		messages, err := c.Queries().ListMessages(r.Context(), 20)
		if err != nil {
			data.Flashes = append(data.Flashes, session.NewErrorFlash("Failed to fetch messages"))
			renderer.RenderHTML(w, "home", data, http.StatusInternalServerError)
			return
		}
		data.Messages = messages
		renderer.RenderOkHTML(w, "home", data)
	})
}
