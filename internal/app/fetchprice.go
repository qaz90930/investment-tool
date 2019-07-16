package fetchprice

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/gocolly/colly"
	"github.com/labstack/echo"
)

var connStr = "postgres://hank:password@localhost:5432/tool_db?sslmode=disable"

func landingPage(c echo.Context) error {
	return c.String(http.StatusOK, "Landing Page")
}

func showCryptoPrice(c echo.Context) error {
	scheduler.Every(1).Minutes().Run(fetchBitcoinPrice)
	scheduler.Every(1).Minutes().Run(fetchStockInfo)
	return c.String(http.StatusOK, "Price list")
}

func fetchBitcoinPrice() {
	c := colly.NewCollector()
	c.OnHTML("#id-bitcoin", func(e *colly.HTMLElement) {
		rawData := e.ChildText(".price")
		price := strings.Trim(rawData, "$")
		fmt.Println("bitcoin", price)
		db, err := sql.Open("postgres", connStr)
		db.QueryRow(`INSERT INTO fund (fund_name, price, created) VALUES ($1, $2, $3)`, "bitcoin", price, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")
}

func fetchStockInfo() {
	c := colly.NewCollector()
	c.OnHTML(".small-12.medium-5.columns", func(e *colly.HTMLElement) {
		rawData := e.ChildText("h1")
		price := strings.Trim(rawData, "$")
		fmt.Println("air new zealand", price)
		db, err := sql.Open("postgres", connStr)
		db.QueryRow(`INSERT INTO fund (fund_name, price, created) VALUES ($1, $2, $3)`, "Air New Zealand", price, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	})
	c.Visit("https://www.nzx.com/instruments/AIR")
}

func Index() {
	// This function contains all the routes
	e := echo.New()
	e.GET("/", landingPage)
	e.GET("/price-list", showCryptoPrice)
	e.Start(":8001")
}
