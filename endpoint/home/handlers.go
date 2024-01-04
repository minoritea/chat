package home

import (
	"log"
	"net/http"

	"github.com/minoritea/chat/domain/message"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/resource"
)

type Container = resource.Container

func GetHandler(c Container) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			message.Data
			session.FlashData
		}
		data.Flashes = session.MustGetFlashes(c, w, r)
		baseData, err := message.GetMessageData(r.Context(), c.Querier().ListNewestMessages, message.FetchLimit)
		if err != nil {
			log.Println(err)
			data.Flashes = append(data.Flashes, session.NewErrorFlash("Failed to fetch messages"))
			c.Renderer().RenderHTML(w, "home", data, http.StatusInternalServerError)
			return
		}
		data.Data = baseData
		data.IsTerminal = len(data.Messages) < 20
		c.Renderer().RenderOkHTML(w, "home", data)
	})
}
