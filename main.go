package main

import (
	"hris_backend/database"
	"hris_backend/internal/seeders"
	"hris_backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()
	database.MigrationAll()

	db := database.DB
	actorID := uint(1)
	err := seeders.SeedPermissions(db, actorID)
	if err != nil {
		panic(err)
	}

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
