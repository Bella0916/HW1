package controller

import (
	//"go_gin_example/model"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Course struct {
	Id         uuid.UUID `json:"id"`
	CourseCode string    `json:"course_code"`
	Name       string    `json:"name"`
}

func GetCategory(c *gin.Context) {
	courseCode := c.Query("course_code")

	db := connectDB()

	if courseCode != "" {
		var course *Course
		db.Where("course_code = $1", courseCode).Find(&course)
		closeDB(db)
		c.JSON(200, course)
		return
	}
	var courses []*Course

	db.Find(&courses)

	closeDB(db)
	c.JSON(200, courses)

}

func GetCourseById(c *gin.Context) {
	db := connectDB()
	var course *Course
	log.Println("CourseId: " + c.Param("CourseId"))
	db.Where("id = $1", c.Param("CourseId")).Take(&course)

	closeDB(db)
	c.JSON(200, course)
}

func GetCourseByCode(c *gin.Context) {
	db := connectDB()
	var course *Course
	db.Where("course_code = $1", c.Param("course_code")).Take(&course)

	closeDB(db)
	c.JSON(200, course)
}
