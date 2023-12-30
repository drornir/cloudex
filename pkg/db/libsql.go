package db

import (
	"database/sql"
	"fmt"

	_ "github.com/libsql/go-libsql"
)

type DB struct {
	db *sql.DB
	*Queries
}

func OpenLibSQL(url string) (*DB, error) {
	conn, err := sql.Open("libsql", url)
	if err != nil {
		return nil, fmt.Errorf("opening db: %w", err)
	}

	qs := New(conn)
	return &DB{
		db:      conn,
		Queries: qs,
	}, nil
}
