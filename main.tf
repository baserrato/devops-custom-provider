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


resource "devops-bootcamp_engineer" "bobby" {
  name  = "Bobby"
  email = "Bobby@gmail.com"
}
resource "devops-bootcamp_engineer" "bob" {
  name  = "Bob"
  email = "Bob222@gmail.com"
}

resource "devops-bootcamp_engineer" "bobb" {
  name  = "Bobb"
  email = "B@gmail.com"
}

resource "devops-bootcamp_dev" "bengal" {
  name      = "bengal"
  engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}
resource "devops-bootcamp_ops" "ferrets" {
  name      = "ferrets"
  engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}
resource "devops-bootcamp_ops" "ferrets2" {
  name      = "ferrets2"
  engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}







