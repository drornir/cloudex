package db

import (
	"database/sql"
	"fmt"

	"github.com/drornir/cloudex/pkg/db/queries"
	_ "github.com/libsql/go-libsql"
)

type DB struct {
	conn *sql.DB
	*queries.Queries
}

func New(dbURL string) (*DB, error) {
	conn, err := sql.Open("libsql", dbURL)
	if err != nil {
		return nil, fmt.Errorf("opening db file: %w", err)
	}
	qs := queries.New(conn)

	return &DB{
		conn:    conn,
		Queries: qs,
	}, nil
}
