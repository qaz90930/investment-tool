package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/labstack/echo"
)

func showCryptoPrice(c echo.Context) error {
	// scheduler.Every(5).Seconds().Run(fetchBitcoinPrice)
	// scheduler.Every(5).Seconds().Run(fetchStockInfo)
	fetchStockInfo()
	return c.String(http.StatusOK, "Crypto Page")
}

func fetchBitcoinPrice() {
	c := colly.NewCollector()
	c.OnHTML("#id-bitcoin", func(e *colly.HTMLElement) {
		rawPrice := e.ChildText(".price")
		price := strings.Trim(rawPrice, "$")
		fmt.Println(price)
		connStr := "postgres://hank:password@localhost:5432/tool_db?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		db.QueryRow(`INSERT INTO fund (fund_name, price, created) VALUES ($1, $2, $3)`, "bitcoin", price, time.Now())
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")
}

func fetchStockInfo() {
	c := colly.NewCollector()
	c.OnHTML(".ListFundListing__marketPrice__3yFc0", func(e *colly.HTMLElement) {
		rawPrice := e.ChildText(".NumberElements__dollarValue__1BUEi")
		// price := strings.Trim(rawPrice, "$")
		fmt.Println(e.Text, rawPrice)
		fmt.Println("124")
	})
	c.Visit("https://app.sharesies.nz/invest/companies")
}

func Index() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/price-list", showCryptoPrice)
	e.Start(":8001")
}
