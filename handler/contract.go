package handler

import (
	"net/http"

	"github.com/yudhapratama10/search-publisher/usecase"
)

type FootballHandler struct {
	footballUsecase usecase.FootballUsecaseContract
}

type FootballHandlerContract interface {
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewFootballClubHandler(footballUsecase usecase.FootballUsecaseContract) FootballHandlerContract {
	return &FootballHandler{
		footballUsecase: footballUsecase,
	}
}
