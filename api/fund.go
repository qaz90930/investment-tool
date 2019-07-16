package fund

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
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
var fundName string
var price decimal.Decimal
var created time.Time

func Price(c echo.Context) error {
	db, err := sql.Open("postgres", connStr)
	rows, err := db.Query(`SELECT fund_name, price, created FROM fund ORDER BY fund_name`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println(reflect.TypeOf(rows), &rows)
	for rows.Next() {
		rows.Scan(&fundName, &price, &created)
		fmt.Println(fundName, price, created)
	}
	return json.NewEncoder(c.Response()).Encode(rows)
}
