package handler

import (
	"bp-ms-accounts/domain/entities"
	"bp-ms-accounts/internal/package/service"
	"net/http"

	"github.com/labstack/gommon/log"
)

type helloWorld struct {
	configs map[string]string
}

func NewHelloWorldService(configs map[string]string) service.Service {
	return &helloWorld{configs: configs}
}

func (h *helloWorld) Resolver(w http.ResponseWriter, r *http.Request) {
	log.Info("se solicito saludo")
	entities.AccountsAnswer(http.StatusAccepted, "hello world", w)
}
