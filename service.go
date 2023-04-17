package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetFact(context.Context) (*Fact, error)
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type FactService struct {
	url   string
	token string
}

var Client HTTPClient

// init() is a very special function meant to execute prior to the main() function
func init() {
	Client = &http.Client{}
}

func NewFactService(url string, token string) Service {
	return &FactService{
		url:   url,
		token: token,
	}
}

func (s *FactService) GetFact(ctx context.Context) (*Fact, error) {
	// create a new http request
	request, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return nil, err
	}
	// add x-api-key to header
	request.Header.Add("x-api-key", s.token)
	// send http request
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer response.Body.Close()
	// fact
	fact := &Fact{}
	if err := json.NewDecoder(response.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
