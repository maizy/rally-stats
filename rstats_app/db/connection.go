package db

import (
	"database/sql"
	"errors"
	"os"

	_ "modernc.org/sqlite"
)

type Connection struct {
	Times *sql.DB
	Data  *sql.DB
}

var DBNotFound = errors.New("DB not found")
var UnableToOpenDB = errors.New("Unable to open DB")
var UnableToCheckDB = errors.New("Unable to check DB")

func CheckAndOpenReadonly(defaultPath, envVar string) (*sql.DB, error) {
	path := defaultPath
	if pathFromEnv, ok := os.LookupEnv(envVar); ok {
		path = pathFromEnv
	}
	if _, err := os.Stat(path); err != nil {
		return nil, DBNotFound
	}
	conn, err := sql.Open("sqlite", path+"?mode=readonly")
	if err != nil {
		return nil, UnableToOpenDB
	}
	if err := conn.Ping(); err != nil {
		return nil, UnableToCheckDB
	}
	return conn, nil
}
