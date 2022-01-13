package pg

import (
	"context"

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

// ============================================================
// Internal Repository Mock

type mockDBConnection struct {
	mockQueryRow MockQueryRow
	mockQuery    MockQuery
	mockExec     MockExec
}

type mockRow struct {
	scan Scan
}

// each mockDBConnection will implement its own queryRow
type MockQueryRow func(ctx context.Context, sql string, args ...interface{}) pgx.Row

type MockQuery func(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)

type MockExec func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)

type Scan func(dest ...interface{}) error

func (mock mockRow) Scan(dest ...interface{}) error {
	return mock.scan(dest...)
}

func (mock mockDBConnection) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return mock.mockQueryRow(ctx, sql, args...)
}

func (mock mockDBConnection) Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error) {
	return mock.mockQuery(ctx, sql, optionsAndArgs...)
}

func (mock mockDBConnection) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return mock.mockExec(ctx, sql, arguments...)
}
