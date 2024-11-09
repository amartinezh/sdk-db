package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type DB struct {
	Conn *pgx.Conn
}

func NewDB(connString string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &DB{
		Conn: conn,
	}, nil
}

func (db *DB) Close() {
	db.Conn.Close(context.Background())
}
