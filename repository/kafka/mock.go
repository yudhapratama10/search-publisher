package kafka

import (
	"github.com/stretchr/testify/mock"
	model "github.com/yudhapratama10/search-publisher/model"
)

type FootaballMock struct {
	mock.Mock
}

func (f *FootaballMock) Produce(footballClub model.FootballClub, operation string) (model.FootballClub, error) {
	args := f.Called(footballClub, operation)

	return args.Get(0).(model.FootballClub), args.Error(1)
}
