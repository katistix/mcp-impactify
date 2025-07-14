package descriptions

var CreateDescriptions = map[string]string{
	"create_widget": `
Create a new widget by specifying its type and required properties. The arguments object must match ONE of the following formats, depending on the widget type:

Depending on the widget type, MAKE SURE TO PROVIDE ALL THE REQUIRED FIELDS AND PROPERTIES.

Arguments:
	{
		// Only one should be provided at a time
		"info": Info | undefined,
		"markdown": Markdown | undefined,
		"poll": Poll | undefined,
	}
`,
}
