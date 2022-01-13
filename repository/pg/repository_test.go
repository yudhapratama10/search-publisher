package pg

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	model "github.com/yudhapratama10/search-publisher/model"
)

func TestInsertFootbalClub(t *testing.T) {

	t.Parallel()

	t.Run("Should Success Insert", func(t *testing.T) {
		data := model.FootballClub{
			Name:        "Newcastle United",
			Nation:      "Inggris",
			Tournaments: []string{"English Premier League", "FA Cup"},
			Rating:      2.5,
			Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
		}

		mockScanReturn := func(dest ...interface{}) error {
			id := dest[0].(*int) // scan pointer
			*id = data.Id

			return nil
		}

		mockQueryRowReturn := func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return mockRow{scan: mockScanReturn}
		}

		mock := mockDBConnection{mockQueryRow: mockQueryRowReturn}
		repo := NewFootballRepository(mock)

		res, _, err := repo.Insert(data)
		fmt.Println(res)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("Should Success Insert 2", func(t *testing.T) {
		data := model.FootballClub{
			Name:        "FC Barcelona",
			Nation:      "Spanyol",
			Tournaments: []string{"Laliga", "Copa Del Rey", "Uefa Europa League"},
			Rating:      4.5,
			Description: `Fútbol Club Barcelona, juga dikenal sebagai Barcelona atau Barça, adalah klub sepak bola profesional yang berbasis di Barcelona, Catalunya, Spanyol.Didirikan pada tahun 1899 oleh sekelompok pemain Swiss, Inggris, Jerman dan Katalan yang dipimpin oleh Joan Gamper, klub telah menjadi simbol budaya Catalan dan Catalanisme, yang mempunyai motto "Més que un club" (Lebih dari sebuah klub).`,
		}

		mockScanReturn := func(dest ...interface{}) error {
			id := dest[0].(*int) // scan pointer
			*id = data.Id

			return nil
		}

		mockQueryRowReturn := func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return mockRow{scan: mockScanReturn}
		}

		mock := mockDBConnection{mockQueryRow: mockQueryRowReturn}
		repo := NewFootballRepository(mock)

		res, _, err := repo.Insert(data)
		// fmt.Println(res)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
}

func TestUpdateFootballClub(t *testing.T) {

	t.Parallel()

	t.Run("Should Success Update", func(t *testing.T) {

		data := model.FootballClub{
			Id:          50,
			Name:        "Newcastle United",
			Nation:      "Inggris",
			Tournaments: []string{"English Premier League", "FA Cup"},
			Rating:      2.5,
			Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
		}

		// Mock Response
		mockExec := func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
			if len(arguments) != 7 {
				return nil, errors.New("Supposed to be 7 params")
			}
			return nil, nil
		}

		mock := mockDBConnection{mockExec: mockExec}
		repo := NewFootballRepository(mock)

		res, _, err := repo.Update(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
}

func TestDeleteFootballClub(t *testing.T) {

	t.Parallel()

	t.Run("Should Success Delete", func(t *testing.T) {
		data := model.FootballClub{
			Id:          20,
			Name:        "Newcastle United",
			Nation:      "Inggris",
			Tournaments: []string{"English Premier League", "FA Cup"},
			Rating:      2.5,
			Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
		}

		mockExec := func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
			if len(arguments) != 1 {
				return nil, errors.New("Supposed to be 1 params")
			}

			if arguments[0] == nil || arguments[0] == 0 {
				return nil, errors.New("Args `id` cannot be empty")
			}
			return nil, nil
		}

		mock := mockDBConnection{mockExec: mockExec}
		repo := NewFootballRepository(mock)

		res, _, err := repo.Delete(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
}
