package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/muhola/goPokedex/internal/cache"
)

const baseLocationUrl = "https://pokeapi.co/api/v2/location-area/"

func NewClient() *Client {
	return &Client{
		BaseURL:    "https://pokeapi.co/api/v2/",
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) GetLocation(url string, pokeCache *cache.Cache) LocationArea {
	var pokeLocation LocationArea
	cachedData, found := pokeCache.Get(url)
	fmt.Println(url)
	if found {
		fmt.Println("Cache Data found")
		err := json.Unmarshal(cachedData, &pokeLocation)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
		return pokeLocation
	}

	if url == "" {
		url = c.BaseURL + "location-area/"
	}
	fmt.Println("Make Request")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	pokeCache.Add(url, body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &pokeLocation); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return pokeLocation
}
