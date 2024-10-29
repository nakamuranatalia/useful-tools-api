package main

import (
	"github.com/nakamuranatalia/useful-tools-api/internal/database"
	"github.com/nakamuranatalia/useful-tools-api/internal/routes"
)

func main() {
	database.DatabaseConnection()
	routes.HandleRequest()

}
