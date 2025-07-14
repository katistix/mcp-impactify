package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type EditingInfoWidgetData struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type EditingMarkdownWidgetData struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type EditingPollWidgetData struct {
	Type     string `json:"type"`
	IsActive bool   `json:"isActive"`
}

// Fix the Data field to be an interface{} (or any) instead of string
type UpdateWidgetArguments struct {
	WidgetID string `json:"widgetId"`

	Info EditingInfoWidgetData     `json:"info,omitempty"`
	Md   EditingMarkdownWidgetData `json:"md,omitempty"`
	Poll EditingPollWidgetData     `json:"poll,omitempty"`
}

func (s *Service) HandleUpdateWidget(arguments UpdateWidgetArguments) (*mcp_golang.ToolResponse, error) {
	updateWidgetUrl := fmt.Sprintf("%s/widget/%s", s.baseURL, arguments.WidgetID)

	var data interface{}
	if arguments.Info.Type != "" {
		data = arguments.Info
	} else if arguments.Md.Type != "" {
		data = arguments.Md
	} else if arguments.Poll.Type != "" {
		data = arguments.Poll
	} else {
		return nil, fmt.Errorf("no widget data provided")
	}

	bodyBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", updateWidgetUrl, io.NopCloser(bytes.NewReader(bodyBytes)))
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
		return nil, fmt.Errorf("API call to %s returned non-200 status: %d, body: %s", updateWidgetUrl, resp.StatusCode, string(respBodyBytes))
	}

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(respBodyBytes))), nil

}
