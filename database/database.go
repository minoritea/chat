package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New(path string) (*Querier, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &Querier{db: db}, nil
}
