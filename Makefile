#makefile for custom terraform provider this is required for terraform plan

default: plan

plan: main.tf fmt generate
	terraform plan

generate: main.go
	go generate

fmt: main.tf
	terraform fmt

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
