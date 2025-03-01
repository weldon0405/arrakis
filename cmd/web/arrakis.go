package main

import (
	"log"
	"net/http"
)

func main() {
	Get_tickers()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	log.Print("Welcome to Arrakis!")

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
