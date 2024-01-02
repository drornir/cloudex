package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	RawConn *sql.DB
	*Queries
}

func OpenLibSQL(url string) (*DB, error) {
	if strings.HasPrefix(url, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("expanding `~` to user home dir: %w", err)
		}

		url = strings.Replace(url, "~", home, 1)
	}

	log.Printf("connecting to db at %q", url)
	conn, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, fmt.Errorf("opening db: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("pinging db: %w", err)
	}

	qs := New(conn)
	return &DB{
		RawConn: conn,
		Queries: qs,
	}, nil
}
