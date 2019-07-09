package main

import (
	"fmt"

	"github.com/hank/investment-tool/route"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=tool_db password=password sslmode=disable")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("connect to DB has succeed")
	}
	route.Route()
	db.LogMode(true)
	db.Close()
}
