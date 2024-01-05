package message

import (
	"context"

	"github.com/minoritea/chat/database"
)

const FetchLimit = 20

type Query[T database.IMessage, P any] func(context.Context, P) ([]T, error)

func GetMessageData[T database.IMessage, P any](ctx context.Context, query Query[T, P], param P) (data Data, err error) {
	rows, err := query(ctx, param)
	if err != nil {
		return data, err
	}
	data.Messages = database.RowsToMessages(rows)
	return data, nil
}
