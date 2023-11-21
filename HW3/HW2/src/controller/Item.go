package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Item struct {
	Id         uuid.UUID `json:"Item_id"`
	Order_id   string    `json:"order_id"`
	Product_id string    `json:"product_id"`
	Is_shipped string    `json:"is_shipped"`
}

var items []*Item

func GetItems(c *gin.Context) {
	id := c.Query("id")
	db := connectDB()
	if id != "" {
		log.Println("id: " + id)
		db.Preload(clause.Associations).Where("id = $1", id).Find(&items)
	} else {
		db.Preload(clause.Associations).Find(&items)
	}
	closeDB(db)
	c.JSON(200, items)
}

func GetItemsById(c *gin.Context) {
	db := connectDB()
	db.Where("id = $1", c.Param("id")).Find(&items)

	closeDB(db)
	c.JSON(200, items)
}

func GetItemsByOrderId(c *gin.Context) {
	db := connectDB()
	defer closeDB(db)
	orderID := c.Param("order_id") // 取得訂單ID
	var items []Item
	db.Where("order_id = $1", orderID).Find(&items) // 查詢物品
	var order Order
	db.Where("id = $1", orderID).Find(&order) // 查詢訂單
	var customer Customer
	db.Where("id = $1", order.Customer_id).Find(&customer) // 查詢客戶

	response := gin.H{
		"items":    items,
		"order_id": order.Id,
		"customer": customer.Name,
	}
	c.JSON(200, response)
}

func GetItemsByProductId(c *gin.Context) {
	db := connectDB()
	defer closeDB(db)
	productID := c.Param("product_id") // 取得訂單ID
	var items []Item
	db.Where("product_id = $1", productID).Find(&items) // 查詢物品
	var order Order
	db.Where("id = $1", productID).Find(&order) // 查詢訂單

	response := gin.H{
		"items":      items,
		"product_id": productID,
	}
	c.JSON(200, response)

}
