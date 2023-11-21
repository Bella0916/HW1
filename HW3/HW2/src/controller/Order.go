package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type Order struct {
	Id          string `json:"Order_id"`
	Customer_id string `json:"customer_id"`
	Is_shipped  string `json:"is_shipped"`
}

var orders []*Order

func GetOrders(c *gin.Context) {
	id := c.Query("id")
	db := connectDB()
	if id != "" {
		log.Println("id: " + id)
		db.Preload(clause.Associations).Where("id = $1", id).Find(&orders)
	} else {
		db.Preload(clause.Associations).Find(&orders)
	}
	closeDB(db)
	c.JSON(200, orders)
}

func GetOrderByOrd_ID(c *gin.Context) {
	db := connectDB()
	db.Where("id = $1", c.Param("ord_id")).Find(&orders)

	closeDB(db)
	c.JSON(200, orders)
}

func GetOrderByCus_ID(c *gin.Context) {
	db := connectDB()
	var customer *Customer
	db.Where("id = $1", c.Param("cus_id")).Find(&customer)
	var Cus_name string = customer.Name
	var order []*Order
	db.Where("customer_id = $1", c.Param("cus_id")).Find(&order)
	//var Odr_id[] string = order.Id

	closeDB(db)
	c.JSON(200, "customer_NAME: "+Cus_name)
	c.JSON(200, order)
}
func GetOrderByCus_Name(c *gin.Context) {
	db := connectDB()
	var customer *Customer
	db.Where("name = $1", c.Param("cus_name")).Find(&customer)
	var Cus_id string = customer.Id
	var order []*Order
	db.Where("customer_id = $1", Cus_id).Find(&order)
	//var Odr_id[] string = order.Id

	closeDB(db)
	c.JSON(200, customer)
	c.JSON(200, order)
}
