package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB
var err error

type BitcoinPrice struct {
	gorm.Model
	ID      int       `json: "id"`
	Price   string    `json: "price"`
	Created time.Time `json: "created"`
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=app_db sslmode=disable")
	if err != nil {
		panic("failed to connect the DB")
	}
	defer db.Close()
	r := gin.Default()
	r.GET("/bitcoin", GetBitcoinPrice)
	r.Run()
}

func GetBitcoinPrice(c *gin.Context) {
	var price BitcoinPrice
	if err := db.Find(price).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, price)
	}
}
