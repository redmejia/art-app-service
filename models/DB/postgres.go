package db

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	Db *sql.DB
}

const (
	OpenConns = 10
	IdleConns = 3
	LifeTime  = 60 * time.Second
)

// DSN connection
func DSNConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(OpenConns)
	db.SetMaxIdleConns(IdleConns)
	db.SetConnMaxLifetime(LifeTime)

	return db, nil
}
