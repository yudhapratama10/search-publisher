package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
	"github.com/yudhapratama10/search-service/model"
)

// type SourceResult struct {
// 	Source `json:"_source"`
// }

type footballRepository struct {
	db    *pgxpool.Pool
	kafka *kafka.Writer
}

type FootballRepositoryContract interface {
	Insert(footballClub model.FootballClub) ([]model.FootballClub, error)
	// Update(keyword string) ([]model.FootballClub, error)
	// Delete
}

func NewRecipeRepository(db *pgxpool.Pool, kafka *kafka.Writer) FootballRepositoryContract {
	return &footballRepository{
		db:    db,
		kafka: kafka,
	}
}
