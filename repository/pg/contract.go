package pg

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	model "github.com/yudhapratama10/search-publisher/model"
)

type footballRepository struct {
	db dbConn
}

type dbConn interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
}

type FootballRepositoryContract interface {
	Get(id int) (model.FootballClub, error)
	Insert(footballClub model.FootballClub) (model.FootballClub, string, error)
	Update(footballClub model.FootballClub) (model.FootballClub, string, error)
	Delete(footballClub model.FootballClub) (model.FootballClub, string, error)
}

func NewFootballRepository(dbConn dbConn) FootballRepositoryContract {
	return &footballRepository{
		db: dbConn,
	}
}
