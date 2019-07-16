package main

import (
	fetch "github.com/hank/investment/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	fetch.Index()
}
