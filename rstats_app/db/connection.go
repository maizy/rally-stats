package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "modernc.org/sqlite"

	"dev.maizy.ru/rstats/rstats_app/dicts"
)

type DBContext struct {
	Times *sql.DB
	Dicts dicts.Dicts
}

var UnableToOpenDB = errors.New("unable to open DB")
var UnableToCheckDB = errors.New("unable to check DB")

func CheckAndOpenReadonly(defaultPath, envVar string) (*sql.DB, error) {
	path := defaultPath
	if pathFromEnv, ok := os.LookupEnv(envVar); ok {
		path = pathFromEnv
	}
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("not found %s", path)
	}
	conn, err := sql.Open("sqlite", "file:"+path+"?mode=ro")
	if err != nil {
		return nil, UnableToOpenDB
	}
	if err := conn.Ping(); err != nil {
		return nil, UnableToCheckDB
	}
	return conn, nil
}
