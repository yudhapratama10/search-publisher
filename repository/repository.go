package repository

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/segmentio/kafka-go"
	model "github.com/yudhapratama10/search-publisher/model"
)

var topic string = "test-messages"

func (repo *footballRepository) Insert(footballClub model.FootballClub) (model.FootballClub, error) {
	var id int

	err := repo.db.QueryRow(context.Background(),
		"INSERT INTO footballclub (name, tournaments, nation, hasstadium, description, rating) VALUES ($1, $2, $3, $4, $5, $6) returning id",
		footballClub.Name, footballClub.Tournaments, footballClub.Nation, footballClub.HasStadium, footballClub.Description, footballClub.Rating).Scan(&id)
	if err != nil {
		return model.FootballClub{}, err
	}

	footballClub.Id = id

	dataJson, err := json.Marshal(map[string]interface{}{
		"data":      footballClub,
		"operation": "insert",
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

func (repo *footballRepository) Update(footballClub model.FootballClub) (model.FootballClub, error) {
	_, err := repo.db.Exec(context.Background(),
		"UPDATE footballclub set name = $1, tournaments = $2, nation = $3, hasstadium = $4, description = $5, rating = $6 where id = $7",
		footballClub.Name, footballClub.Tournaments, footballClub.Nation, footballClub.HasStadium, footballClub.Description, footballClub.Rating, footballClub.Id)
	if err != nil {
		return model.FootballClub{}, err
	}

	dataJson, err := json.Marshal(map[string]interface{}{
		"data":      footballClub,
		"operation": "update",
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

func (repo *footballRepository) Delete(footballClub model.FootballClub) (model.FootballClub, error) {
	_, err := repo.db.Exec(context.Background(),
		"DELETE from footballclub where id = $1", footballClub.Id)
	if err != nil {
		return model.FootballClub{}, err
	}

	dataJson, err := json.Marshal(map[string]interface{}{
		"data":      footballClub,
		"operation": "delete",
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

func (repo *footballRepository) Get(id int) (model.FootballClub, error) {

	rows, err := repo.db.Query(context.Background(),
		"SELECT id, name, tournaments, nation, hasstadium, description, rating from footballclub where id = $1",
		id)
	if err != nil {
		return model.FootballClub{}, err
	}
	defer rows.Close()

	var data model.FootballClub

	if rows.Next() {
		err := rows.Scan(&data.Id, &data.Name, &data.Tournaments, &data.Nation, &data.HasStadium, &data.Description, &data.Rating)
		if err != nil {
			return model.FootballClub{}, err
		}
	}

	return data, nil
}
