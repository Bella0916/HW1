package main

import (
	"go_gin_example/controller"
	"go_gin_example/envconfig"
	"log"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET Users by " + controller.GetUser(),
	})
}

func getDBInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Port" + envconfig.GetEnv("DB_PORT"),
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Private-Network", "true")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Category, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	server := gin.Default()

	//server.Use(cors.Default())
	server.Use(CORSMiddleware())

	//GET
	//users
	server.GET("/users", getUsers)
	server.GET("/dbinfo", getDBInfo)
	//categories
	server.GET("/categories", controller.GetCategories)
	//poducts
	server.GET("/products", controller.GetProducts)
	server.GET("/products/:ProductId", controller.GetProductsByID)
	//customers
	server.GET("/customers", controller.GetCustomers)
	server.GET("/customers/:CustomerId", controller.GetCustomersByID)
	//orders
	server.GET("/orders", controller.GetCustomerOrders)
	server.GET("/orders/:CustomerId", controller.GetOrdersByCustomerID)
	//item
	server.GET("/items", controller.GetOrderItems)
	server.GET("/items/:OrderId", controller.GetItemsByOrderID)

	//Post
	//products
	server.POST("/products", controller.CreateProducts)
	//customers
	server.POST("/customers", controller.CreateCustomers)
	//orders
	server.POST("/orders", controller.CreateOrders)

	//PUT
	//products
	server.PUT("/products/:ProductId", controller.UpdateProductByID)
	//orders
	server.PUT("/orders/:CustomerId", controller.UpdateOrdersByCustomerID)
	server.PUT("/orders/:CustomerId/products/:ProductId", controller.CustomerSelectProduct)
	//item
	server.PUT("/items/:OrderId", controller.UpdateItemsByOrderID)

	//DELETE
	//products
	server.DELETE("/products/:ProductId", controller.DeleteProductsByID)
	//orders
	server.DELETE("/orders/:OrderId", controller.DeleteOrdersByID)

	if err := server.Run(":" + envconfig.GetEnv("PORT")); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
