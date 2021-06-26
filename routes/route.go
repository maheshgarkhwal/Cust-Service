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
	app.Post("/update-item/:id", updateItem)
	app.Post("/list-item", listItem)
	app.Post("/delete-item/:id", deleteItem)

	//customer master routes
	app.Post("/add-customer", addCustomer)
	app.Post("/update-customer/:id", updateCustomer)
	app.Post("/list-customer", listCustomer)
	app.Post("/delete-customer/:id", deleteCustomer)

}
