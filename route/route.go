package route

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
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
		rarPrice := e.ChildText(".price")
		price := strings.Trim(rarPrice, "$")
		fmt.Println(price)
		connStr := "postgres://hank:password@localhost:5432/tool_db?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		fmt.Printf("Successfully Connected")
		db.QueryRow(`INSERT INTO fund (fund_name, price, created) VALUES ($1, $2, $3)`, "bitcoin", price, time.Now())
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
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
