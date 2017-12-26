package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"home.dev/toster/flypoolapi/api"
	"home.dev/toster/flypoolapi/endpoint"
	"home.dev/toster/flypoolapi/handlers"
)

func main() {
	_wallet := flag.String("w", "", "Wallet.")
	wallet := *_wallet
	endpoints := endpoint.Endpoints{
		{"Ethermine.org", "https://api.ethermine.org"},
		{"Ethpool.org", "http://api.ethpool.org"},
		{"Ethermine ETC", "https://api-etc.ethermine.org"},
		{"Flypool Zcash", "https://api-zcash.flypool.org"},
	}
	endpoint, err := endpoints.Get("Flypool Zcash")
	if err != nil {
		log.Fatal(err)
	}

	api := api.Api{http.Client{Timeout: time.Second * 2}}

	payouts := handlers.Payouts{}
	payouts.Request(endpoint, wallet, api)

	fmt.Println(endpoint.Url, "statistic:", payouts.Sum())
}
