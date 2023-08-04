package main

type Rack struct {
	Name          string
	TORSwitchName string
	UplinkCount   int
	UplinkSpeed   string
}

const (
	rack_resource = `resource "apstra_rack_type" "{{.Name}}" {
	name                       = "{{.Name}}"
	description                = "Rack {{.Name}}"
	fabric_connectivity_design = "l3clos"
	leaf_switches = { // leaf switches are a map keyed by switch name, so
		tor_switch = { 
		logical_device_id   = apstra_logical_device.{{.TORSwitchName}}.id
		spine_link_count    = {{.UplinkCount}}
		spine_link_speed    = "{{.UplinkSpeed}}"
		}
	}
	}`
)
