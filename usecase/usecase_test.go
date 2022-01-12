package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudhapratama10/search-publisher/model"
	kafkaRepo "github.com/yudhapratama10/search-publisher/repository/kafka"
	pgRepo "github.com/yudhapratama10/search-publisher/repository/pg"
)

func TestInsert(t *testing.T) {
	t.Parallel()

	t.Run("Success Insert", func(t *testing.T) {
		pgRepo := new(pgRepo.FootballMock)
		kafkaRepo := new(kafkaRepo.FootaballMock)
		uc := NewFootballClubUsecase(pgRepo, kafkaRepo)

		var (
			// ctx          = context.Background()
			// id           = 30
			operation    = "insert"
			footballClub = model.FootballClub{
				Name:        "Newcastle United",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup"},
				Rating:      2.5,
				Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			}
		)

		// pgRepo.On("Get", id).Return([]model.FootballClub{}, nil)
		pgRepo.On("Insert", footballClub).Return(footballClub, operation, nil)
		kafkaRepo.On("Produce", footballClub, operation).Return(footballClub, nil)

		prod, err := uc.Insert(footballClub)
		assert.NoError(t, err)
		assert.NotEmpty(t, prod)
	})

	t.Run("Failed Insert Duplicate", func(t *testing.T) {
		pgRepo := new(pgRepo.FootballMock)
		kafkaRepo := new(kafkaRepo.FootaballMock)
		uc := NewFootballClubUsecase(pgRepo, kafkaRepo)

		var (
			// ctx          = context.Background()
			operation = "insert"
			// id           = 30
			footballClub = model.FootballClub{
				Name:        "Newcastle United",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup"},
				Rating:      2.5,
				Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			}
		)

		pgRepo.On("Insert", footballClub).Return(model.FootballClub{}, operation, errors.New("Duplicate on inserting data"))
		// kafkaRepo.On("Produce", footballClub, operation).Return(footballClub, nil)

		prod, err := uc.Insert(footballClub)
		fmt.Println(err)
		assert.Error(t, err)
		assert.Empty(t, prod)
	})

}
