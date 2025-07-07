package main

import (
	"hris_backend/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	database.MigrationAll()

	log.Fatal(app.Listen(":8080"))
}
