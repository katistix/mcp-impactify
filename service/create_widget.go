package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type CreateInfoWidgetData struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type CreateMarkdownWidgetData struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PollOption struct {
	ID       string   `json:"id"`
	Text     string   `json:"text"`
	VoterIDs []string `json:"voterIds"`
}

type CreatePollWidgetData struct {
	Type     string       `json:"type"`
	Question string       `json:"question"`
	Options  []PollOption `json:"options"`
	IsActive bool         `json:"isActive"`
}

// Fix the Data field to be an interface{} (or any) instead of string
type CreateWidgetArguments struct {
	Info CreateInfoWidgetData     `json:"info,omitempty"`
	Md   CreateMarkdownWidgetData `json:"md,omitempty"`
	Poll CreatePollWidgetData     `json:"poll,omitempty"`
}

func (s *Service) HandleCreateWidget(arguments CreateWidgetArguments) (*mcp_golang.ToolResponse, error) {
	createWidgetUrl := fmt.Sprintf("%s/widget", s.baseURL)

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
	req, err := http.NewRequest("POST", createWidgetUrl, io.NopCloser(bytes.NewReader(bodyBytes)))
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
		return nil, fmt.Errorf("API call to %s returned non-200 status: %d, body: %s", createWidgetUrl, resp.StatusCode, string(respBodyBytes))
	}

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(respBodyBytes))), nil

}
