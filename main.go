package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func fetchBitcoinPrice() {
	c := colly.NewCollector()
	c.OnHTML("#id-bitcoin", func(e *colly.HTMLElement) {
		price := e.ChildText(".market-cap")
		fmt.Println(price)
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=tool_db")
	if err != nil {
		fmt.Printf("failed")
	}
	defer db.Close()
	fetchBitcoinPrice()
}
