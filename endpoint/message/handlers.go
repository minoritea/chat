package message

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/message"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
	"github.com/minoritea/chat/resource"
	"github.com/samber/lo"
)

type Container = resource.Container
type Data = message.StreamData

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

		http.Redirect(w, r, "/messages/more", http.StatusSeeOther)
	}
}

func GetHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		beforeID := r.URL.Query().Get("before_id")
		afterID := r.URL.Query().Get("after_id")

		switch [2]bool{beforeID != "", afterID != ""} {
		case [2]bool{false, false}: // no query parameters are set
			handleOldestMessages(c, w, r)
		case
			[2]bool{true, true}: // multiple query parameters are set
			log.Println("multiple query parameters are set")
			session.MustAddFlash(c, w, r, session.NewErrorFlash("both before_id and after_id are set"))
			http.Redirect(w, r, "/", http.StatusSeeOther)

		case [2]bool{true, false}: // only before_id is set
			handleGetMessagesBeforeID(c, beforeID, w, r)

		case [2]bool{false, true}: // only after_id is set
			handleGetMessagesAfterID(c, afterID, w, r)
		}
	}
}

func GetMoreHandler(c Container) http.HandlerFunc {
	renderer := c.Renderer()
	return func(w http.ResponseWriter, r *http.Request) {
		var data Data
		data.TargetID = "messages"
		data.Action = "append"
		data.MightHaveMore = true
		renderer.RenderStream(w, "messages", data, http.StatusOK)
	}
}

func handleGetMessagesByID[P any, R database.IMessage](
	c Container,
	w http.ResponseWriter,
	r *http.Request,
	listQuery func(context.Context, P) ([]R, error),
	param P,
	dataHandler func([]R) Data,
) {
	messages, err := listQuery(r.Context(), param)
	if err != nil {
		log.Println(err)
		session.MustAddFlash(c, w, r, session.NewErrorFlash("Internal Server Error"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data := dataHandler(messages)
	c.Renderer().RenderStream(w, "messages", data, http.StatusOK)
}

func handleGetMessagesBeforeID(c Container, beforeID string, w http.ResponseWriter, r *http.Request) {
	handleGetMessagesByID(c, w, r, c.Queries().ListMessagesBeforeID, database.ListMessagesBeforeIDParams{
		ID:    beforeID,
		Limit: 20,
	}, func(messages []database.ListMessagesBeforeIDRow) Data {
		var data Data
		data.Action = "before"
		data.TargetID = beforeID
		data.Messages = lo.Reverse(database.RowsToMessages(messages))
		if len(messages) < 20 {
			data.IsTerminal = true
		}
		return data
	})
}

func handleGetMessagesAfterID(c Container, afterID string, w http.ResponseWriter, r *http.Request) {
	handleGetMessagesByID(c, w, r, c.Queries().ListMessagesAfterID, database.ListMessagesAfterIDParams{
		ID:    afterID,
		Limit: 20,
	}, func(messages []database.ListMessagesAfterIDRow) Data {
		var data Data
		data.Messages = lo.Reverse(database.RowsToMessages(messages))
		data.TargetID = afterID
		data.Action = "after"
		data.MightHaveMore = len(data.Messages) == 20
		return data
	})
}

func handleOldestMessages(c Container, w http.ResponseWriter, r *http.Request) {
	messages, err := c.Queries().ListOldestMessages(r.Context(), 20)
	if err != nil {
		log.Println(err)
		session.MustAddFlash(c, w, r, session.NewErrorFlash("Internal Server Error"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var data Data
	data.IsTerminal = true
	data.Action = "prepend"
	data.TargetID = "messages"
	data.Messages = lo.Reverse(database.RowsToMessages(messages))
	data.MightHaveMore = len(messages) == 20
	c.Renderer().RenderStream(w, "messages", data, http.StatusOK)
}
