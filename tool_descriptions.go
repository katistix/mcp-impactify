package main

// ToolDescriptions holds the descriptions for each tool, allowing multiline strings.
var ToolDescriptions = map[string]string{
	"get_event": `
Get details about the event that the api key associated with this instance of the server is authorized to access.

Those details include the event's title, location, and pricing information.

This tool returns information specific to the event context for the current API key.
`,
	"get_widgets": `
Get a list of details about all available widgets.

Widgets are the primary interface for interacting with the event.
They allow the staff members to post new information about the event such as updates, announcements, and other relevant content.
They can also include interactive features like polls.

This tool provides a list and details of widgets accessible to the current event.
`,

	"get_chat": `
Get complete chat message history of the current event.

The Message object includes the sender's user id, message content, and timestamp.

In the live chat, every member of the event can participate in real-time discussions.
`,

	"get_single_widget": `
Get details about a single widget.

This tool provides details of a specific widget accessible to the current event.
`,
}
