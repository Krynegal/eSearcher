package main

import (
	"eSearcher/configs"
	"eSearcher/internal/handlers"
	"eSearcher/internal/service"
	"eSearcher/internal/storage"
	"eSearcher/internal/storage/mongodb"
	"eSearcher/internal/storage/postgres"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := configs.NewConfig()

	dbpool, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	mongo, err := mongodb.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	store := &storage.Storage{
		VacancyStorage:        mongodb.NewVacancyCollection(mongo, "vacancies"),
		ApplicantStorage:      postgres.NewApplicantsStore(dbpool),
		SpecializationStorage: postgres.NewSpecializationsStore(dbpool),
		EmployerStorage:       postgres.NewEmployersStore(dbpool),
		ResponsesStorage:      postgres.NewResponsesStore(dbpool),
	}
	svc := &service.Services{
		VacancyService:         service.NewVacancies(store.VacancyStorage),
		ApplicantsService:      service.NewApplicants(store.ApplicantStorage),
		SpecializationsService: service.NewSpecializations(store.SpecializationStorage),
		EmployersService:       service.NewEmployers(store.EmployerStorage),
		ResponsesService:       service.NewResponses(store.ResponsesStorage),
	}

	//rateLimiter, err := redis.New(cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	r := handlers.NewRouter(cfg, svc)
	serverURL := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("server run on: %s", serverURL)
	log.Fatal(http.ListenAndServe(serverURL, r))
}
