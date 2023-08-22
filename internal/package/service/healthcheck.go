package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type api struct {
	router  http.Handler
	service Service
}

type Server interface {
	Router() http.Handler
}

func New(service Service) Server {
	a := &api{
		service: service,
	}

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", a.service.HealthCheck).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
