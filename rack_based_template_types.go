package main

type RackMapping struct {
	Name  string
	Count int
}
type RackBasedTemplate struct {
	Name             string
	Encapsulation    string
	SpineDeviceName  string
	SpineDeviceCount int
	Racks            []RackMapping
}

const (
	rack_based_template_resource = `resource "apstra_template_rack_based" "{{.Name}}" {
		name                     = "Rack Based Template - {{.Name}}"
		overlay_control_protocol = "evpn"
	  	asn_allocation_scheme = "unique"
		spine = {
		  logical_device_id = apstra_logical_device.{{.SpineDeviceName}}.id
		  count             = {{.SpineDeviceCount}}
		}
		rack_infos = {
			{{- range $r := .Racks}}
		   (apstra_rack_type.{{$r.Name}}.id) = { count = {{$r.Count}} }
		   {{ end -}}
		}
	  }`
)
