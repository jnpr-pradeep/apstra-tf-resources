package main

type Provider struct {
	Name    string
	Version string
	URL     string
}

const (
	provider_resource = `terraform {
	required_providers {
	  apstra = {
		source = "Juniper/apstra"
		version = "0.25.0"
	  }
	}
  }
  
  provider "apstra" {
	url                     = "{{.URL}}"
	tls_validation_disabled = true                         # optional
  }`
)
