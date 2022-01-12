package pg

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type dbConn interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
}

var conn dbConn = nil

func InitClient() (dbConn, error) {
	conString := "postgres://postgres:@localhost:5432/postgres"

	dbConn, err := pgxpool.Connect(context.Background(), conString)
	if err != nil {
		return nil, err
	}

	if err = dbConn.Ping(context.Background()); err != nil {
		return nil, err
	}

	conn = dbConn

	return conn, nil
}
