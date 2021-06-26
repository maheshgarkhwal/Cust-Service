package routes

import (
	"cust-service/model"
	service "cust-service/services"
	"cust-service/validations"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//add Custmer handler
func addCustomer(c *fiber.Ctx) error {
	Customer := new(model.Customer)
	if err := c.BodyParser(Customer); err != nil {
		fmt.Println(err)
		return err
	}
	errs := validations.ValidCustomer(*Customer)
	if len(errs) > 0 {
		return c.Status(401).JSON(fiber.Map{"message": "enter valid input", "result": errs})
	}
	result, err := service.CreateCustomer(*Customer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "unable to create Customer", "result": err})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Customer added sucessfully", "result": result})

}

//update Customer
func updateCustomer(c *fiber.Ctx) error {
	CustomerData := new(model.Customer)
	c.BodyParser(CustomerData)
	errs := validations.ValidCustomer(*CustomerData)
	if len(errs) > 0 {
		return c.Status(401).JSON(fiber.Map{"message": "enter valid input", "result": errs})
	}
	id := c.Params("id")
	result, err := service.CustomerUpdate(*CustomerData, id)
	if result.FirstName == "" {
		return c.Status(400).JSON("no Customer exist for the given id")
	} else {
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "unable to update Customer", "result": err})
		}
		return c.Status(200).JSON(fiber.Map{"message": "Customer updated sucessfully", "result": result})

	}
}

//list Customer
func listCustomer(c *fiber.Ctx) error {
	pg := c.Query("page")
	limit := c.Query("limit")
	result, message := service.CustomerList(pg, limit)
	if message != "" {
		return c.JSON(fiber.Map{"status": "failed", "result": message})
	} else {
		if len(result) < 1 {
			return c.JSON(fiber.Map{"status": "sucess", "result": "no more record left"})
		} else {
			return c.JSON(fiber.Map{"status": "sucess", "result": result})
		}
	}
}

func deleteCustomer(c *fiber.Ctx) error {

	id := c.Params("id")
	result := service.CustomerDelete(id)
	if result.FirstName == "" {
		return c.Status(400).JSON("Customer does not exist with the given id")
	}
	return c.Status(200).JSON(fiber.Map{"message": "Customer deleted sucessfully", "result": result})
}
