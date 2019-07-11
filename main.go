package main

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/hank/investment-tool/route"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=disable")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	m.Steps(2)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("connect to DB has succeed")
	}
	route.Route()
	db.Close()
}
