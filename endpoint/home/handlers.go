package home

import (
	"net/http"

	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/message"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/resource"
	"github.com/samber/lo"
)

type Container = resource.Container

func GetHandler(c Container) http.HandlerFunc {
	renderer := c.Renderer()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			message.Data
			session.FlashData
		}
		data.Flashes = session.MustGetFlashes(c, w, r)
		messages, err := c.Queries().ListNewestMessages(r.Context(), 20)
		if err != nil {
			data.Flashes = append(data.Flashes, session.NewErrorFlash("Failed to fetch messages"))
			renderer.RenderHTML(w, "home", data, http.StatusInternalServerError)
			return
		}
		data.IsTerminal = len(messages) < 20
		data.Messages = lo.Reverse(database.RowsToMessages(messages))
		renderer.RenderOkHTML(w, "home", data)
	})
}
