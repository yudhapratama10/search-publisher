package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudhapratama10/search-publisher/model"
	"github.com/yudhapratama10/search-publisher/usecase"
)

func TestInsert(t *testing.T) {
	t.Parallel()

	t.Run("Success Insert", func(t *testing.T) {
		mock := new(usecase.FootballUsecaseMock)
		handlers := NewFootballClubHandler(mock)

		var (
			footballClub = model.FootballClub{
				Name:        "Newcastle United",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup"},
				Rating:      2.5,
				Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			}
		)

		mock.On("Insert", footballClub).Return(footballClub, nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.Insert)

		reqByte, _ := json.Marshal(footballClub)
		r := io.NopCloser(strings.NewReader(string(reqByte)))

		req, err := http.NewRequest("POST", "/insert", r)
		if err != nil {
			t.Fatal(err)
		}

		handler.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})

}
