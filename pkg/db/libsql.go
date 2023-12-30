package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/libsql/go-libsql"
)

type DB struct {
	db *sql.DB
	*Queries
}

func OpenLibSQL(url string) (*DB, error) {
	if strings.HasPrefix(url, "file://~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("expanding `~` to user home dir: %w", err)
		}

		url = strings.Replace(url, "~", home, 1)
	}

	conn, err := sql.Open("libsql", url)
	if err != nil {
		return nil, fmt.Errorf("opening db: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("pinging db: %w", err)
	}

	qs := New(conn)
	return &DB{
		db:      conn,
		Queries: qs,
	}, nil
}
