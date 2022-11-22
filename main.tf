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
  engineers = []
}

resource "devops-bootcamp_devops" "devops" {
  dev = [{ id = devops-bootcamp_dev.bengal.id }]
  ops = []
}






