package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

func IsRecordNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func NewID() string {
	return ulid.Make().String()
}

func (r ListMessagesAfterIDRow) GetID() string            { return r.ID }
func (r ListMessagesAfterIDRow) GetMessage() string       { return r.Message }
func (r ListMessagesAfterIDRow) GetCreatedAt() time.Time  { return r.CreatedAt }
func (r ListMessagesAfterIDRow) GetAccount() string       { return r.Account }
func (r ListMessagesBeforeIDRow) GetID() string           { return r.ID }
func (r ListMessagesBeforeIDRow) GetMessage() string      { return r.Message }
func (r ListMessagesBeforeIDRow) GetCreatedAt() time.Time { return r.CreatedAt }
func (r ListMessagesBeforeIDRow) GetAccount() string      { return r.Account }
func (r ListNewestMessagesRow) GetID() string             { return r.ID }
func (r ListNewestMessagesRow) GetMessage() string        { return r.Message }
func (r ListNewestMessagesRow) GetCreatedAt() time.Time   { return r.CreatedAt }
func (r ListNewestMessagesRow) GetAccount() string        { return r.Account }
func (r ListOldestMessagesRow) GetID() string             { return r.ID }
func (r ListOldestMessagesRow) GetMessage() string        { return r.Message }
func (r ListOldestMessagesRow) GetCreatedAt() time.Time   { return r.CreatedAt }
func (r ListOldestMessagesRow) GetAccount() string        { return r.Account }

type IMessage interface {
	GetID() string
	GetMessage() string
	GetCreatedAt() time.Time
	GetAccount() string
}

func RowsToMessages[T IMessage](ms []T) []IMessage {
	result := make([]IMessage, len(ms))
	for i, m := range ms {
		result[i] = m
	}
	return result
}
