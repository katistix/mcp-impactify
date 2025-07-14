package service

import (
	"net/http"
	"time"
)

// Service holds the base URL and API key for making API calls.
type Service struct {
	baseURL string
	apiKey  string
	client  *http.Client // http client for making requests
}

// NewService creates a new Service instance.
func NewService(baseURL, apiKey string) *Service {
	return &Service{
		baseURL: baseURL,
		apiKey:  apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second, // Set a reasonable timeout
		},
	}
}
