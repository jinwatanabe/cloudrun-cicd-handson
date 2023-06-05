package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id  int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Age  int `json:"age"`
}

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	// dsn := fmt.Sprintf("%s:%s@tcp(aws.connect.psdb.cloud)/[DB名]?tls=true", user, password)
	dsn := fmt.Sprintf("%s:%s@tcp(aws.connect.psdb.cloud)/neko?tls=true", user, password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Deployed!!",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Unscoped().Find(&users)
		c.JSON(200, users)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}