package main

import (
	"html/template"
)

func RenderLogicalDevice() string {
	s := append([]Speed{}, Speed{Name: "1G", Count: 40})
	s = append(s, Speed{Name: "10G", Count: 10})
	ld := LogicalDevice{Name: "D1", Rows: 2, Columns: 25, Speeds: s}

	tmpl := template.Must(template.New("logicalDevice").Parse(ld_resource))
	return RenderTemplate(tmpl, ld)
}
