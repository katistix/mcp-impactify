package main

// ToolDescriptions holds the descriptions for each tool, allowing multiline strings.
var ToolDescriptions = map[string]string{
	"get_event": `
Get details about the event that the api key associated with this instance of the server is authorized to access.

This tool returns information specific to the event context for the current API key.
`,
	"get_widgets": `
Get details about available widgets.

This tool provides a list and details of widgets accessible to the current event.
`,

	"get_chat": `
Get complete chat message history of the current event.
`,

	"get_single_widget": `
Get details about a single widget.

This tool provides details of a specific widget accessible to the current event.
`,
}
