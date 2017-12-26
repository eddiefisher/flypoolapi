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
