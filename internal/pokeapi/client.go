package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/muhola/goPokedex/internal/cache"
)

func NewClient() *Client {
	return &Client{
		BaseURL:    "https://pokeapi.co/api/v2/",
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// fetchAndUnmarshal retrieves data from cache or API and unmarshals it
func (c *Client) fetchAndUnmarshal(url string, pokeCache *cache.Cache, result interface{}) error {
	// Check cache first
	cachedData, found := pokeCache.Get(url)
	if found {
		err := json.Unmarshal(cachedData, result)
		if err != nil {
			return fmt.Errorf("unmarshalling JSON from cache: %w", err)
		}
		return nil
	}

	// Fetch from API if not in cache
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP GET request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	// Add to cache
	pokeCache.Add(url, body)

	// Unmarshal the JSON
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("unmarshalling JSON from API: %w", err)
	}

	return nil
}

func (c *Client) GetLocation(url string, pokeCache *cache.Cache) LocationArea {
	if url == "" {
		url = c.BaseURL + "location/"
	}
	var pokeLocation LocationArea
	err := c.fetchAndUnmarshal(url, pokeCache, &pokeLocation)

	if err != nil {
		fmt.Println("Can not find location")
	}

	return pokeLocation
}

func (c *Client) GetPokemonBasedOnLocation(parameter string, pokeCache *cache.Cache) PokemonLocationEncounters {
	var pokemonfromLocation PokemonLocationEncounters
	if parameter == "" {
		return pokemonfromLocation
	}
	//cheack if user cityies in without -area prefex
	if !strings.Contains(parameter, "-area") {
		parameter += "-area"
	}
	url := c.BaseURL + "location-area/" + parameter

	err := c.fetchAndUnmarshal(url, pokeCache, &pokemonfromLocation)

	if err != nil {
		fmt.Println("Can not find location")
	}

	return pokemonfromLocation
}

func (c *Client) GetPokemon(parameter string, pokeCache *cache.Cache) Pokemon {
	var pokemon Pokemon
	if parameter == "" {
		return pokemon
	}

	url := c.BaseURL + "pokemon/" + parameter
	err := c.fetchAndUnmarshal(url, pokeCache, &pokemon)

	if err != nil {
		fmt.Printf("Can not find %s /n", parameter)
	}

	return pokemon
}
