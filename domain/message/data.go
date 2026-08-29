// Package message provides domain logic for messages.
package message

import (
	"github.com/minoritea/chat/database"
	"github.com/samber/lo/mutable"
)

// Data holds fetched messages and metadata for rendering.
type Data struct {
	IsTerminal    bool
	Messages      []database.IMessage
	MightHaveMore bool
	Action        string
}

// Reverse reverses the order of the messages.
func (d *Data) Reverse() {
	mutable.Reverse(d.Messages)
}
