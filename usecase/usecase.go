package usecase

import model "github.com/yudhapratama10/search-publisher/model"

func (usecase *footballUsecase) Insert(footballClub model.FootballClub) (model.FootballClub, error) {

	cursor, err := usecase.repo.Insert(footballClub)
	if err != nil {
		return model.FootballClub{}, err
	}

	return cursor, nil
}
