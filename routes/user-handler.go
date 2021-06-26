package routes

import (
	"cust-service/model"
	service "cust-service/services"
	"cust-service/validations"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//create user handler
func createUser(c *fiber.Ctx) error {

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(err)
	}
	errs := validations.ValidUser(*user)
	if len(errs) > 0 {
		return c.Status(401).JSON(fiber.Map{"message": "enter valid input", "result": errs})
	}
	result, err := service.RegisterationService(user)

	//fmt.Printf("%T", result)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "unable to create user", "result": err})
	} else {
		return c.Status(200).JSON(fiber.Map{"message": "user sucessfully created", "result": result})
	}
}

//login handler
func login(c *fiber.Ctx) error {

	userData := new(model.User)
	c.BodyParser(userData)
	result := service.LoginService(userData)
	return c.Status(200).JSON(result)
}
