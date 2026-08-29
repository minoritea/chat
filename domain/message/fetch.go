package message

import (
	"context"

	"github.com/minoritea/chat/database"
)

// FetchLimit is the maximum number of messages fetched at once.
const FetchLimit = 20

// Query is a function that fetches messages.
type Query[T database.IMessage, P any] func(context.Context, P) ([]T, error)

// GetMessageData fetches messages with query and returns them as Data.
func GetMessageData[T database.IMessage, P any](ctx context.Context, query Query[T, P], param P) (data Data, err error) {
	rows, err := query(ctx, param)
	if err != nil {
		return data, err
	}
	data.Messages = database.RowsToMessages(rows)
	return data, nil
}
