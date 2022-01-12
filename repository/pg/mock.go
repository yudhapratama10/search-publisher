package pg

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/mock"
	"github.com/yudhapratama10/search-publisher/model"
)

// External Mock (Usecase)

type FootballMock struct {
	mock.Mock
}

func (f *FootballMock) Get(id int) (model.FootballClub, error) {
	args := f.Called(id)

	return args.Get(0).(model.FootballClub), args.Error(1)
}

func (f *FootballMock) Insert(footballClub model.FootballClub) (model.FootballClub, string, error) {
	args := f.Called(footballClub)

	return args.Get(0).(model.FootballClub), args.String(1), args.Error(2)
}

func (f *FootballMock) Update(footballClub model.FootballClub) (model.FootballClub, string, error) {
	args := f.Called(footballClub)

	return args.Get(0).(model.FootballClub), args.String(1), args.Error(2)
}

func (f *FootballMock) Delete(footballClub model.FootballClub) (model.FootballClub, string, error) {
	args := f.Called(footballClub)

	return args.Get(0).(model.FootballClub), args.String(1), args.Error(2)
}

// =============================================================================

// Internal Repository Mock

type mockDBConnection struct {
}

type mockRow struct {
	id  int
	msg string
}

func (mock mockRow) Scan(dest ...interface{}) error {
	id := dest[0].(*int) // scan pointer
	if mock.msg != "" {
		return errors.New(mock.msg)
	}
	*id = mock.id

	return nil
}

func (mock mockDBConnection) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if args[0] != nil && args[0] == "" {
		return mockRow{id: 0, msg: "Field is empty"}
	} else if args[1] != nil && args[1] == "" {
		return mockRow{id: 0, msg: "Field is empty"}
	} else if args[2] != nil && args[2] == "" {
		return mockRow{id: 0, msg: "Field is empty"}
	}

	rand.Seed(time.Now().UnixNano())
	return mockRow{id: rand.Intn(100-1+1) + 1}
}

func (mock mockDBConnection) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	if string(sql[0]) == "U" {
		if len(arguments) != 7 {
			return nil, errors.New("Supposed to be 7 params")
		}
	} else if string(sql[0]) == "D" {
		if len(arguments) != 1 {
			return nil, errors.New("Supposed to be 1 params")
		}
	}

	return nil, nil
}

func (mock mockDBConnection) Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error) {
	return nil, nil
}
