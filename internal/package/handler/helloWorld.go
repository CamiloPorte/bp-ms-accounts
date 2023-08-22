package handler

import (
	"bp-ms-accounts/domain/entities"
	"bp-ms-accounts/internal/package/service"
	"net/http"
)

type helloWorld struct {
	configs map[string]string
}

func NewHelloWorldService(configs map[string]string) service.Service {
	return &helloWorld{configs: configs}
}

func (h *helloWorld) Resolver(w http.ResponseWriter, r *http.Request) {
	entities.AccountsAnswer(http.StatusAccepted, "Connection healthy", w)
}
