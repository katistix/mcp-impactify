package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type UpdateEventArguments struct {
	Title          *string `json:"title,omitempty"`
	LivestreamLink *string `json:"livestreamLink,omitempty"`
	Location       *struct {
		Latitude  *float64 `json:"latitude"`
		Longitude *float64 `json:"longitude"`
	} `json:"location,omitempty"`
}

func (s *Service) HandleUpdateEvent(arguments UpdateEventArguments) (*mcp_golang.ToolResponse, error) {
	updateEventUrl := fmt.Sprintf("%s/event", s.baseURL)

	// Marshal the arguments to JSON for the request body
	bodyBytes, err := json.Marshal(arguments)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal arguments: %w", err)
	}

	req, err := http.NewRequest("POST", updateEventUrl, io.NopCloser(bytes.NewReader(bodyBytes)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call to %s returned non-200 status: %d, body: %s", updateEventUrl, resp.StatusCode, string(respBodyBytes))
	}

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(respBodyBytes))), nil
}
