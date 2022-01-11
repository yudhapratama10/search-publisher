package kafka

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/segmentio/kafka-go"
	model "github.com/yudhapratama10/search-publisher/model"
)

var topic string = "test-messages"

func (repo *footballRepository) Produce(footballClub model.FootballClub, operation string) (model.FootballClub, error) {

	dataJson, err := json.Marshal(map[string]interface{}{
		"data":      footballClub,
		"operation": operation,
	})

	if err != nil {
		return model.FootballClub{}, err
	}

	err = repo.kafka.WriteMessages(context.Background(),
		kafka.Message{
			Topic: topic,
			Key:   []byte(strconv.Itoa(footballClub.Id)),
			Value: dataJson,
		},
	)

	if err != nil {
		return model.FootballClub{}, err
	}

	return footballClub, nil
}
