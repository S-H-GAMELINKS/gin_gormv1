package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"runtime/debug"
)

type User struct {
	ID uint64 `gorm:'primary_key'`
	Name string
	Tweets []Tweet
}

type Tweet struct {
	ID uint64 `gorm:'primary_key'`
	UserID uint64
	Content string
}

func main() {
	debug.SetGCPercent(100)

	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	sqlite3, err := gorm.Open("sqlite3", "./gorm.db")
	defer sqlite3.Close()
	if err != nil {
		sqlite3.Close()
	}

	sqlite3.LogMode(true)

	r.GET("/", func(c *gin.Context) {
		var users []User
		sqlite3.Preload("Tweets").Find(&users)
		c.JSON(200, "Hello world!")
	})

	r.Run(":19090")
}
