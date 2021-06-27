package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(a *fiber.App) {

	app := a.Group("/api/v1")

	//sign up and login routes
	app.Post("/create-user", createUser)
	app.Post("/login", login)

	//item routes
	app.Post("/add-item", addItem)
	app.Put("/update-item/:id", updateItem)
	app.Get("/list-item", listItem)
	app.Delete("/delete-item/:id", deleteItem)

	//customer master routes
	app.Post("/add-customer", addCustomer)
	app.Put("/update-customer/:id", updateCustomer)
	app.Get("/list-customer", listCustomer)
	app.Delete("/delete-customer/:id", deleteCustomer)

	//
}
