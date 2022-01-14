package usecase

import (
	"github.com/stretchr/testify/mock"
	model "github.com/yudhapratama10/search-publisher/model"
)

type FootballUsecaseMock struct {
	mock.Mock
}

func (f *FootballUsecaseMock) Insert(footballClub model.FootballClub) (model.FootballClub, error) {
	args := f.Called(footballClub)

	return args.Get(0).(model.FootballClub), args.Error(1)
}

func (f *FootballUsecaseMock) Update(footballClub model.FootballClub) (model.FootballClub, error) {
	args := f.Called(footballClub)

	return args.Get(0).(model.FootballClub), args.Error(1)
}
