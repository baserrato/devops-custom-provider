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


/* creating a populated devops resource so the datasource
   has something to grab.  You could also create a resource
   with a manual API call to accomplish the same thing   */
resource "devops-bootcamp_engineer" "bobby" {
  name  = "Bobby"
  email = "Bobby@gmail.com"
}
resource "devops-bootcamp_ops" "ferrets" {
  name      = "ferrets"
  engineers = [{ id = devops-bootcamp_engineer.bobby.id }]
}
resource "devops-bootcamp_devops" "devops" {
  ops = [{ id = devops-bootcamp_ops.ferrets.id }]
  dev = []
}

/* grabbing data on resource with specific id, 
   devops resources don't have a name       */
data "devops-bootcamp_devops_data" "devops" {
  id = devops-bootcamp_devops.devops.id
}
output "devops" {
  value = data.devops-bootcamp_devops_data.devops
}








