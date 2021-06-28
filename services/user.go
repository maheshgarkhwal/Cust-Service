package service

import (
	"cust-service/database"
	"cust-service/model"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	user.ConfirmPassword = string(hashedPassword)
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
	delete(m, "confirm_password")
	return m, nil
}

func LoginService(userData *model.User) string {

	db := database.DBConn
	var user model.User
	db.Find(&user, "user_name = ?", userData.Email)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		return err.Error()
	} else {
		token, err := CreateToken(user.Email)
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

func Calculate(o *model.Order) *model.Order {
	db := database.DBConn
	var item model.Item
	db.Find(&item, "item_name = ?", o.Item)
	if item.ItemName == "" {
		return nil
	}
	o.Rate = item.Rate * o.Qty
	return o
}

func OrderPlace(o *model.Order) *model.Order {
	db := database.DBConn
	var item model.Item
	db.Find(&item, "item_name = ?", o.Item)
	if item.ItemName == "" {
		return nil
	}
	item.Qty = item.Qty - o.Qty
	if err := db.Model(&item).Updates(model.Item{ItemName: item.ItemName, Rate: item.Rate, Qty: item.Qty}).Error; err != nil {
		return nil
	}
	if err := db.Create(&o).Error; err != nil {
		return nil
	}
	return o
}
