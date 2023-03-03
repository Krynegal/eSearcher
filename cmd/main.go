package main

import (
	"eSearcher/configs"
	"eSearcher/internal/handlers"
	"eSearcher/internal/service"
	"eSearcher/internal/storage"
	"eSearcher/internal/storage/mongodb"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := configs.NewConfig()

	mongo, err := mongodb.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	store := &storage.Storage{
		VacancyStorage:  mongodb.NewVacancyCollection(mongo, "vacancies"),
		EmployeeStorage: nil,
		EmployerStorage: nil,
	}
	svc := &service.Services{
		VacancyService:  service.NewVacancies(store.VacancyStorage),
		EmployeeService: nil,
		EmployerService: nil,
	}

	r := handlers.NewRouter(cfg, svc)
	serverURL := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("server run on: %s", serverURL)
	log.Fatal(http.ListenAndServe(serverURL, r))
}
