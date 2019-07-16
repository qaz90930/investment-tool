package fund

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
)

type Fund struct {
	Name    string          `json:"fund_name"`
	Price   decimal.Decimal `json:"price"`
	Created time.Time       `json:"created"`
}

var connStr = "postgres://hank:password@localhost:5432/tool_db?sslmode=disable"

func getPrice(c echo.Context) error {
	db, err := sql.Open("postgres", connStr)
	rows := db.QueryRow(`SELECT fund_name, price, created FROM fund ORDER BY fund_name`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return json.NewEncoder(c.Response()).Encode(rows)
}
