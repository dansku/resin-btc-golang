package main

import (
	"time"
)

func main() {

	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	// Initialize variables
	currentTime := time.Now().Local().Format("02-01-2006")
	initialTime := currentTime

	currentPrice := getBitcoinPrice(url)
	openingPrice := currentPrice
	todayHigh := currentPrice
	todayLow := currentPrice
	daysRunning := 1

	for {

		currentPrice = getBitcoinPrice(url)

		currentTime = time.Now().Local().Format("02-01-2006")

		// Reset data on new day
		if currentTime != initialTime {
			initialTime = currentTime
			openingPrice = currentPrice
			todayHigh = currentPrice
			todayLow = currentPrice
			daysRunning++
		}

		if currentPrice > todayHigh {
			todayHigh = currentPrice
		}

		if currentPrice < todayLow {
			todayLow = currentPrice
		}

		printPrices(daysRunning, currentTime, currentPrice, todayHigh, todayLow, openingPrice)

		time.Sleep(10 * time.Second)

	}

}
