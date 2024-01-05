package message_test

import (
	"context"
	"testing"
	"time"

	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/message"
)

type Message struct{ id string }

func (m Message) GetID() string           { return m.id }
func (m Message) GetMessage() string      { return "" }
func (m Message) GetCreatedAt() time.Time { return time.Now() }
func (m Message) GetAccount() string      { return "" }

var _ database.IMessage = (*Message)(nil)

func TestGetMessageData(t *testing.T) {
	type Param struct{}
	messages := []Message{{id: "1"}, {id: "2"}, {id: "3"}}
	query := func(ctx context.Context, param Param) ([]Message, error) {
		return messages, nil
	}
	data, err := message.GetMessageData(context.Background(), query, Param{})
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Messages) != len(messages) {
		t.Fatalf("len(data) = %d, want %d", len(data.Messages), len(messages))
	}
	for i, m := range data.Messages {
		want := messages[i]
		if m.GetID() != want.id {
			t.Errorf("data[%d].ID = %s, want %s", i, m.GetID(), want.id)
		}
	}
}
