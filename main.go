package main

import (
	fetchprice "github.com/hank/investment-tool/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	fetchprice.Index()
}
