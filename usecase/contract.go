package usecase

import (
	model "github.com/yudhapratama10/search-publisher/model"
	repoKafka "github.com/yudhapratama10/search-publisher/repository/kafka"
	repoPg "github.com/yudhapratama10/search-publisher/repository/pg"
)

type footballUsecase struct {
	repoPg    repoPg.FootballRepositoryContract
	repoKafka repoKafka.FootballRepositoryContract
}

type FootballUsecaseContract interface {
	Insert(footballClub model.FootballClub) (model.FootballClub, error)
	Update(footballClub model.FootballClub) (model.FootballClub, error)
}

func NewFootballClubUsecase(repoPg repoPg.FootballRepositoryContract, repoKafka repoKafka.FootballRepositoryContract) FootballUsecaseContract {
	return &footballUsecase{
		repoPg:    repoPg,
		repoKafka: repoKafka,
	}
}
