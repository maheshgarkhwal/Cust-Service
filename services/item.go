package service

import (
	"cust-service/database"
	"cust-service/model"
	"fmt"
	"strconv"
)

func CreateItem(i model.Item) (model.Item, error) {
	db := database.DBConn

	if err := db.Create(&i).Error; err != nil {
		return i, err
	}
	return i, nil
}

func ItemUpdate(i model.Item, id string) (model.Item, error) {
	db := database.DBConn
	var item model.Item
	db.First(&item, id)
	if item.ItemName == "" {
		return item, nil
	} else {
		if err := db.Model(&item).Updates(model.Item{ItemName: i.ItemName, Rate: i.Qty, Qty: i.Rate}).Error; err != nil {
			return item, err
		}
		return item, nil
	}
}

//item list
func ItemList(pg string, limit string) ([]model.Item, string) {
	db := database.DBConn
	var items []model.Item
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
		db.Select("id", "item_name", "rate", "qty").Limit(RecordLimit).Offset(offset).Find(&items)
		return items, ""
	}
}

//delete item

func ItemDelete(id string) model.Item {
	db := database.DBConn
	var item model.Item
	db.First(&item, id)
	if item.ItemName == "" {
		return item
	}
	db.Delete(&item)
	return item
}
