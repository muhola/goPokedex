package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/muhola/goPokedex/internal/cache"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client.BaseURL != "https://pokeapi.co/api/v2/" {
		t.Errorf("Expected BaseURL top be https://pokeapi.co/api/v2/")
	}
	if client.httpClient == nil {
		t.Errorf("Expected httpClient to not be nil")
	}
	if client.httpClient.Timeout != 10*time.Second {
		t.Errorf("Expected Default timeout to be 10 seconds, got %v", client.httpClient.Timeout)
	}
}

func TestGetLocation(t *testing.T) {
	mockLocationArea := mockLocationArea()
	// Create a test server that returns our mock response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockLocationArea)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL + "/"
	pokeCache := cache.NewCache(5 * time.Minute)
	// Test the GetLocation function
	result := client.GetLocation("", pokeCache)

	// Assertions
	assertLocationAreaEqual(t, mockLocationArea, result)
}
func assertLocationAreaEqual(t *testing.T, expected, actual LocationArea) {
	// Check top-level fields
	if expected.Count != actual.Count {
		t.Errorf("Count mismatch: expected %d, got %d", expected.Count, actual.Count)
	}

	if expected.Next != actual.Next {
		t.Errorf("Next URL mismatch: expected %s, got %s", expected.Next, actual.Next)
	}

	// Check if Previous is nil or has the correct value
	if (expected.Previous == "nil" && actual.Previous != "nil") ||
		(expected.Previous != "nil" && actual.Previous == "nil") {
		t.Errorf("Previous mismatch: expected %v, got %v", expected.Previous, actual.Previous)
	}

	// Check Results length
	if len(expected.Results) != len(actual.Results) {
		t.Errorf("Results length mismatch: expected %d, got %d",
			len(expected.Results), len(actual.Results))
		return // Early return to avoid index out of range errors
	}

	// Check each result
	for i, expResult := range expected.Results {
		actResult := actual.Results[i]
		if expResult.Name != actResult.Name {
			t.Errorf("Result[%d] name mismatch: expected %s, got %s",
				i, expResult.Name, actResult.Name)
		}
		if expResult.URL != actResult.URL {
			t.Errorf("Result[%d] name mismatch: expected %s, got %s",
				i, expResult.URL, actResult.URL)
		}
	}
}

// Helper Function to Mock-LocationArea
func mockLocationArea() LocationArea {
	return LocationArea{
		Count:    1000,
		Next:     "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
		Previous: "nil",
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{
				Name: "canalave-city-area",
				URL:  "https://pokeapi.co/api/v2/location-area/1/",
			},
			{
				Name: "eterna-city-area",
				URL:  "https://pokeapi.co/api/v2/location-area/2/",
			},
		},
	}
}
