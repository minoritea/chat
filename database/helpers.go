// Package database provides database access.
package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

// ErrRecordNotFound is an alias for sql.ErrNoRows.
var ErrRecordNotFound = sql.ErrNoRows

// IsRecordNotFound reports whether err is sql.ErrNoRows.
func IsRecordNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

// NewID returns a new ULID for database records.
func NewID() string {
	return ulid.Make().String()
}

// GetID returns the ID.
func (r ListMessagesAfterIDRow) GetID() string { return r.ID }

// GetMessage returns the message.
func (r ListMessagesAfterIDRow) GetMessage() string { return r.Message }

// GetCreatedAt returns the creation time.
func (r ListMessagesAfterIDRow) GetCreatedAt() time.Time { return r.CreatedAt }

// GetAccount returns the account name.
func (r ListMessagesAfterIDRow) GetAccount() string { return r.Account }

// GetID returns the ID.
func (r ListMessagesBeforeIDRow) GetID() string { return r.ID }

// GetMessage returns the message.
func (r ListMessagesBeforeIDRow) GetMessage() string { return r.Message }

// GetCreatedAt returns the creation time.
func (r ListMessagesBeforeIDRow) GetCreatedAt() time.Time { return r.CreatedAt }

// GetAccount returns the account name.
func (r ListMessagesBeforeIDRow) GetAccount() string { return r.Account }

// GetID returns the ID.
func (r ListNewestMessagesRow) GetID() string { return r.ID }

// GetMessage returns the message.
func (r ListNewestMessagesRow) GetMessage() string { return r.Message }

// GetCreatedAt returns the creation time.
func (r ListNewestMessagesRow) GetCreatedAt() time.Time { return r.CreatedAt }

// GetAccount returns the account name.
func (r ListNewestMessagesRow) GetAccount() string { return r.Account }

// GetID returns the ID.
func (r ListOldestMessagesRow) GetID() string { return r.ID }

// GetMessage returns the message.
func (r ListOldestMessagesRow) GetMessage() string { return r.Message }

// GetCreatedAt returns the creation time.
func (r ListOldestMessagesRow) GetCreatedAt() time.Time { return r.CreatedAt }

// GetAccount returns the account name.
func (r ListOldestMessagesRow) GetAccount() string { return r.Account }

// MessageData is implemented by message row types.
type MessageData interface {
	GetID() string
	GetMessage() string
	GetCreatedAt() time.Time
	GetAccount() string
}

// RowsToMessages converts a slice of T to a slice of MessageData.
func RowsToMessages[T MessageData](ms []T) []MessageData {
	result := make([]MessageData, len(ms))
	for i, m := range ms {
		result[i] = m
	}
	return result
}
