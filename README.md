# Apstrat Terraform Sample Code Generator 

To Create a main.tf for the apstra blueprint resource, Blue Print, with Logical Device, Rack, Template.

Currently this generator assumes the below provider and version,

```
apstra = {
	source = "Juniper/apstra"
	version = "0.25.0"
}
```
It can be modified based on the needs of the version.

## Resources it supports
This repo will try to generate the apstra resource(s), limited to Logical Device, Rack, Template, Blue Print, based on the object structure.

1. Added Logical Device resource.
2. Added Rack resource

## Command to execute 
Use the below command to generate all the resources (sample ones):

Below command will run and generates the sample terrform file.
```
go run .
```
To run with the Provider URL, 
```
PROVIDER_URL="https://admin:password@1.1.1.1:12345/" go run .
```
