package main

import (
	"cust-service/database"
	"cust-service/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	database.InitDatabase()
	app := fiber.New()
	app.Static("/", "./view")
	app.Use(cors.New())
	routes.SetupRoutes(app)
	if err := app.Listen(os.Getenv("PORT_NO")); err != nil {
		log.Fatal(err)
	}
}
