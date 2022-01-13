package kafka

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"
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

// ============================================================
// Internal Repository Mock

type mockKafkaConnection struct {
	mockWriteMessages MockWriteMessages
}

type MockWriteMessages func(ctx context.Context, msgs ...kafka.Message) error

func (mock mockKafkaConnection) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	if len(msgs) < 1 {
		return errors.New("Should contain message")
	}

	return mock.mockWriteMessages(ctx, msgs...)
}
