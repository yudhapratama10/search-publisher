package usecase

import (
	"errors"
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
		// fmt.Println(err)
		assert.Error(t, err)
		assert.Empty(t, prod)
	})

}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("Success Update", func(t *testing.T) {
		pgRepo := new(pgRepo.FootballMock)
		kafkaRepo := new(kafkaRepo.FootaballMock)
		uc := NewFootballClubUsecase(pgRepo, kafkaRepo)

		var (
			operation    = "update"
			footballClub = model.FootballClub{
				Id:          25,
				Name:        "Newcastle United",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup"},
				Rating:      2.5,
				Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			}
		)

		pgRepo.On("Update", footballClub).Return(footballClub, operation, nil)
		pgRepo.On("Get", footballClub.Id).Return(footballClub, nil)
		kafkaRepo.On("Produce", footballClub, operation).Return(footballClub, nil)

		prod, err := uc.Update(footballClub)
		assert.NoError(t, err)
		assert.NotEmpty(t, prod)
	})

	t.Run("Failed Update Duplicate", func(t *testing.T) {
		pgRepo := new(pgRepo.FootballMock)
		kafkaRepo := new(kafkaRepo.FootaballMock)
		uc := NewFootballClubUsecase(pgRepo, kafkaRepo)

		var (
			operation    = "update"
			footballClub = model.FootballClub{
				Id:          0,
				Name:        "",
				Description: "",
			}
		)

		pgRepo.On("Get", footballClub.Id).Return(footballClub, nil)
		pgRepo.On("Update", footballClub).Return(model.FootballClub{}, operation, errors.New("Duplicate on updating data"))

		prod, err := uc.Update(footballClub)
		assert.Error(t, err)
		assert.Empty(t, prod)
	})

}

func TestDelete(t *testing.T) {
	t.Parallel()

	t.Run("Success Delete", func(t *testing.T) {
		pgRepo := new(pgRepo.FootballMock)
		kafkaRepo := new(kafkaRepo.FootaballMock)
		uc := NewFootballClubUsecase(pgRepo, kafkaRepo)

		var (
			operation    = "delete"
			footballClub = model.FootballClub{
				Id:          25,
				Name:        "Newcastle United",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup"},
				Rating:      2.5,
				Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			}
		)

		// pgRepo.On("Get", id).Return([]model.FootballClub{}, nil)
		pgRepo.On("Delete", footballClub).Return(footballClub, operation, nil)
		pgRepo.On("Get", footballClub.Id).Return(footballClub, nil)
		kafkaRepo.On("Produce", footballClub, operation).Return(footballClub, nil)

		prod, err := uc.Delete(footballClub)
		assert.NoError(t, err)
		assert.NotEmpty(t, prod)
	})

	t.Run("Failed Delete Duplicate", func(t *testing.T) {
		pgRepo := new(pgRepo.FootballMock)
		kafkaRepo := new(kafkaRepo.FootaballMock)
		uc := NewFootballClubUsecase(pgRepo, kafkaRepo)

		var (
			// operation    = "update"
			footballClub = model.FootballClub{
				Id:          0,
				Name:        "",
				Description: "",
			}
		)

		pgRepo.On("Get", footballClub.Id).Return(footballClub, nil)
		//pgRepo.On("Delete", footballClub).Return(model.FootballClub{}, operation, errors.New("Duplicate on updating data"))

		prod, err := uc.Update(footballClub)
		assert.Error(t, err)
		assert.Empty(t, prod)
	})

}
