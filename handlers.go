package main

import (
	"fmt"
	"io"
	"net/http"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type GetEventArguments struct{}

// handleGetEvent is a method of Service that makes an API call.
func (s *Service) handleGetEvent(arguments GetEventArguments) (*mcp_golang.ToolResponse, error) {
	// Construct the URL for the event endpoint
	eventURL := fmt.Sprintf("%s/event", s.baseURL) // Assuming an /event endpoint

	req, err := http.NewRequest("GET", eventURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+s.apiKey)
	req.Header.Add("Content-Type", "application/json") // Assuming JSON response

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API call to %s: %w", eventURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call to %s returned non-200 status: %d, body: %s", eventURL, resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// You might want to parse the JSON response here and return a structured content.
	// For now, returning raw JSON as text content.
	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(bodyBytes))), nil
}

type GetWidgetsArguments struct{}

// handleGetWidgets is a method of Service that makes an API call.
func (s *Service) handleGetWidgets(arguments GetWidgetsArguments) (*mcp_golang.ToolResponse, error) {
	// Construct the URL for the widgets endpoint
	widgetsURL := fmt.Sprintf("%s/widgets", s.baseURL) // Assuming a /widgets endpoint

	req, err := http.NewRequest("GET", widgetsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+s.apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API call to %s: %w", widgetsURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call to %s returned non-200 status: %d, body: %s", widgetsURL, resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(bodyBytes))), nil
}

type GetChatArguments struct{}

// handleGetChat is a method of Service that makes an API call.
func (s *Service) handleGetChat(arguments GetChatArguments) (*mcp_golang.ToolResponse, error) {
	// Construct the URL for the chat endpoint
	chatURL := fmt.Sprintf("%s/chat", s.baseURL) // Assuming a /chat endpoint

	req, err := http.NewRequest("GET", chatURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+s.apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API call to %s: %w", chatURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call to %s returned non-200 status: %d, body: %s", chatURL, resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(bodyBytes))), nil
}

type GetSingleWidgetArguments struct {
	ID string `json:"id"`
}

// handleGetSingleWidget is a method of Service that makes an API call.
func (s *Service) handleGetSingleWidget(arguments GetSingleWidgetArguments) (*mcp_golang.ToolResponse, error) {
	// Construct the URL for the widget endpoint
	widgetURL := fmt.Sprintf("%s/widget/%s", s.baseURL, arguments.ID)

	req, err := http.NewRequest("GET", widgetURL, nil)
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

type DeleteWidgetArguments struct {
	ID string `json:"id"`
}

func (s *Service) handleDeleteWidget(arguments DeleteWidgetArguments) (*mcp_golang.ToolResponse, error) {
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
