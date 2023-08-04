package main

import (
	"html/template"
)

func RenderBluePrint() string {
	bp := BluePrint{Name: "Stage3_BluePrint_1", TemplateName: "Stage3_Template_1"}

	tmpl := template.Must(template.New("blueprint").Parse(blueprint_resource))

	return RenderTemplate(tmpl, bp)
}
