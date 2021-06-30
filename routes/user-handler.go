package routes

import (
	"cust-service/model"
	service "cust-service/services"
	"cust-service/validations"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

//authentication
func auth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if len(tokenString) == 0 {
		return c.Status(400).JSON("missing json token")
	}
	tokenString1 := strings.Split(tokenString, " ")

	hmacSecretString := os.Getenv("ACCESS_SECRET") // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenString1[1], func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		fmt.Println("error:", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return c.Next()
	} else {
		fmt.Println("Invalid JWT Token")
		return c.Status(400).JSON("invalid json token")
	}
}

type Person struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

//create user handler
func createUser(c *fiber.Ctx) error {
	u := new(Person)
	if err := c.BodyParser(u); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(err)

	}
	if u.Password != u.ConfirmPassword {
		return c.Status(400).JSON("password and confirm password must be same")
	}

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(err)

	}
	fmt.Print(user)
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
	result, err := service.LoginService(userData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "incorrect email password", "result": err.Error()})
	} else {
		return c.Status(200).JSON(fiber.Map{"message": "login sucess", "token": result})
	}
}

func rateCalculate(c *fiber.Ctx) error {
	orderData := new(model.Order)
	c.BodyParser(orderData)
	result := service.Calculate(orderData)

	if result != nil {
		return c.Status(200).JSON(result)
	} else {
		return c.Status(400).JSON("item name you enter does not exist")
	}
}

func placeOrder(c *fiber.Ctx) error {
	orderData := new(model.Order)
	c.BodyParser(orderData)
	result := service.OrderPlace(orderData)

	if result != nil {
		return c.Status(200).JSON(fiber.Map{"message": "order placed sucessfully", "result": result})
	} else {
		return c.Status(400).JSON("erro in placing order")
	}

}
