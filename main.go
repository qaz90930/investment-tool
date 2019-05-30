package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BitcoinPrice struct {
	gorm.Model
	Price string
	// Date  *time.Time
}

func Bitcoin(c *gin.Context) {
	db, err := gorm.Open("postgres", "host=localhost dbname=app_db")
	if err != nil {
		panic("failed to connect the DB")
	}
	defer db.Close()
	var price BitcoinPrice
	data := db.First(price)
	fmt.Printf("%s", data)
}

func main() {
	r := gin.Default()
	r.GET("/bitcoin", Bitcoin)
	r.Run()
}
