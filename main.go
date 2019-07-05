package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
)

var db *gorm.DB
var err error

type CryptoPrice struct {
	ID      uint
	Name    string
	Price   decimal.Decimal
	Created time.Time
}

func showCryptoPrice(c echo.Context) error {
	fetchBitcoinPrice()
	return c.String(http.StatusOK, "Crypto Page")
}

func fetchBitcoinPrice() {
	c := colly.NewCollector()
	c.OnHTML("#id-bitcoin", func(e *colly.HTMLElement) {
		price := e.ChildText(".market-cap")
		fmt.Println(price)
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")
	var bitcoin = CryptoPrice{Name: "Bitcoin", Price: decimal.NewFromFloat(1000.00), Created: time.Now()}
	db.Create(&bitcoin)
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

	e.Start(":8001")
	db.Close()
}
