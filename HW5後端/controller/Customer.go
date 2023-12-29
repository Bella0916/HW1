package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type Customer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetCustomers(c *gin.Context) {
	db := connectDB()
	var customers []*Customer
	db.Preload(clause.Associations).Find(&customers)

	closeDB(db)
	c.JSON(200, customers)
}

func GetCustomersByID(c *gin.Context) {
	db := connectDB()
	var customer *Customer
	db.Where("id = $1", c.Param("CustomerId")).Take(&customer)

	closeDB(db)
	c.JSON(200, customer)
}

func CreateCustomers(c *gin.Context) {
	db := connectDB()
	var customer *Customer
	c.BindJSON(&customer)
	result := db.Create(&customer)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(500, gin.H{
			"message": "Create customer failed with error: " + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, customer)
}
