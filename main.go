package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yudhapratama10/search-publisher/handler"
	kafkaClient "github.com/yudhapratama10/search-publisher/infrastructure/kafka"
	pgClient "github.com/yudhapratama10/search-publisher/infrastructure/pg"
	"github.com/yudhapratama10/search-publisher/repository"
	"github.com/yudhapratama10/search-publisher/usecase"
)

func main() {

	dbClient, err := pgClient.InitClient()
	if err != nil {
		log.Fatal(err)
	}

	kClient := kafkaClient.InitClient()
	defer kClient.Close()

	repo := repository.NewFootballRepository(dbClient, kClient)
	uc := usecase.NewFootballClubUsecase(repo)
	handler := handler.NewFootballClubHandler(uc)

	http.HandleFunc("/insert", handler.Insert)

	fmt.Println("Starting Service")
	http.ListenAndServe(":8082", nil)
}
