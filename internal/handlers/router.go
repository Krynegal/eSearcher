package handlers

import (
	"eSearcher/configs"
	"eSearcher/internal/service"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	Services   *service.Services
	GoogleAuth *GoogleAuth
}

func NewRouter(cfg *configs.Config, services *service.Services) *Router {
	router := &Router{
		Router:     mux.NewRouter(),
		Services:   services,
		GoogleAuth: NewGoogleAuth(cfg),
	}
	router.InitRoutes()
	return router
}

func (r *Router) InitRoutes() {
	r.Router.HandleFunc("/", r.HandleMain)
	r.Router.HandleFunc("/login-gl", r.HandleGoogleLogin)
	r.Router.HandleFunc("/callback-gl", r.CallBackFromGoogle)

	r.Router.HandleFunc("/api/vacancy", r.CreateVacancy)
	//r.Router.HandleFunc("/api/user/register", r.registration).Methods(http.MethodPost)
	//r.Router.HandleFunc("/api/user/login", r.authentication).Methods(http.MethodPost)
	//r.Router.Handle("/api/user/orders", middlewares.AuthMiddleware(r.loadOrders)).Methods(http.MethodPost)
}
