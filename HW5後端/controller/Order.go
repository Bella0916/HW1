package controller

import (
	"log"
	"strconv"

	//"strconv"

	"github.com/gin-gonic/gin"
)

type Order struct {
	Id         int  `json:"id"`
	CustomerID int  `json:"customer_id"`
	IsPaid     bool `json:"is_paid"`
}

/*type CustomerProduct struct {
	Id         int `json:"id"`
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"order_id"`
}*/

type CustomerProductSelection struct {
	Id         int `gorm:"primaryKey;type:int;default:gen_random_int()"  json:"id"`
	CustomerId int `gorm:"unique,composite:customer_order_unique_constraint" json:"customer_id"`
	ProductId  int `gorm:"unique,composite:customer_order_unique_constraint" json:"order_id"`
}

func GetCustomerOrders(c *gin.Context) {
	db := connectDB()
	var orders []*Order
	db.Find(&orders)

	closeDB(db)
	c.JSON(200, orders)
}

func GetOrdersByCustomerID(c *gin.Context) {
	db := connectDB()
	var order *Order
	db.Where("customer_id = $1", c.Param("CustomerId")).Take(&order)

	closeDB(db)
	c.JSON(200, order)
}

func (CustomerProductSelection) TableName() string {
	return "customer_order"
}

var idcount int = 6001

func CustomerSelectProduct(c *gin.Context) {
	db := connectDB()
	customerId, err := strconv.Atoi(c.Param("CustomerId"))

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "CustomerId is not a valid Id",
		})
		closeDB(db)
		return
	}

	orderId, err := strconv.Atoi(c.Param("ProductId"))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "ProductId is not a valid Id",
		})
		closeDB(db)
		return
	}

	customerProductSelection := &CustomerProductSelection{
		Id:         idcount,
		CustomerId: customerId,
		ProductId:  orderId,
	}

	result := db.Create(&customerProductSelection)
	if result.Error != nil {
		log.Println(result.Error)

		c.JSON(500, gin.H{
			"message": "do CustomerProductSelection failed with error: " + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, customerProductSelection)
	idcount++
}

func CreateOrders(c *gin.Context) {
	db := connectDB()
	var order *Order
	c.BindJSON(&order)
	result := db.Create(&order)

	if result.Error != nil {
		log.Println(result.Error)

		c.JSON(500, gin.H{
			"message": "Create order failed with error: " + result.Error.Error(),
		})
		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, order)
}

/*func CreateOrdersByCustomerID(c *gin.Context) {
	db := connectDB()
	var order *Order
	queryResult := db.Where("customer_id = $1", c.Param("CustomerId")).Take(&order)
	if queryResult.Error != nil {
		log.Println(queryResult.Error)
		c.JSON(500, gin.H{
			"message": "Create order failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}
	result := db.Create(&order)

	if result.Error != nil {
		log.Println(result.Error)

		c.JSON(500, gin.H{
			"message": "create order failed with error: " + result.Error.Error(),
		})

		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Create orders Susseccfully",
	})
}


func DeletecustomerProductByCustomerID(c *gin.Context) {
	db := connectDB()
	var customerProduct *CustomerProduct
	queryResult := db.Where("customer_id = $1", c.Param("CustomerId")).Take(&customerProduct)

	if queryResult.Error != nil {
		log.Println(queryResult.Error)
		c.JSON(500, gin.H{
			"message": "Delete customerProduct failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	result := db.Delete(&customerProduct)

	if result.Error != nil {
		log.Println(result.Error)

		c.JSON(500, gin.H{
			"message": "Delete customerProduct failed with error: " + result.Error.Error(),
		})

		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Delete customerProduct Susseccfully",
	})
}


*/

func DeleteOrdersByID(c *gin.Context) {
	db := connectDB()
	var order *Order
	queryResult := db.Where("id = $1", c.Param("OrderId")).Take(&order)

	if queryResult.Error != nil {
		log.Println(queryResult.Error)
		c.JSON(500, gin.H{
			"message": "Delete order failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	result := db.Delete(&order)

	if result.Error != nil {
		log.Println(result.Error)

		c.JSON(500, gin.H{
			"message": "Delete order failed with error: " + result.Error.Error(),
		})

		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Delete orders Susseccfully",
	})
}

func UpdateOrdersByCustomerID(c *gin.Context) {
	db := connectDB()
	var order *Order
	queryResult := db.Where("customer_id = $1", c.Param("CustomerId")).Take(&order)

	if queryResult.Error != nil {
		log.Println(queryResult.Error)
		c.JSON(500, gin.H{
			"message": "Update order failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	var orderBody *Order
	c.BindJSON(&orderBody)
	orderBody.CustomerID = order.CustomerID
	result := db.Model(&order).Where("customer_id = ?", order.CustomerID).Updates(orderBody)

	if result.Error != nil {
		log.Println(result.Error)

		c.JSON(500, gin.H{
			"message": "Update order_IsPaid failed with error: " + result.Error.Error(),
		})

		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Update orders Susseccfully",
	})
}
