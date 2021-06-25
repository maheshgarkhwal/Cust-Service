package main

import (
	"cust-service/database"
	routes "cust-service/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New()
	routes.SetupRoutes(app)
	if err := app.Listen(os.Getenv("PORT_NO")); err != nil {
		log.Fatal(err)
	} else {
		log.Fatal("server started on port", os.Getenv("PORT_NO"))
	}
}
