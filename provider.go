package main

import (
	"bytes"
	"html/template"
	"os"
)

func getProviderURL() string {
	return os.Getenv("PROVIDER_URL")
}

func RenderTemplate(tmpl *template.Template, obj interface{}) string {
	// Create a buffer to store the rendered output
	var resultBuffer bytes.Buffer

	// Render the template and store the result in the buffer
	err := tmpl.Execute(&resultBuffer, obj)
	if err != nil {
		panic(err)
	}

	// Get the rendered output as a string from the buffer
	return resultBuffer.String()
}

func RenderProvider() string {
	p := Provider{URL: getProviderURL()}
	tmpl := template.Must(template.New("provider").Parse(provider_resource))

	return RenderTemplate(tmpl, p)
}
