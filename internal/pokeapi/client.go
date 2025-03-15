package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const baseLocationUrl = "https://pokeapi.co/api/v2/location-area/"

func GetLocation(url string) LocationData {
	if url == "" {
		url = baseLocationUrl
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	pokiLocation := LocationData{}

	if err := json.Unmarshal(body, &pokiLocation); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return pokiLocation
}
