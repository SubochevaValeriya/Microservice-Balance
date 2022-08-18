package main

import (
	balance "github.com/SubochevaValeriya/microservice-balance"
	"github.com/SubochevaValeriya/microservice-balance/pkg/handler"
	"github.com/SubochevaValeriya/microservice-balance/pkg/repository"
	"github.com/SubochevaValeriya/microservice-balance/pkg/service"
	"log"
)

func main() {
	// dependency injection
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(balance.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ccured while running http server: %s", err.Error())
	}

}
