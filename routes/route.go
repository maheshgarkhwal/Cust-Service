package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	//sign up and login routes
	api.Post("/create-user", createUser)
	api.Post("/login", login)

	//item routes
	api.Post("/add-item", auth, addItem)
	api.Put("/update-item/:id", auth, updateItem)
	api.Get("/list-item", auth, listItem)
	api.Delete("/delete-item/:id", auth, deleteItem)

	//customer master routes
	api.Post("/add-customer", auth, addCustomer)
	api.Put("/update-customer/:id", auth, updateCustomer)
	api.Get("/list-customer", auth, listCustomer)
	api.Delete("/delete-customer/:id", auth, deleteCustomer)

	//
}
