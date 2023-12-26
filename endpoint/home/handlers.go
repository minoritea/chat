package home

import (
	"net/http"

	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/resource"
	"github.com/samber/lo"
)

type Container = resource.Container

type Data struct {
	Messages []database.ListMessagesRow
	Flashes  []session.Flash
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
		data.Messages = lo.Reverse(messages)
		renderer.RenderOkHTML(w, "home", data)
	})
}
