package main

import (
	"html/template"
)

func RenderRackBasedTemplate() string {
	r := RackMapping{Name: "R1", Count: 2}
	racks := append([]RackMapping{}, r)
	rt := RackBasedTemplate{Name: "Stage3_Template_1", Encapsulation: "vxlan", SpineDeviceName: "D1", SpineDeviceCount: 2, Racks: racks}

	tmpl := template.Must(template.New("rack_based_template").Parse(rack_based_template_resource))

	return RenderTemplate(tmpl, rt)
}
