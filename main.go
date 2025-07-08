package main

import (
	"hris_backend/database"
	"hris_backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	database.MigrationAll()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
