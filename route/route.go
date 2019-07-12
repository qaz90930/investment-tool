package route

import (
	"fmt"
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
		price := e.ChildText(".price")
		fmt.Println(price)
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")
}

func Route() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/crypto", showCryptoPrice)
	e.Start(":8001")
}
