package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	var err error
	dsn := "mahesh:Mahesh@g7@tcp(localhost:3306)/crud?parseTime=true"
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect", err)
	} else {
		fmt.Println("connected with database")
	}
}
