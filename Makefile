#makefile for custom terraform provider this is required for terraform plan
GOOS?=$$(go env GOOS)
GOARCH?=$$(go env GOARCH)

.PHONY: testacc clean init plan

plan: clean init provider devops-resource devops-datasource
	terraform -chdir=examples/allCombined init -plugin-dir=../../.plugin-cache
	terraform -chdir=examples/allCombined plan 

build: main.go generate
	go build -o terraform-provider-devops-bootcamp

generate: main.go
	go generate

fmt: main.tf
	terraform fmt

init: clean build
	#makes a directory including making gome directories should they not exist (-p)
	mkdir -p .plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/$(GOOS)_$(GOARCH)
	ln -s "${PWD}/terraform-provider-devops-bootcamp" "${PWD}/.plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/$(GOOS)_$(GOARCH)/terraform-provider-devops-bootcamp"

clean: 
	rm -rf .plugin-cache .terraform .terraform.lock.hcl terraform-provider-devops-bootcamp

provider:
	terraform -chdir=examples/provider init -plugin-dir=../../.plugin-cache
	terraform -chdir=examples/provider plan

engineer-resource: 
	terraform -chdir=examples/resources/Engineer init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/resources/Engineer plan

dev-resource: engineer-resource
	terraform -chdir=examples/resources/Dev init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/resources/Dev plan

ops-resource: dev-resource
	terraform -chdir=examples/resources/Ops init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/resources/Ops plan

devops-resource: ops-resource clean
	terraform -chdir=examples/resources/DevOps init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/resources/DevOps plan

engineer-datasource:
	terraform -chdir=examples/data-sources/Engineer init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/data-sources/Engineer plan

dev-datasource: engineer-datasource
	terraform -chdir=examples/data-sources/Dev init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/data-sources/Dev plan

ops-datasource: dev-datasource
	terraform -chdir=examples/data-sources/Ops init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/data-sources/Ops plan

devops-datasource: ops-datasource clean
	terraform -chdir=examples/data-sources/DevOps init -plugin-dir=../../../.plugin-cache/
	terraform -chdir=examples/data-sources/DevOps plan

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
