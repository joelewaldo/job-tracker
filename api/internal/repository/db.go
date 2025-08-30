package repository

import (
	"database/sql"
	"time"

	"github.com/joelewaldo/job-tracker/api/pkg/logger"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Log.WithError(err).Fatal("Failed to open database connection")
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	if err := conn.Ping(); err != nil {
		logger.Log.WithError(err).Fatal("Failed to ping database")
		return nil, err
	}

	logger.Log.Info("Database connection established")
	return &DB{Conn: conn}, nil
}

func (db *DB) Close() error {
	return db.Conn.Close()
}
