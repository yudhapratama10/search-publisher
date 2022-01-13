package kafka

import (
	"context"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	model "github.com/yudhapratama10/search-publisher/model"
)

func TestProduceMessage(t *testing.T) {

	t.Parallel()

	var data = model.FootballClub{
		Id:          25,
		Name:        "Newcastle United",
		Nation:      "Inggris",
		Tournaments: []string{"English Premier League", "FA Cup"},
		Rating:      2.5,
		Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
	}

	t.Run("Should Success Produce Insert", func(t *testing.T) {
		var (
			operation = "insert"
		)

		mockWriteMesasges := func(ctx context.Context, msgs ...kafka.Message) error {
			return nil
		}

		mock := mockKafkaConnection{mockWriteMessages: mockWriteMesasges}
		repo := NewFootballRepository(mock)

		res, err := repo.Produce(data, operation)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.NotZero(t, res.Id)
	})

	t.Run("Should Success Produce Update", func(t *testing.T) {
		var (
			operation = "update"
		)

		mockWriteMesasges := func(ctx context.Context, msgs ...kafka.Message) error {
			return nil
		}

		mock := mockKafkaConnection{mockWriteMessages: mockWriteMesasges}
		repo := NewFootballRepository(mock)

		res, err := repo.Produce(data, operation)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.NotZero(t, res.Id)
	})

	t.Run("Should Success Produce Delete", func(t *testing.T) {
		var (
			operation = "delete"
		)

		mockWriteMesasges := func(ctx context.Context, msgs ...kafka.Message) error {
			return nil
		}

		mock := mockKafkaConnection{mockWriteMessages: mockWriteMesasges}
		repo := NewFootballRepository(mock)

		res, err := repo.Produce(data, operation)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.NotZero(t, res.Id)
	})

}
