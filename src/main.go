// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Wed,  3 Jan 2018]

package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/eddiefisher/flypoolapi/src/api"
	"github.com/eddiefisher/flypoolapi/src/endpoint"
	"github.com/eddiefisher/flypoolapi/src/handlers"
	"github.com/eddiefisher/flypoolapi/src/version"
)

func init() {
	logrus.Printf(
		"commit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
}

func main() {
	_wallet := flag.String("w", "t1Txkfef1wWfpqYsAWjhe4otJnwE7DZeBwY", "Wallet.")
	wallet := *_wallet
	if wallet == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	endpoints := endpoint.Endpoints{
		{"Ethermine.org", "https://api.ethermine.org"},
		{"Ethpool.org", "http://api.ethpool.org"},
		{"Ethermine ETC", "https://api-etc.ethermine.org"},
		{"Flypool Zcash", "https://api-zcash.flypool.org"},
	}
	endpoint, err := endpoints.Get("Flypool Zcash")
	if err != nil {
		logrus.Fatal(err)
	}

	api := api.Api{http.Client{Timeout: time.Second * 2}}

	payouts := handlers.Payouts{}
	payouts.Request(endpoint, wallet, api)

	logrus.Println(endpoint.Url, "statistic:", payouts.Sum())
}
