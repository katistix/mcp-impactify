package main

import (
	"log"
	"net/http"
	"os"
)

func initializeWithHealthCheck() (string, string, error) {
	// Use positional argument for baseURL
	var baseURL string
	if len(os.Args) > 1 {
		baseURL = os.Args[1]
	} else {
		baseURL = "https://impactify.ro/api/rest/v1"
	}

	// Print the base path that is decided
	log.Printf("Using base URL: %s", baseURL)

	// Do a health check fetch to `/health`
	resp, err := http.Get(baseURL + "/health")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Health check failed")
	} else {
		log.Println("Health check passed")
	}

	// Read environment variables
	apiKey := os.Getenv("IMPACTIFY_EVENT_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing required environment variable: IMPACTIFY_EVENT_API_KEY")
	}
	// Fetch `/event` and log the response body with Bearer token
	req, err := http.NewRequest("GET", baseURL+"/event", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Event fetch failed")
	} else {
		log.Println("Event fetch passed")
	}

	log.Println("Initialization complete. All systems are go!")

	return baseURL, apiKey, nil
}
