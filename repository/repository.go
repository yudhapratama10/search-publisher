package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	model "github.com/yudhapratama10/search-publisher/model"
)

// var resp map[string]interface{}

func (repo *footballRepository) Insert(footballClub model.FootballClub) (model.FootballClub, error) {
	var id int

	err := repo.db.QueryRow(context.Background(),
		"INSERT INTO footballclub (name, tournaments, nation, hasstadium, description, rating) VALUES ($1, $2, $3, $4, $5, $6) returning id",
		footballClub.Name, footballClub.Tournaments, footballClub.Nation, footballClub.HasStadium, footballClub.Description, footballClub.Rating).Scan(&id)
	fmt.Println(err)
	if err != nil {
		return model.FootballClub{}, err
	}

	footballClub.Id = id
	topic := "test-messages"

	dataJson, err := json.Marshal(footballClub)
	if err != nil {
		return model.FootballClub{}, err
	}

	err = repo.kafka.WriteMessages(context.Background(),
		kafka.Message{
			Topic: topic,
			Key:   []byte(""),
			Value: dataJson,
		},
	)
	if err != nil {
		return model.FootballClub{}, err
	}

	return footballClub, nil
}
