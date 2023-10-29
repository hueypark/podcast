package persistence

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

// NewSqliteDB creates a new sqlite database connection.
func NewSqliteDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
