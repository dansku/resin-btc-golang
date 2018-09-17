package main

import (
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// https://mholt.github.io/json-to-go/
type BitcoinData struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		USD struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		GBP struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		EUR struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func getBitcoinPrice(url string) float64 {

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	jsn, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var bitcoindata BitcoinData

	err = json.Unmarshal(jsn, &bitcoindata)

	if err != nil {
		log.Fatal(err)
	}

	price := convertStringPriceToFloat(bitcoindata.Bpi.USD.Rate)

	return price

}

func convertStringPriceToFloat(price string) float64 {

	str := strings.Replace(price, ",", "", -1)
	floatPrice, err := strconv.ParseFloat(str, 64)

	if err != nil {
		log.Fatal(err)
	}

	return floatPrice

}

func printPrices(daysRunning int, currentTime string, currentPrice float64, todayHigh float64, todayLow float64, openingPrice float64) {

	c := color.New(color.FgRed, color.BgBlack, color.Bold)

	if currentPrice >= openingPrice {
		c = color.New(color.FgGreen, color.BgBlack, color.Bold)
	}

	c.Printf("[%v] - %v /  Current Price: $%.2f  - High Price: $%.2f  - Low Price: $%.2f  - Open Price: $%.2f \n", daysRunning, currentTime, currentPrice, todayHigh, todayLow, openingPrice)
}
