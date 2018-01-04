// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// endpoint.go [created: Wed,  3 Jan 2018]

package endpoint

import "errors"

type Endpoints []Endpoint

type Endpoint struct {
	Key string
	Url string
}

func (endpoints Endpoints) Get(key string) (Endpoint, error) {
	for _, k := range endpoints {
		if k.Key == key {
			return k, nil
		}
	}
	return Endpoint{}, errors.New("Endpoint not found")
}
