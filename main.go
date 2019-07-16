package main

import (
	fetch "github.com/hank/investment-tool/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	fetch.Index()
}
