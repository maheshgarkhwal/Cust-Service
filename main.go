package main

import (
	"cust-service/database"
	"cust-service/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New()
	routes.SetupRoutes(app)
	if err := app.Listen(os.Getenv(":3000")); err != nil {
		log.Fatal(err)
	}
}
