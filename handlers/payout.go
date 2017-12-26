package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"home.dev/toster/flypoolapi/api"
	"home.dev/toster/flypoolapi/endpoint"
)

type Payouts struct {
	Status string `json:"status"`
	Data   []struct {
		Start  int    `json:"start"`
		End    int    `json:"end"`
		Amount int    `json:"amount"`
		TxHash string `json:"txHash"`
		PaidOn int    `json:"paidOn"`
	} `json:"data"`
}

func (payouts Payouts) Sum() (result float32) {
	for _, k := range payouts.Data {
		f := float32(k.Amount) / 100000000.0
		result = result + f
	}
	return result
}

func (payouts *Payouts) Request(endpoint endpoint.Endpoint, wallet string, api api.Api) {
	url := strings.Join([]string{endpoint.Url, "miner", wallet, "payouts"}, "/")
	body := api.Request(http.MethodGet, url, nil)
	err := json.Unmarshal(body, &payouts)
	if err != nil {
		log.Fatal(err)
	}
}
