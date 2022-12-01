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


/*  Dev resource including an engineer*/
resource "devops-bootcamp_engineer" "bob" {
  name  = "Bob"
  email = "Bob222@gmail.com"
}

resource "devops-bootcamp_dev" "bengal" {
  name      = "bengal"
  engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}

/* Empty Dev resource */
resource "devops-bootcamp_dev" "ferrets" {
  name      = "ferrets"
  engineers = []
}






