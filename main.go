// main.go
package main

import (
	"log" // Added for logging errors in main

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

func main() {
	done := make(chan struct{})

	baseURL, apiKey, err := initializeWithHealthCheck()
	if err != nil {
		log.Fatalf("Failed to initialize with health check: %v", err) // Use log.Fatalf for critical errors
	}

	// Create an instance of your Service
	myService := NewService(baseURL, apiKey)

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// REGISTER THE TOOLS

	err = server.RegisterTool("get_event", ToolDescriptions["get_event"], myService.handleGetEvent)
	if err != nil {
		log.Fatalf("Failed to register get_event tool: %v", err)
	}
	err = server.RegisterTool("get_widgets", ToolDescriptions["get_widgets"], myService.handleGetWidgets)
	if err != nil {
		log.Fatalf("Failed to register get_widgets tool: %v", err)
	}
	err = server.RegisterTool("get_chat", ToolDescriptions["get_chat"], myService.handleGetChat)
	if err != nil {
		log.Fatalf("Failed to register get_chat tool: %v", err)
	}
	err = server.RegisterTool("get_single_widget", ToolDescriptions["get_single_widget"], myService.handleGetSingleWidget)
	if err != nil {
		log.Fatalf("Failed to register get_widget tool: %v", err)
	}

	// Start the server
	log.Println("Starting MCP server...") // Add a message before starting
	err = server.Serve()
	if err != nil {
		log.Fatalf("Server stopped with an error: %v", err)
	}

	<-done
}
