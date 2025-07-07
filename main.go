package main

import (
	"hris_backend/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	if err := database.MigrateAll(database.DB); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Fatal(app.Listen(":3000"))
}
