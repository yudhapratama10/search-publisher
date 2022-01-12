package main

import (
	"log"
	"net/http"

	"github.com/yudhapratama10/search-publisher/handler"
	kafkaClient "github.com/yudhapratama10/search-publisher/infrastructure/kafka"
	pgClient "github.com/yudhapratama10/search-publisher/infrastructure/pg"
	kafkaRepo "github.com/yudhapratama10/search-publisher/repository/kafka"
	pgRepo "github.com/yudhapratama10/search-publisher/repository/pg"
	"github.com/yudhapratama10/search-publisher/usecase"
)

func main() {

	dbClient, err := pgClient.InitClient()
	if err != nil {
		log.Fatal(err)
	}

	kClient := kafkaClient.InitClient()
	defer kClient.Close()

	// Initialize Repository, Usecase, & Handler
	pgRepo := pgRepo.NewFootballRepository(dbClient)
	kafkaRepo := kafkaRepo.NewFootballRepository(kClient)
	uc := usecase.NewFootballClubUsecase(pgRepo, kafkaRepo)
	handler := handler.NewFootballClubHandler(uc)

	http.HandleFunc("/insert", handler.Insert)
	http.HandleFunc("/update", handler.Update)

	log.Println("Starting Service")
	http.ListenAndServe(":8082", nil)
}
