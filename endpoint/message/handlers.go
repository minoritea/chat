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
)

type Container = resource.Container
type Data = message.Data

func PostHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := session.MustGet(c, r)
		message := r.PostFormValue("message")
		if message == "" {
			log.Println("message is empty")
			session.RedirectWithErrorFlash(w, r, s, "/", "Message is empty")
			return
		}

		_, err := c.Querier().CreateMessage(r.Context(), database.CreateMessageParams{
			ID:        database.NewID(),
			UserID:    user.FromContext(r.Context()).ID,
			Message:   message,
			CreatedAt: time.Now(),
		})
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(w, r, s, "/", "Internal Server Error")
			return
		}

		session.MustSave(s, r, w)

		http.Redirect(w, r, "/messages/more", http.StatusSeeOther)
	}
}

func GetHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := session.MustGet(c, r)

		var (
			data Data
			err  error
		)

		beforeID := r.URL.Query().Get("before_id")
		afterID := r.URL.Query().Get("after_id")

		switch [2]bool{beforeID != "", afterID != ""} {
		case [2]bool{false, false}: // no query parameters are set
			data, err = getOldestMessagesData(r.Context(), c)
			if err != nil {
				log.Println(err)
				session.RedirectWithErrorFlash(w, r, s, "/", "Internal Server Error")
				return
			}
		case [2]bool{true, true}: // multiple query parameters are set
			log.Println("multiple query parameters are set")
			session.RedirectWithErrorFlash(w, r, s, "/", "Failed to fetch messages")
			return
		case [2]bool{true, false}: // only before_id is set
			data, err = getMessagesDataBeforeID(r.Context(), c, beforeID)
			if err != nil {
				log.Println(err)
				session.RedirectWithErrorFlash(w, r, s, "/", "Internal Server Error")
				return
			}
		case [2]bool{false, true}: // only after_id is set
			data, err = getMessagesDataAfterID(r.Context(), c, afterID)
			if err != nil {
				log.Println(err)
				session.RedirectWithErrorFlash(w, r, s, "/", "Internal Server Error")
				return
			}
		}

		session.MustSave(s, r, w)

		c.Renderer().RenderStream(w, "messages", data, http.StatusOK)
	}
}

func GetMoreHandler(c Container) http.HandlerFunc {
	renderer := c.Renderer()
	return func(w http.ResponseWriter, r *http.Request) {
		var data Data
		data.Action = "append"
		data.MightHaveMore = true
		renderer.RenderStream(w, "messages", data, http.StatusOK)
	}
}

func getMessagesDataBeforeID(ctx context.Context, c Container, beforeID string) (data Data, err error) {
	data, err = message.GetMessageData(ctx, c.Querier().ListMessagesBeforeID, database.ListMessagesBeforeIDParams{
		ID:    beforeID,
		Limit: message.FetchLimit,
	})
	if err != nil {
		return data, err
	}
	data.Action = "prepend"
	data.IsTerminal = len(data.Messages) < message.FetchLimit
	return data, nil
}

func getMessagesDataAfterID(ctx context.Context, c Container, afterID string) (data Data, err error) {
	data, err = message.GetMessageData(ctx, c.Querier().ListMessagesAfterID, database.ListMessagesAfterIDParams{
		ID:    afterID,
		Limit: message.FetchLimit,
	})
	if err != nil {
		return data, err
	}

	data.Action = "append"
	data.Reverse()
	data.MightHaveMore = len(data.Messages) == message.FetchLimit
	return data, nil
}

func getOldestMessagesData(ctx context.Context, c Container) (data Data, err error) {
	data, err = message.GetMessageData(ctx, c.Querier().ListOldestMessages, message.FetchLimit)
	if err != nil {
		return data, err
	}

	data.IsTerminal = true
	data.Action = "prepend"
	data.MightHaveMore = len(data.Messages) == message.FetchLimit
	return data, nil
}
