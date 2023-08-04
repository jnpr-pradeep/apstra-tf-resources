package main

import (
	"fmt"
	"strings"
)

func main() {
	// It uses the builder pattern to construct the final object.
	var builder strings.Builder
	builder.WriteString(RenderProvider())
	builder.WriteString("\n")

	builder.WriteString(RenderLogicalDevice())
	builder.WriteString("\n")

	builder.WriteString(RenderRack())
	builder.WriteString("\n")

	builder.WriteString(RenderRackBasedTemplate())
	builder.WriteString("\n")

	builder.WriteString(RenderBluePrint())
	builder.WriteString("\n")

	fmt.Println(builder.String())
}
