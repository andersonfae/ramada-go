package main

import (
	"api-produtos/database"
	"api-produtos/models"
	"api-produtos/routes"
	"log"
)

func main() {
	database.Connect()

	database.DB.AutoMigrate(&models.Produto{})

	r := routes.SetupRoutes()
	log.Fatal(r.Run(":8080"))
}
