package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(a *fiber.App) {

	app := a.Group("/api/v1")
	app.Post("/create-user", CreateUser)
	app.Post("/login", Login)
}
