package kafka

import (
	"github.com/segmentio/kafka-go"
	model "github.com/yudhapratama10/search-publisher/model"
)

type footballRepository struct {
	kafka *kafka.Writer
}

type FootballRepositoryContract interface {
	Produce(footballClub model.FootballClub, operation string) (model.FootballClub, error)
}

func NewFootballRepository(kafka *kafka.Writer) FootballRepositoryContract {
	return &footballRepository{
		kafka: kafka,
	}
}
