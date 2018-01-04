// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// handlers.go [created: Wed,  3 Jan 2018]

package handlers

import (
	"github.com/eddiefisher/flypoolapi/src/api"
	"github.com/eddiefisher/flypoolapi/src/endpoint"
)

type Handlers interface {
	Request(endpoint.Endpoint, string, api.Api)
}
