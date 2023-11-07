package controller

import (
	//"go_gin_example/model"

	"fmt"
	"go_gin_example/envconfig"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

type Category struct {
	Id   uuid.UUID `json:"Category_id"`
	Name string    `json:"name"`
}

func GetCategories(c *gin.Context) {
	name := c.Query("name")
	db := connectDB()
	var categories []*Category
	if name != "" {
		log.Println("name: " + name)
		db.Preload(clause.Associations).Where("name = $1", name).Find(&categories)
	} else {
		db.Preload(clause.Associations).Find(&categories)
	}
	closeDB(db)
	c.JSON(200, categories)
}

type Customer struct {
	Id string `json:"Customer_id"`
	//Cus_id string    `json:"Customer_id"`
	Name string `json:"Customer_name"`
	// Department represents the department of a student.
	//**Department *Department `json:"department,omitempty"`
	//**Courses    []Course    `gorm:"many2many:student_course;"`
}

func GetCustomers(c *gin.Context) {
	id := c.Query("id")
	db := connectDB()
	var customers []*Customer
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
	var customers []*Customer
	db.Where("id = $1", c.Param("cus_id")).Find(&customers)

	closeDB(db)
	c.JSON(200, customers)
}
func GetCustomersByCus_Name(c *gin.Context) {
	db := connectDB()
	var customers *Customer
	db.Where("name = $1", c.Param("cus_name")).Find(&customers)

	closeDB(db)
	c.JSON(200, customers)
}

func GetOrderByOrd_ID(c *gin.Context) {
	db := connectDB()
	var order []*Order
	db.Where("id = $1", c.Param("ord_id")).Find(&order)

	closeDB(db)
	c.JSON(200, order)
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

func GetItemsByOrderId(c *gin.Context) {
	db := connectDB()
	var item *Item
	db.Where("order_id = $1", c.Param("order_id")).Find(&item)

	var order *Order
	db.Where("id = $1", c.Param("order_id")).Find(&order)
	var Cus_id string = order.Customer_id
	var customer *Customer
	db.Where("id = $1", Cus_id).Find(&customer)

	closeDB(db)

	c.JSON(200, item.Id)
	c.JSON(200, order.Id)
	c.JSON(200, customer.Name)
}

type Order struct {
	Id          string `json:"Order_id"`
	Customer_id string `json:"customer_id"`
	Is_shipped  string `json:"is_shipped"`
}

func GetOrders(c *gin.Context) {
	id := c.Query("id")
	db := connectDB()
	var orders []*Order
	if id != "" {
		log.Println("Id: " + id)
		db.Preload(clause.Associations).Where("Id = $1", id).Find(&orders)
	} else {
		db.Preload(clause.Associations).Find(&orders)
	}
	closeDB(db)
	c.JSON(200, orders)
}

type Item struct {
	Id         uuid.UUID `json:"Item_id"`
	Order_id   string    `json:"order_id"`
	Product_id string    `json:"product_id"`
	Is_shipped string    `json:"is_shipped"`
}

func GetItems(c *gin.Context) {
	id := c.Query("id")
	db := connectDB()
	var items []*Item
	if id != "" {
		log.Println("Id: " + id)
		db.Preload(clause.Associations).Where("Id = $1", id).Find(&items)
	} else {
		db.Preload(clause.Associations).Find(&items)
	}
	closeDB(db)
	c.JSON(200, items)
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
