package pg

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
	model "github.com/yudhapratama10/search-publisher/model"
)

type footballRepository struct {
	db *pgxpool.Pool
}

type FootballRepositoryContract interface {
	Get(id int) (model.FootballClub, error)
	Insert(footballClub model.FootballClub) (model.FootballClub, string, error)
	Update(footballClub model.FootballClub) (model.FootballClub, string, error)
	Delete(footballClub model.FootballClub) (model.FootballClub, string, error)
}

func NewFootballRepository(db *pgxpool.Pool, kafka *kafka.Writer) FootballRepositoryContract {
	return &footballRepository{
		db: db,
	}
}
