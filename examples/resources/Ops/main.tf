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


/*  Ops resource including an engineer*/
resource "devops-bootcamp_engineer" "johny" {
  name  = "johny"
  email = "johny@gmail.com"
}

resource "devops-bootcamp_ops" "bengal" {
  name      = "bengal"
  engineers = [{ id = devops-bootcamp_engineer.johny.id }]
}

/* Empty Ops resource */
resource "devops-bootcamp_ops" "ferrets" {
  name      = "ferrets"
  engineers = []
}















