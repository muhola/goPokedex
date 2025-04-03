package pokeapi

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	httpClient *http.Client
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
