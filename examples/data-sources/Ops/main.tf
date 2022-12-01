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

/* creating a populated ops resource so the datasource
   has something to grab.  You could also create a resource
   with a manual API call to accomplish the same thing
*/
resource "devops-bootcamp_engineer" "bobby" {
  name  = "Bobby"
  email = "Bobby@gmail.com"
}

resource "devops-bootcamp_ops" "bengal" {
  name      = "bengal"
  engineers = [{ id = devops-bootcamp_engineer.bobby.id }]
}

/* grabbing ops resource with name "bengal" */
data "devops-bootcamp_ops_data" "bengal" {
  name = "bengal"
}
output "dev_bengal" {
  value = data.devops-bootcamp_ops_data.bengal
}
