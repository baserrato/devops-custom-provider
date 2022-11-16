terraform {
  required_providers {
    devops-bootcamp = {
      source = "liatr.io/terraform/devops-bootcamp"
    }
  }
}

provider "devops-bootcamp" {
  endpoint = "http://localhost:8080"
}

data "devops-bootcamp_engineer_data" "api" {
}

output "api_engineers" {
  value = data.devops-bootcamp_engineer_data.api
}

data "devops-bootcamp_ops_data" "api" {
}

output "api_ops" {
  value = data.devops-bootcamp_ops_data.api
}

data "devops-bootcamp_dev_data" "api" {
}

output "api_dev" {
  value = data.devops-bootcamp_dev_data.api
}

