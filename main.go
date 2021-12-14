package main

import (
	"context"
	"log"
	"zerovwap/configuration"
	"zerovwap/service"
	"zerovwap/ws"
)

func main() {
	ctx := context.Background()

	config := configuration.InitConfig()

	wsClient, err := ws.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.NewService(wsClient, config)

	err = svc.Execute(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
