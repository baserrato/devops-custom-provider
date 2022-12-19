package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestOpsResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "devops-bootcamp_engineer" "bob" {
	name  = "Bob"
	email = "Bob222@gmail.com"
}
resource "devops-bootcamp_engineer" "bobby" {
	name  = "bobby"
	email = "bobby@gmail.com"
}
				  
resource "devops-bootcamp_ops" "test" {
	name      = "bengal"
	engineers = [{ id = devops-bootcamp_engineer.bob.id }, { id = devops-bootcamp_engineer.bobby.id }]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name
					resource.TestCheckResourceAttr("devops-bootcamp_ops.test", "name", "bengal"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("devops-bootcamp_ops.test", "id"),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_ops.test", "engineers.*", map[string]string{
						"name":  "Bob",
						"email": "Bob222@gmail.com",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_ops.test", "engineers.*", map[string]string{
						"name":  "bobby",
						"email": "bobby@gmail.com",
					}),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig +
					`
resource "devops-bootcamp_engineer" "bob" {	
	name  = "Bob2"
	email = "Bob@gmail.com"
}
resource "devops-bootcamp_engineer" "bobby" {
	name  = "bobby2"
	email = "bobby2@gmail.com"
}
				  
resource "devops-bootcamp_ops" "test" {
	name      = "bengal2"
	engineers = [{ id = devops-bootcamp_engineer.bob.id }, { id = devops-bootcamp_engineer.bobby.id }]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name
					resource.TestCheckResourceAttr("devops-bootcamp_ops.test", "name", "bengal2"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("devops-bootcamp_ops.test", "id"),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_ops.test", "engineers.*", map[string]string{
						"name":  "Bob2",
						"email": "Bob@gmail.com",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_ops.test", "engineers.*", map[string]string{
						"name":  "bobby2",
						"email": "bobby2@gmail.com",
					}),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
