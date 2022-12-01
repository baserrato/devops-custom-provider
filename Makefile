#makefile for custom terraform provider this is required for terraform plan
.PHONY: testacc clean

plan: clean main.tf fmt build init
	terraform plan

build: main.go generate
	go build -o terraform-provider-devops-bootcamp

generate: main.go
	go generate

fmt: main.tf
	terraform fmt

init: clean main.tf build
	
	#makes a directory including making gome directories should they not exist (-p)
	mkdir -p .plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/$$(go env GOOS)_$$(go env GOARCH)
	ln -s "${PWD}/terraform-provider-devops-bootcamp" "${PWD}/.plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/$$(go env GOOS)_$$(go env GOARCH)/terraform-provider-devops-bootcamp"
	terraform init -plugin-dir=./.plugin-cache/


clean: 
	rm -rf .plugin-cache .terraform .terraform.lock.hcl terraform-provider-devops-bootcamp

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
