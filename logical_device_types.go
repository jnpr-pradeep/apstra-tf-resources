package main

type Speed struct {
	Name string
	Count int
}
type LogicalDevice struct {
	Name string
	Rows int
	Columns int
	Speeds []Speed
}

const (
	ld_resource = `resource "apstra_logical_device" "{{.Name}}" {
	name = "{{.Name}}"
	panels = [
	{
		rows = {{.Rows}}
		columns = {{.Columns}}
		port_groups = [
		{{- range $speed := .Speeds }}
		{
			port_count = {{$speed.Count}}
			port_speed = "{{$speed.Name}}"
		},
		{{- end }}
		]
		}
	]
}`
)
