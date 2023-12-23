package database

import (
	"errors"

	"github.com/mattn/go-sqlite3"
	"github.com/oklog/ulid/v2"
)

func IsSqliteNotFound(err error) bool {
	return IsSqliteError(err, sqlite3.ErrNotFound)
}

func IsSqliteError(err error, code sqlite3.ErrNo) bool {
	if err == nil {
		return false
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		return sqliteErr.Code == code
	}

	// log.Println("not sqlite error")

	var sqliteErrPtr *sqlite3.Error
	if errors.As(err, &sqliteErrPtr) && sqliteErrPtr != nil {
		return sqliteErrPtr.Code == code
	}
	//	log.Println("not sqlite error pointer")

	return false
}

func NewID() string {
	return ulid.Make().String()
}
