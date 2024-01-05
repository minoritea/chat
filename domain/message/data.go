package message

import (
	"github.com/minoritea/chat/database"
	"github.com/samber/lo"
)

type Data struct {
	IsTerminal    bool
	Messages      []database.IMessage
	MightHaveMore bool
	Action        string
}

func (d *Data) Reverse() {
	d.Messages = lo.Reverse(d.Messages)
}
