package controller

import (
	//"fmt"
	//"go_gin_example/envconfig"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	//"github.com/google/uuid"

	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	Id   uuid.UUID `json:"Category_id"`
	Name string    `json:"name"`
}

var categories []*Category

func GetCategories(c *gin.Context) {
	name := c.Query("name")
	db := connectDB()
	if name != "" {
		log.Println("name: " + name)
		db.Preload(clause.Associations).Where("name = $1", name).Find(&categories)
	} else {
		db.Preload(clause.Associations).Find(&categories)
	}
	closeDB(db)
	c.JSON(200, categories)
}

func GetCategories_Name(c *gin.Context) {
	db := connectDB()
	db.Where("name = $1", c.Param("crg_name")).Find(&categories)

	closeDB(db)
	c.JSON(200, categories)
}

func GetCategories_ID(c *gin.Context) {
	db := connectDB()
	db.Where("id = $1", c.Param("crg_id")).Find(&categories)

	closeDB(db)
	c.JSON(200, categories)
}
