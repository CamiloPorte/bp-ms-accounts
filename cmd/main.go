package main

import (
	"bp-ms-accounts/cmd/conf"
	"bp-ms-accounts/domain/consts"
	"fmt"
	"log"
	"net/http"

	handlers "bp-ms-accounts/internal/package/handler"
	"bp-ms-accounts/internal/package/service"
)

func main() {
	configs, err := conf.GetConfigs()
	if err != nil {
		log.Fatalf("Error loading env values %e", err)
		return
	}

	services := handlers.NewService(configs)
	s := service.New(services)
	fmt.Println("servidor creado correctamente")

	log.Fatal(http.ListenAndServe(configs[consts.PortConst], s.Router()))
}
