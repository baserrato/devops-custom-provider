#makefile for custom terraform provider this is required for terraform plan
.PHONY: testacc clean init plan

plan: clean init devops-resource devops-datasource
	terraform -chdir=examples/allCombined plan 

build: main.go generate
	go build -o terraform-provider-devops-bootcamp

generate: main.go
	go generate

fmt: main.tf
	terraform fmt

init: clean build
	#makes a directory including making gome directories should they not exist (-p)
	mkdir -p .plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/$$(go env GOOS)_$$(go env GOARCH)
	ln -s "${PWD}/terraform-provider-devops-bootcamp" "${PWD}/.plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/$$(go env GOOS)_$$(go env GOARCH)/terraform-provider-devops-bootcamp"
	terraform -chdir=examples/allCombined init -plugin-dir=../../.plugin-cache/

clean: 
	rm -rf .plugin-cache .terraform .terraform.lock.hcl terraform-provider-devops-bootcamp

engineer-resource: 
	terraform -chdir=examples/resources/Engineer plan

dev-resource: engineer-resource 
	terraform -chdir=examples/resources/Dev plan

ops-resource: dev-resource
	terraform -chdir=examples/resources/Ops plan

devops-resource: ops-resource clean
	terraform -chdir=examples/resources/DevOps plan

devops-datasource:
	terraform -chdir=examples/resources/DevOps plan

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
