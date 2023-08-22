package main

import (
	"bp-ms-accounts/cmd/conf"
	"bp-ms-accounts/domain/consts"
	"net/http"

	"github.com/labstack/gommon/log"

	"bp-ms-accounts/internal/package/handler"
	"bp-ms-accounts/internal/package/service"
)

func main() {
	configs, err := conf.GetConfigs()
	if err != nil {
		log.Fatalf("Error loading env values %e", err)
		return
	}

	services := map[string]service.Service{
		consts.Healthcheck: handler.NewAccountsService(configs),
		consts.HelloWorld:  handler.NewHelloWorldService(configs),
	}

	s := service.New(services)

	log.Info("servidor creado correctamente")

	log.Fatal(http.ListenAndServe(configs[consts.PortConst], s.Router()))
}
