// main.go
package main

import (
	"log" // Added for logging errors in main

	"github.com/katistix/mcp-impactify/descriptions"
	"github.com/katistix/mcp-impactify/service"
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
	myService := service.NewService(baseURL, apiKey)
	log.Println("Service initialized", baseURL)

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// REGISTER THE TOOLS

	// GET TOOLS
	err = server.RegisterTool("get_event", descriptions.GetDescriptions["get_event"], myService.HandleGetEvent)
	if err != nil {
		log.Fatalf("Failed to register get_event tool: %v", err)
	}
	err = server.RegisterTool("get_widgets", descriptions.GetDescriptions["get_widgets"], myService.HandleGetWidgets)
	if err != nil {
		log.Fatalf("Failed to register get_widgets tool: %v", err)
	}
	err = server.RegisterTool("get_chat", descriptions.GetDescriptions["get_chat"], myService.HandleGetChat)
	if err != nil {
		log.Fatalf("Failed to register get_chat tool: %v", err)
	}
	err = server.RegisterTool("get_single_widget", descriptions.GetDescriptions["get_single_widget"], myService.HandleGetSingleWidget)
	if err != nil {
		log.Fatalf("Failed to register get_widget tool: %v", err)
	}

	// UPDATE TOOLS
	err = server.RegisterTool("update_event", descriptions.UpdateDescriptions["update_event"], myService.HandleUpdateEvent)
	if err != nil {
		log.Fatalf("Failed to register update_widget tool: %v", err)
	}
	err = server.RegisterTool("update_widget", descriptions.UpdateDescriptions["update_widget"], myService.HandleUpdateWidget)
	if err != nil {
		log.Fatalf("Failed to register update_widget tool: %v", err)
	}

	// CREATE TOOLS
	err = server.RegisterTool("create_widget", descriptions.CreateDescriptions["create_widget"], myService.HandleCreateWidget)
	if err != nil {
		log.Fatalf("Failed to register create_widget tool: %v", err)
	}

	// DELETE TOOLS
	err = server.RegisterTool("delete_widget", descriptions.DeleteDescriptions["delete_widget"], myService.HandleDeleteWidget)
	if err != nil {
		log.Fatalf("Failed to register delete_widget tool: %v", err)
	}

	// Start the server
	log.Println("Starting MCP server...") // Add a message before starting
	err = server.Serve()
	if err != nil {
		log.Fatalf("Server stopped with an error: %v", err)
	}

	<-done
}
