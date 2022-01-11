package usecase

import (
	"errors"

	model "github.com/yudhapratama10/search-publisher/model"
)

func (usecase *footballUsecase) Insert(footballClub model.FootballClub) (model.FootballClub, error) {

	cursor, action, err := usecase.repoPg.Insert(footballClub)
	if err != nil {
		return model.FootballClub{}, err
	}

	// Produce message to Kafka
	_, err = usecase.repoKafka.Produce(footballClub, action)
	if err != nil {
		return model.FootballClub{}, err
	}

	return cursor, nil
}

func (usecase *footballUsecase) Update(footballClub model.FootballClub) (model.FootballClub, error) {

	// Checking Exist
	respGet, err := usecase.repoPg.Get(footballClub.Id)
	if err != nil {
		return model.FootballClub{}, err
	}

	if respGet.Id == 0 {
		return model.FootballClub{}, errors.New("Data not found")
	}

	// Update
	respUpdate, action, err := usecase.repoPg.Update(footballClub)
	if err != nil {
		return model.FootballClub{}, err
	}

	// Produce message to Kafka
	_, err = usecase.repoKafka.Produce(footballClub, action)
	if err != nil {
		return model.FootballClub{}, err
	}

	return respUpdate, nil
}
