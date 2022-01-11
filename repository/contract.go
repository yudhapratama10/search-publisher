package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
	model "github.com/yudhapratama10/search-publisher/model"
)

// type SourceResult struct {
// 	Source `json:"_source"`
// }

type footballRepository struct {
	db    *pgxpool.Pool
	kafka *kafka.Writer
}

type FootballRepositoryContract interface {
	Insert(footballClub model.FootballClub) (model.FootballClub, error)
	Update(footballClub model.FootballClub) (model.FootballClub, error)
	Get(id int) (model.FootballClub, error)
	// Delete
}

func NewFootballRepository(db *pgxpool.Pool, kafka *kafka.Writer) FootballRepositoryContract {
	return &footballRepository{
		db:    db,
		kafka: kafka,
	}
}
