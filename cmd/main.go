package main

import (
	c "github.com/nakamuranatalia/useful-tools-api/internal/controller"
	db "github.com/nakamuranatalia/useful-tools-api/internal/database"
	repo "github.com/nakamuranatalia/useful-tools-api/internal/repository"
	"github.com/nakamuranatalia/useful-tools-api/internal/routes"
	s "github.com/nakamuranatalia/useful-tools-api/internal/service"
)

func main() {
	database := db.DatabaseConnection()

	repository := repo.NewRepository(database)
	service := s.NewService(repository)
	controller := c.NewController(service)
	routes.HandleRequest(controller)
}
