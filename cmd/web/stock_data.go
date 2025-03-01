package main

import (
	"fmt"
	"net/http"
	"os"
)

func Get_tickers() {
	fmt.Println("Lets get some stock tickers!")
	av_api_key := os.Getenv("ALPHA_VANTAGE_API_KEY")
	fmt.Println(av_api_key)

	// resp, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", ticker, av_api_key)
	resp, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=IBM&apikey=demo")
	if err != nil {
		fmt.Printf("Error during request: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Client: Received response!")
	fmt.Println("Client: Status code:")
	fmt.Println(resp)

}
