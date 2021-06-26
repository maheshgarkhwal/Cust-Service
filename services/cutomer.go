package service

import (
	"cust-service/database"
	"cust-service/model"
	"fmt"
	"strconv"
)

func CreateCustomer(i model.Customer) (model.Customer, error) {
	db := database.DBConn

	if err := db.Create(&i).Error; err != nil {
		return i, err
	}
	return i, nil
}

func CustomerUpdate(i model.Customer, id string) (model.Customer, error) {
	db := database.DBConn
	var Customer model.Customer
	db.First(&Customer, id)
	if Customer.FirstName == "" {
		return Customer, nil
	} else {
		if err := db.Model(&Customer).Updates(model.Customer{FirstName: Customer.FirstName, LastName: Customer.LastName, Phone: Customer.Phone}).Error; err != nil {
			return Customer, err
		}
		return Customer, nil
	}
}

//Customer list
func CustomerList(pg string, limit string) ([]model.Customer, string) {
	db := database.DBConn
	var Customers []model.Customer
	RecordLimit, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Print(err)
	}
	page, err := strconv.Atoi(pg)
	if err != nil {
		fmt.Print(err)
	}
	offset, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Print(err)
	}
	if page < 1 {
		return nil, "no page exist"
	} else {
		offset = (page - 1) * offset
		db.Select("id", "first_name", "last_name", "phone", "email").Limit(RecordLimit).Offset(offset).Find(&Customers)
		return Customers, ""
	}
}

//delete Customer

func CustomerDelete(id string) model.Customer {
	db := database.DBConn
	var Customer model.Customer
	db.First(&Customer, id)
	if Customer.FirstName == "" {
		return Customer
	}
	db.Delete(&Customer)
	return Customer
}
