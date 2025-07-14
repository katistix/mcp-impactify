package service

import (
	"fmt"
	"io"
	"net/http"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type DeleteWidgetArguments struct {
	ID string `json:"id"`
}

func (s *Service) HandleDeleteWidget(arguments DeleteWidgetArguments) (*mcp_golang.ToolResponse, error) {
	// Construct the URL for the widget endpoint
	widgetURL := fmt.Sprintf("%s/widget/%s", s.baseURL, arguments.ID)

	req, err := http.NewRequest("DELETE", widgetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+s.apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API call to %s: %w", widgetURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call to %s returned non-200 status: %d, body: %s", widgetURL, resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(bodyBytes))), nil
}
