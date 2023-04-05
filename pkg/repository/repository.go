package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
}

func New(conn string) (*PostgresDB, error) {
	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		Pool: pool,
	}, nil
}
