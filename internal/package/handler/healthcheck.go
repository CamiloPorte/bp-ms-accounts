package handlers

import (
	"bp-ms-accounts/domain/entities"
	"bp-ms-accounts/internal/package/service"
	"net/http"
)

type accounts struct {
	configs map[string]string
}

func NewService(configs map[string]string) service.Service { return &accounts{configs: configs} }

func (a *accounts) HealthCheck(w http.ResponseWriter, r *http.Request) {
	entities.AccountsAnswer(http.StatusAccepted, "Connection healthy", w)
}
