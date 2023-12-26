package database

import (
	"database/sql"
	"errors"

	"github.com/oklog/ulid/v2"
)

func IsRecordNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func NewID() string {
	return ulid.Make().String()
}
