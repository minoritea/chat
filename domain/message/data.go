package message

import (
	"github.com/minoritea/chat/database"
)

type Data struct {
	IsTerminal    bool
	Messages      []database.IMessage
	MightHaveMore bool
}

type StreamData struct {
	Data
	TargetID string
	Action   string
}
