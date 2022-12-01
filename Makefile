#makefile for custom terraform provider this is required for terraform plan
# Run acceptance tests
.PHONY: testacc clean

default: plan

clean:
	rm -rf .plugin-cache .terraform .terraform.lock.hcl terraform-provider-devops-bootcamp

setup: clean main.tf build
	mkdir -p .plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/darwin_arm64
	ln -s "${PWD}/terraform-provider-devops-bootcamp" "${PWD}/.plugin-cache/liatr.io/terraform/devops-bootcamp/0.0.1/darwin_arm64/terraform-provider-devops-bootcamp"
	terraform init -plugin-dir=./.plugin-cache/
	terraform plan



plan: main.tf fmt build
	terraform plan

build: main.go generate
	go build -o terraform-provider-devops-bootcamp

generate: main.go
	go generate

fmt: main.tf
	terraform fmt

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
