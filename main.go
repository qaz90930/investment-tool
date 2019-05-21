package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BitcoinPrice struct {
	Ask  float64 `json:"ask"`
	Bid  float64 `json:"bid"`
	Last float64 `json:"last"`
	High float64 `json:"high"`
	Low  float64 `json:"low"`
	Open struct {
		Hour   float64 `json:"hour"`
		Day    float64 `json:"day"`
		Week   float64 `json:"week"`
		Month  float64 `json:"month"`
		Month3 float64 `json:"month_3"`
		Month6 float64 `json:"month_6"`
		Year   float64 `json:"year"`
	} `json:"open"`
	Averages struct {
		Day   float64 `json:"day"`
		Week  float64 `json:"week"`
		Month float64 `json:"month"`
	} `json:"averages"`
	Volume  float64 `json:"volume"`
	Changes struct {
		Percent struct {
			Hour   float64 `json:"hour"`
			Day    float64 `json:"day"`
			Week   float64 `json:"week"`
			Month  float64 `json:"month"`
			Month3 float64 `json:"month_3"`
			Month6 float64 `json:"month_6"`
			Year   float64 `json:"year"`
		} `json:"percent"`
		Price struct {
			Hour   float64 `json:"hour"`
			Day    float64 `json:"day"`
			Week   float64 `json:"week"`
			Month  float64 `json:"month"`
			Month3 float64 `json:"month_3"`
			Month6 float64 `json:"month_6"`
			Year   float64 `json:"year"`
		} `json:"price"`
	} `json:"changes"`
	VolumePercent    float64 `json:"volume_percent"`
	Timestamp        int     `json:"timestamp"`
	DisplayTimestamp string  `json:"display_timestamp"`
	DisplaySymbol    string  `json:"display_symbol"`
}

func main() {
	r := gin.Default()
	respone, err := http.Get("https://apiv2.bitcoinaverage.com/indices/global/ticker/BTCUSD")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		data, _ := ioutil.ReadAll(respone.Body)
		var bp BitcoinPrice
		json.Unmarshal([]byte(data), &bp)
		fmt.Println(bp.Last)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
