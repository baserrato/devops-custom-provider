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
engineers = []
}

output "api_engineers" {
  value = data.devops-bootcamp_engineer_data.api
}

