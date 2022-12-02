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


/* creating an engineer resource so the datasource
   has something to grab.  You could also create a resource
   with a manual API call to accomplish the same thing
*/
resource "devops-bootcamp_engineer" "bob" {
  name  = "Bob"
  email = "Bob222@gmail.com"
}

/* grabbing engineer resource with name "Bob" */
data "devops-bootcamp_engineer_data" "bob" {
  name = "Bob"
}
output "engineer_bob" {
  value = data.devops-bootcamp_engineer_data.bob
}
