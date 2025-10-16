package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type catfact struct {
	Fact string `json:"fact"`
}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

// FACTAPIURL Define a constant for the API URL to fetch cat facts
const FACTAPIURL = "https://catfact.ninja/fact"

// GetFact fetches a random fact from the external catfact.ninja API.
func GetFact(url string) (string, error) {

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request to catfacts API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		log.Printf("CatFacts API returned status code: %d\n", resp.StatusCode)
		return "", fmt.Errorf("catfacts API returned non-OK status: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var factResponse catfact
	if err := json.Unmarshal(bodyBytes, &factResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	log.Printf("Successfully fetched fact: %s\n", factResponse.Fact)

	return factResponse.Fact, nil
}
