package controller

import (
	"database/sql"
	"fmt"
	"go_gin_example/envconfig"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUser() string {
	return envconfig.GetEnv("NAME")
}

type user struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func GetUsersOldMethod(c *gin.Context) {

	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", envconfig.GetEnv("DB_HOST"), envconfig.GetEnv("DB_USER"), envconfig.GetEnv("DB_PASSWORD"), envconfig.GetEnv("DB_NAME"))
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	log.Println("Successfully created connection to database")
	rows, err := db.Query("SELECT * FROM users")
	checkError(err)
	var users []user
	for rows.Next() {
		var id uuid.UUID
		var name string
		err = rows.Scan(&id, &name)
		checkError(err)
		log.Printf("ID: %s, Name: %s\n", id, name)
		users = append(users, user{Id: id, Name: name})
	}

	db.Close()

	c.JSON(200, users)

}

func GetUserByIdOldMethod(c *gin.Context) {
	userID := c.Param("UserID")
	log.Println("UserID: " + userID)

	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s TimeZone=Asia/Taipei", envconfig.GetEnv("DB_HOST"), envconfig.GetEnv("DB_USER"), envconfig.GetEnv("DB_PASSWORD"), envconfig.GetEnv("DB_NAME"), envconfig.GetEnv("DB_WITH_SSL"))
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	log.Println("Successfully created connection to database")
	rows, err := db.Query("SELECT * FROM users WHERE id = $1", userID)
	checkError(err)
	var u user

	for rows.Next() {
		var id uuid.UUID
		var name string
		err = rows.Scan(&id, &name)
		checkError(err)
		log.Printf("ID: %s, Name: %s\n", id, name)
		u = user{Id: id, Name: name}
		break
	}

	db.Close()

	c.JSON(200, u)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
