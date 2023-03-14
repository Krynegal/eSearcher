package handlers

import (
	"eSearcher/internal/middlewares"
	"eSearcher/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
	Services *service.Services
	//RateLimiter redis.RateLimiter
}

func NewRouter(services *service.Services) *Router {
	router := &Router{
		Router:   mux.NewRouter(),
		Services: services,
		//RateLimiter: rateLimiter,
	}
	router.InitRoutes()
	return router
}

func (r *Router) InitRoutes() {
	r.Router.HandleFunc("/api/user/register", r.registration).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/user/login", r.authentication).Methods(http.MethodPost)

	//r.Router.HandleFunc("/api/vacancy/create", r.CreateVacancy).Methods(http.MethodPost)
	r.Router.Handle("/api/vacancy/create", middlewares.AuthMiddleware(r.CreateVacancy)).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/vacancy/search", r.KeyWordSearchVacancy).Methods(http.MethodPost)

	r.Router.HandleFunc("/api/applicant/{id}", r.GetApplicant).Methods(http.MethodGet)
	r.Router.HandleFunc("/api/applicant/create", r.CreateApplicant).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/applicant/search", r.SearchApplicant).Methods(http.MethodPost)

	r.Router.HandleFunc("/api/specializations", r.GetAllSpecializations).Methods(http.MethodGet)

	r.Router.HandleFunc("/api/response/add", r.AddResponse).Methods(http.MethodPost)
	r.Router.HandleFunc("/api/response/delete", r.DeleteResponse).Methods(http.MethodPost)
}
