package routes

import (
	"cust-service/model"
	service "cust-service/services"
	"cust-service/validations"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//add item handler
func addItem(c *fiber.Ctx) error {
	item := new(model.Item)
	if err := c.BodyParser(item); err != nil {
		fmt.Println(err)
		return err
	}
	errs := validations.ValidItem(*item)
	if len(errs) > 0 {
		return c.Status(401).JSON(fiber.Map{"message": "enter valid input", "result": errs})
	}
	fmt.Printf("%v,%T \n", item.ItemName, item.ItemName)
	fmt.Printf("%v,%T \n", item.Qty, item.Qty)
	fmt.Printf("%v,%T \n", item.Rate, item.Rate)

	result, err := service.CreateItem(*item)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "unable to create Item", "result": err})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Item added sucessfully", "result": result})

}

//update item
func updateItem(c *fiber.Ctx) error {
	itemData := new(model.Item)
	c.BodyParser(itemData)
	errs := validations.ValidItem(*itemData)
	if len(errs) > 0 {
		return c.Status(401).JSON(fiber.Map{"message": "enter valid input", "result": errs})
	}
	id := c.Params("id")
	result, err := service.ItemUpdate(*itemData, id)
	if result.ItemName == "" {
		return c.Status(400).JSON("no item exist for the given id")
	} else {
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "unable to update item", "result": err})
		}
		return c.Status(200).JSON(fiber.Map{"message": "item updated sucessfully", "result": result})

	}
}

//list item
func listItem(c *fiber.Ctx) error {
	pg := c.Query("page")
	limit := c.Query("limit")
	result, message := service.ItemList(pg, limit)
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

func deleteItem(c *fiber.Ctx) error {

	id := c.Params("id")
	result := service.ItemDelete(id)
	if result.ItemName == "" {
		return c.Status(400).JSON("item does not exist with the given id")
	}
	return c.Status(200).JSON(fiber.Map{"message": "item deleted sucessfully", "result": result})
}
