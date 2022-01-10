package pg

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitClient() (*pgxpool.Pool, error) {
	conString := "postgres://postgres:@localhost:5432/postgres"

	conn, err := pgxpool.Connect(context.Background(), conString)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
