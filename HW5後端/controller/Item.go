package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Id        int  `json:"id"`
	OrderID   int  `json:"order_id"`
	ProductID int  `json:"product_id"`
	IsShipped bool `json:"is_shipped"`
}

func GetOrderItems(c *gin.Context) {
	db := connectDB()
	var items []*Item
	db.Find(&items)

	closeDB(db)
	c.JSON(200, items)
}

func GetItemsByOrderID(c *gin.Context) {
	db := connectDB()
	var item *Item
	db.Where("order_id = $1", c.Param("OrderId")).Take(&item)

	closeDB(db)
	c.JSON(200, item)
}

func UpdateItemsByOrderID(c *gin.Context) {
	db := connectDB()
	var item *Item
	queryResult := db.Where("order_id = $1", c.Param("OrderId")).Take(&item)

	if queryResult.Error != nil {
		log.Println(queryResult.Error)
		c.JSON(500, gin.H{
			"message": "Update IsShipped failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	var itemBody *Item
	c.BindJSON(&itemBody)
	itemBody.OrderID = item.OrderID
	result := db.Model(&item).Where("order_id = ?", item.OrderID).Updates(itemBody)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(500, gin.H{
			"message": "Update item_IsShipped failed with error: " + result.Error.Error(),
		})

		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Update items Susseccfully",
	})
}
