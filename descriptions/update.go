package descriptions

/*
 */
var UpdateDescriptions = map[string]string{
	"update_event": `
	Update the details of an event. All fields are optional; only the fields you provide will be changed.

	Fields:
	{
		title?: string
		livestreamLink?: string
		location?: {
			latitude: float64   (must be between -90 and 90)
			longitude: float64  (must be between -180 and 180)
		}
	}
`,

	"update_widget": `
	Update an existing widget by specifying its widgetId and the new data for the widget. The data structure must match the widget's type.

	The data field should NOT be a string. It should be a JSON object that matches the widget's type.


	type PollOption{
		id: string
		text: string
		voterIds: []string
	}

	POLL WIDGET DATA SCHEMA: {
		type: "POLL" // string literal
		question: string
		options: []PollOption // Array of PollOption objects (not a string!!!)
		isActive: bool
	}

	Arguments:
	{
		"widgetId": string    The ID of the widget to update

		// Only one should be provided at a time
		"info": Info | undefined,
		"markdown": Markdown | undefined,
		"poll": Poll | undefined,

	}
`,
}
