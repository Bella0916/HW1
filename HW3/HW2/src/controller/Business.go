package controller

import (
	"fmt"
	"go_gin_example/envconfig"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 新增商品
func GetCreateProduct(c *gin.Context) {
	db := connectDB()
	var product Product
	c.BindJSON(&product)
	db.Create(&product)
	c.JSON(200, product)
	closeDB(db)
}

// 刪除商品
func GetDeleteProduct(c *gin.Context) {
	db := connectDB()
	id := c.Query("id")
	//var product Product
	db.Delete(clause.Associations).Where("id = $1", c.Param("products_id")).Find(&id)
	//fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// 新增客戶
func GetCreateCustomer(c *gin.Context) {
	db := connectDB()
	var customer Customer
	c.BindJSON(&customer)
	db.Create(&customer)
	c.JSON(200, customer)
}

// 客戶下訂單
func GetCreateOrder(c *gin.Context) {
	db := connectDB()
	var order Order
	c.BindJSON(&order)
	db.Create(&order)
	c.JSON(200, order)
}

// 更新訂單付款狀態
func GetUpdateOrderPaymentStatus(c *gin.Context) {
	id := c.Params.ByName("id")
	db := connectDB()
	var item Item
	if err := db.Where("id = $1", id).First(&item).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&item)
	db.Save(&item)
	c.JSON(200, item)
}

// 更新item送貨狀態
func GetUpdateItemShippingStatus(c *gin.Context) {
	db := connectDB()
	//id := c.Params.ByName("id")
	var item Item
	//c.BindJSON(&item)
	db.Save(&item)
	c.JSON(200, item)
}

func connectDB() *gorm.DB {
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Taipei",
		envconfig.GetEnv("DB_HOST"), "postgres", "postgres", "nutc_csic", envconfig.GetEnv("DB_PORT"), envconfig.GetEnv("DB_WITH_SSL"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}
