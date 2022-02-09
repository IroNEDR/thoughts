package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Conn struct {
	Pool *pgxpool.Pool
}

func NewConnectionPool(dsn string) (*Conn, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(context.Background()); err != nil {
		return nil, err
	}
	return &Conn{Pool: db}, nil
}
