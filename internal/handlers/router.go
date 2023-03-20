package handlers

import (
	"eSearcher/configs"
	"eSearcher/internal/middlewares"
	"eSearcher/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
	Services *service.Services
	Config   *configs.Config
}

func NewRouter(cfg *configs.Config, services *service.Services) *Router {
	router := &Router{
		Router:   mux.NewRouter(),
		Services: services,
		Config:   cfg,
	}
	router.InitRoutes()
	return router
}

func (r *Router) InitRoutes() {
	r.Router.Use(middlewares.RateLimitMiddleWare(r.Config))
	r.Router.HandleFunc("/api/user/register", r.registration).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/user/login", r.authentication).Methods(http.MethodPost)

	// Data for applicant and employer pages
	r.Router.HandleFunc("/api/options", r.GetAllOptions).Methods(http.MethodGet)

	// Vacancy handlers
	r.Router.HandleFunc("/api/vacancy/", r.CreateVacancy).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/vacancy/search", r.KeyWordSearchVacancy).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/vacancy/my", r.GetMyVacancies).Methods(http.MethodGet)
	r.Router.HandleFunc("/api/vacancy/", r.UpdateVacancy).Methods(http.MethodPatch)

	// Applicant handlers
	r.Router.HandleFunc("/api/applicant/{id}", r.GetApplicant).Methods(http.MethodGet)
	r.Router.HandleFunc("/api/applicant/", r.CreateApplicant).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/applicant/", r.UpdateApplicant).Methods(http.MethodPatch)
	r.Router.HandleFunc("/api/applicant/search", r.SearchApplicant).Methods(http.MethodPost)

	// Employer handlers
	r.Router.HandleFunc("/api/employer/{id}", r.GetEmployer).Methods(http.MethodGet)
	r.Router.HandleFunc("/api/employer/", r.CreateEmployer).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/employer/", r.UpdateEmployer).Methods(http.MethodPatch)

	// Response handlers
	r.Router.HandleFunc("/api/response/my", r.GetMyResponses).Methods(http.MethodGet)
	r.Router.HandleFunc("/api/response/", r.AddResponse).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/response/", r.ChangeStatus).Methods(http.MethodPatch)
	r.Router.HandleFunc("/api/response/", r.DeleteResponse).Methods(http.MethodDelete)
	r.Router.HandleFunc("/api/response/respondents/{vacancyID}", r.GetRespondents).Methods(http.MethodGet)
}
