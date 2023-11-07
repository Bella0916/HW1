package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Department struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ShortName string    `json:"short_name"`
}

func GetDepartments(c *gin.Context) {

	db := connectDB()

	var departments []*Department

	db.Find(&departments)

	closeDB(db)
	c.JSON(200, departments)

}

func GetDepartmentById(c *gin.Context) {
	db := connectDB()
	var department *Department
	db.Where("id = $1", c.Param("DepartmentId")).Take(&department)

	closeDB(db)
	c.JSON(200, department)
}

func GetStudentsByDepartmentId(c *gin.Context) {
	db := connectDB()
	var students []*Customer
	db.Where("department_id = $1", c.Param("DepartmentId")).Find(&students)

	closeDB(db)
	c.JSON(200, students)
}
