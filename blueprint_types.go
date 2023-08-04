package main

type BluePrint struct {
	Name         string
	TemplateName string
}

const (
	blueprint_resource = `resource "apstra_datacenter_blueprint" "{{.Name}}" {
	  name        = "{{.Name}}"
	  template_id = apstra_template_rack_based.{{.TemplateName}}.id
	}`
)
