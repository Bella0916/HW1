package controller

import (
	//"fmt"
	//"go_gin_example/envconfig"
	"log"

	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"

	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Customer struct {
	Id   string `json:"Customer_id"`
	Name string `json:"Customer_name"`
}

var customers []*Customer

func GetCustomers(c *gin.Context) {
	id := c.Query("id")
	db := connectDB()
	if id != "" {
		log.Println("id: " + id)
		db.Preload(clause.Associations).Where("id = $1", id).Find(&customers)
	} else {
		db.Preload(clause.Associations).Find(&customers)
	}
	closeDB(db)
	c.JSON(200, customers)
}

func GetCustomersByCus_ID(c *gin.Context) {
	db := connectDB()
	db.Where("id = $1", c.Param("cus_id")).Find(&customers)

	closeDB(db)
	c.JSON(200, customers)
}

func GetCustomersByCus_Name(c *gin.Context) {
	db := connectDB()
	db.Where("name = $1", c.Param("cus_name")).Find(&customers)

	closeDB(db)
	c.JSON(200, customers)
}
