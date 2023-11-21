package main

import (
	"go_gin_example/controller"
	"go_gin_example/envconfig"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()
	//---HW3
	// Business
	//Invoke-RestMethod -Uri http://localhost:8080/products/ -Method Post
	server.POST("/products", controller.GetCreateProduct)
	//Invoke-RestMethod -Uri http://localhost:8080/products/0187bed0-0071-4e23-8448-c105127163ff -Method Delete
	server.DELETE("/products/:products_id", controller.GetDeleteProduct)
	//Invoke-RestMethod -Uri http://localhost:8080/customers/ -Method Post
	server.POST("/customers", controller.GetCreateCustomer)
	//Invoke-RestMethod -Uri http://localhost:8080/orders/ -Method Delete
	server.POST("/orders", controller.GetCreateOrder)
	//Invoke-RestMethod -Uri http://localhost:8080/orders/ce7377ee-596c-4b9d-a26f-44a12c667cf1/pay -Method Put
	server.PUT("/orders/:id/pay", controller.GetUpdateOrderPaymentStatus)
	//Invoke-RestMethod -Uri http://localhost:8080/items/0551c33f-54c6-46b8-b161-fc9dd8d73c9a/ship -Method Put
	server.PUT("/items/:id/ship", controller.GetUpdateItemShippingStatus)
	//---HW3

	//---HW2
	// Products----1
	//http://localhost:8080/products/
	server.GET("/products", controller.GetProducts) //
	//http://localhost:8080/products/ctg/c7bf349c-7d86-4db2-9db8-4fc55fa75113
	server.GET("/products/ctg/:ctg_id", controller.GetProducts_CategoryID) // by CategoryID
	//----1

	// Categories----2
	//http://localhost:8080/categories/
	server.GET("/categories", controller.GetCategories) // 讀取 Categories for ALL
	//http://localhost:8080/categories/name/CPU
	server.GET("/categories/name/:crg_name", controller.GetCategories_Name) // 讀取 Categories By Name
	//http://localhost:8080/categories/id/4e9981fb-3060-4980-99c9-cb91a61c734b
	server.GET("/categories/id/:crg_id", controller.GetCategories_ID) // 讀取 Categories By ID
	//----2

	// Customers----3
	//http://localhost:8080/customers
	server.GET("/customers", controller.GetCustomers) // 查詢 Customers for ALL
	//http://localhost:8080/customers/id/06d4a6ed-2af0-475b-a141-f8caaff9473b
	server.GET("/customers/id/:cus_id", controller.GetCustomersByCus_ID) // 讀取 GetCustomers By ID
	//http://localhost:8080/customers/name/Ying
	server.GET("/customers/name/:cus_name", controller.GetCustomersByCus_Name) // 讀取 Customers By Name
	//----3

	// Orders----4
	//http://localhost:8080/orders
	server.GET("/orders", controller.GetOrders) // 讀取 Orders for ALL
	//http://localhost:8080/orders/id/ce7377ee-596c-4b9d-a26f-44a12c667cf1
	server.GET("/orders/id/:ord_id", controller.GetOrderByOrd_ID) // 讀取 Order By ID
	//http://localhost:8080/orders/cus_id/06d4a6ed-2af0-475b-a141-f8caaff9473b
	server.GET("/orders/cus_id/:cus_id", controller.GetOrderByCus_ID) // 讀取 Order By Cus_ID
	//http://localhost:8080/orders/name/Ying
	server.GET("/orders/name/:cus_name", controller.GetOrderByCus_Name) // 讀取 Order By Name
	//----4

	// Items----5
	//http://localhost:8080/items
	server.GET("/items", controller.GetItems) // 讀取 Items for ALL
	//http://localhost:8080/items/id/b211df2f-d6c1-49d7-86e9-4095e8167ee5
	server.GET("/items/id/:id", controller.GetItemsById) // 讀取 Items By id
	//http://localhost:8080/items/order_id/f361f657-a5a1-4607-a9fd-f01679bdab88
	server.GET("/items/order_id/:order_id", controller.GetItemsByOrderId) // 讀取 Items By order_id
	//http://localhost:8080/items/product_id/0297a69a-6730-485a-be7f-be1fb93a3c66
	server.GET("/items/product_id/:product_id", controller.GetItemsByProductId) // 讀取 Items By product_id
	//----5

	if err := server.Run(":" + envconfig.GetEnv("PORT")); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
