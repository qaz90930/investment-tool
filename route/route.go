package route

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
)

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
	db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/tool_db?sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	// var bitcoin = CryptoPrice{Name: "Bitcoin", Price: decimal.NewFromFloat(1000.00), Created: time.Now()}
	db.QueryRow(`INSERT INTO price(name, price, created) VALUES ('Name', 'Price', 'Created')`)
}

func Route() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/crypto", showCryptoPrice)
	e.Start(":8001")
}
