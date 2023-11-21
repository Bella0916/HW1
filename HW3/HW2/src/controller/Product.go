package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Product struct {
	Id         uuid.UUID `json:"Product_id"`
	Name       string    `json:"name"`
	Price      string    `json:"price"`
	CategoryID string    `json:"category_id"`
}

func GetProducts(c *gin.Context) {
	name := c.Query("name")
	db := connectDB()
	var products []*Product
	if name != "" {
		log.Println("name: " + name)
		db.Preload(clause.Associations).Where("name = $1", name).Find(&products)
	} else {
		db.Preload(clause.Associations).Find(&products)
	}
	closeDB(db)
	c.JSON(200, products)
}
func GetProducts_CategoryID(c *gin.Context) {
	db := connectDB()
	defer closeDB(db)
	CategoryID := c.Param("ctg_id")
	var products []Product
	db.Where("category_id = $1", CategoryID).Find(&products)
	var categories []*Category
	db.Where("id = $1", CategoryID).Find(&categories)

	response := gin.H{
		"products":   products,
		"categories": categories,
	}
	c.JSON(200, response)
}
