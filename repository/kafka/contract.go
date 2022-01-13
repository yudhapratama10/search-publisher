package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	model "github.com/yudhapratama10/search-publisher/model"
)

type writeMessages interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}

type footballRepository struct {
	kafka writeMessages
}

type FootballRepositoryContract interface {
	Produce(footballClub model.FootballClub, operation string) (model.FootballClub, error)
}

func NewFootballRepository(kafka writeMessages) FootballRepositoryContract {
	return &footballRepository{
		kafka: kafka,
	}
}
