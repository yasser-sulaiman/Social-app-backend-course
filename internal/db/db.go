package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns) // Set the maximum number of open connections to the database
	db.SetMaxIdleConns(maxIdleConns) // Set the maximum number of idle connections in the pool

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration) // Set the maximum amount of time a connection may be idle before being closed

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the connection is alive
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}