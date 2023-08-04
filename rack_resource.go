package main

import (
	"html/template"
)

func RenderRack() string {
	r := Rack{Name: "R1", TORSwitchName: "D1", UplinkCount: 2, UplinkSpeed: "10G"}

	tmpl := template.Must(template.New("rack").Parse(rack_resource))

	return RenderTemplate(tmpl, r)
}
