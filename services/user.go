package service

import (
	"cust-service/database"
	"cust-service/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func RegisterationService(user *model.User) (map[string]interface{}, error) {
	db := database.DBConn
	password := []byte(user.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Print(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Print(err)
	}
	delete(m, "password")
	return m, nil
}

func LoginService(userData *model.User) string {

	db := database.DBConn
	var user model.User
	db.Find(&user, "Email = ?", userData.UserName)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))

	if err != nil {
		fmt.Print("error in compareHash", err)
		return err.Error()
	} else {
		token, err := CreateToken(user.UserName)
		if err != nil {
			fmt.Print(err)
		}
		return token
	}
}

func CreateToken(userId string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Authentication(c *fiber.Ctx) {
	tokenString := c.Get("Authorization")
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
		c.Next()
	} else {
		log.Println("Invalid JWT Token")
		c.Status(400).JSON("invalid json token")
	}
}
