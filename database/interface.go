// Code generated by ifacemaker; DO NOT EDIT.

package database

import (
	"context"
)

// Querier ...
type Querier interface {
	CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByAccount(ctx context.Context, account string) (User, error)
	GetUserBySessionID(ctx context.Context, id string) (User, error)
	ListMessagesAfterID(ctx context.Context, arg ListMessagesAfterIDParams) ([]ListMessagesAfterIDRow, error)
	ListMessagesBeforeID(ctx context.Context, arg ListMessagesBeforeIDParams) ([]ListMessagesBeforeIDRow, error)
	ListNewestMessages(ctx context.Context, limit int64) ([]ListNewestMessagesRow, error)
	ListOldestMessages(ctx context.Context, limit int64) ([]ListOldestMessagesRow, error)
}
