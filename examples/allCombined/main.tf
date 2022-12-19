terraform {
  required_providers {
    devops-bootcamp = {
      source = "liatr.io/terraform/devops-bootcamp"
      version = "1.0.0"
    }
  }
}

provider "devops-bootcamp" {
  endpoint = "http://localhost:8080"
}


resource "devops-bootcamp_engineer" "bobby" {
  name  = "Bobby"
  email = "Bobby@gmail.com"
}
resource "devops-bootcamp_engineer" "bob" {
  name  = "Bob"
  email = "Bob222@gmail.com"
}


resource "devops-bootcamp_dev" "bengal" {
  name      = "bengal"
  engineers = [{ id = devops-bootcamp_engineer.bobby.id }]
}
resource "devops-bootcamp_ops" "ferrets" {
  name      = "ferrets"
  engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}
resource "devops-bootcamp_ops" "ferrets2" {
  name      = "ferrets2"
  engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}

resource "devops-bootcamp_devops" "devops" {
  ops = [{ id = devops-bootcamp_ops.ferrets.id }, { id = devops-bootcamp_ops.ferrets2.id }]
  dev = [{ id = devops-bootcamp_dev.bengal.id }]
}

data "devops-bootcamp_engineer_data" "bob" {
  name = "Bob"
}
output "engineer_bob" {
  value = data.devops-bootcamp_engineer_data.bob
}


data "devops-bootcamp_dev_data" "bengal" {
  name = "bengal"
}
output "dev_ferrets" {
  value = data.devops-bootcamp_dev_data.bengal
}


data "devops-bootcamp_ops_data" "ferrets" {
  name = "ferrets"
}
output "ops_ferrets" {
  value = data.devops-bootcamp_ops_data.ferrets
}


data "devops-bootcamp_devops_data" "devops" {
  id = devops-bootcamp_devops.devops.id
}
output "devops" {
  value = data.devops-bootcamp_devops_data.devops
}








