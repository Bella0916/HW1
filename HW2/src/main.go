package main

import (
	"go_gin_example/controller"
	"go_gin_example/envconfig"
	"log"

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

func main() {
	server := gin.Default()

	//GET /users
	server.GET("/users", getUsers)   // 讀取Users
	server.GET("/dbinfo", getDBInfo) // 讀取Users

	// 作業
	// 查詢商品清單 ok
	// 查詢客戶清單 ok
	// 查詢客戶有哪些訂單 ok
	// 查詢訂單明細
	//GET /customers
	server.GET("/products", controller.GetProducts)   // 查詢商品清單
	server.GET("/customers", controller.GetCustomers) // 查詢客戶清單

	server.GET("/orders/cus_id/:cus_id", controller.GetOrderByCus_ID)   // 讀取 Order By Cus_ID 查詢客戶有哪些訂單 By Cus_ID
	server.GET("/orders/name/:cus_name", controller.GetOrderByCus_Name) // 讀取 Order By Name 查詢客戶有哪些訂單 By Name

	//server.GET("/items/:order_id", controller.GetItemsByOrderId) // 讀取 Items By Order_ID

	//server.GET("/customers/id/:cus_id", controller.GetCustomersByCus_ID)       // 讀取 GetCustomers By ID
	//server.GET("/customers/name/:cus_name", controller.GetCustomersByCus_Name) // 讀取 Customers By Name

	//GET /orders
	server.GET("/orders/id/:ord_id", controller.GetOrderByOrd_ID) // 讀取 Order By ID

	server.GET("/items/:order_id", controller.GetItemsByOrderId) // 讀取 Items

	server.GET("/categories", controller.GetCategories) // 讀取 Categories
	server.GET("/orders", controller.GetOrders)         // 讀取 Orders
	server.GET("/items", controller.GetItems)           // 讀取 Items
	//server.GET("/students/:StudentId", controller.GetStudentById)                      // 讀取Student
	//server.GET("/students/:StudentId/department", controller.GetDepartmentByStudentId) // 讀取Department of Student

	//GET /departments
	server.GET("/departments", controller.GetDepartments)                                   // 讀取Departments
	server.GET("/departments/:DepartmentId", controller.GetDepartmentById)                  // 讀取Department
	server.GET("/departments/:DepartmentId/students", controller.GetStudentsByDepartmentId) // 讀取Students of Department
	//GET /ousers with old method
	server.GET("/ousers", controller.GetUsersOldMethod)            // 讀取Users
	server.GET("/ousers/:UserId", controller.GetUserByIdOldMethod) // 讀取Users
	//GET /courses
	//server.GET("/courses", controller.GetCourses)              // 讀取Courses
	server.GET("/courses/:CourseId", controller.GetCourseById) // 讀取Courses
	if err := server.Run(":" + envconfig.GetEnv("PORT")); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
