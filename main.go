package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func showCryptoPrice(c echo.Context) error {
	return c.String(http.StatusOK, "Crypto Page")
}

func fetchBitcoinPrice() {
	c := colly.NewCollector()
	c.OnHTML("#id-bitcoin", func(e *colly.HTMLElement) {
		price := e.ChildText(".market-cap")
		fmt.Println(price)
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=tool_db password=password sslmode=disable")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("connect to DB has succeed")
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/crypto", showCryptoPrice)
	// db.Close()
	e.Start(":8001")
	db.Close()
}
