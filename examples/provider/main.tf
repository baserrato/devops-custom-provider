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

resource "devops-bootcamp_engineer" "bob" {
  name  = "bob"
  email = "bob@bob.com"
}


data "devops-bootcamp_engineer_data" "bob2" {
  engineers = [{ name = "bob", email = "bob@bob.com" }, ]
}

