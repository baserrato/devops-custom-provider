package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDevOpsResource(t *testing.T) {
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
				  
resource "devops-bootcamp_dev" "test" {
	name      = "bengal"
	engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}
resource "devops-bootcamp_devops" "test" {
	dev = [{ id = devops-bootcamp_dev.test.id }]
	ops = []
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name
					resource.TestCheckResourceAttrSet("devops-bootcamp_devops.test", "id"),
					resource.TestCheckResourceAttr("devops-bootcamp_devops.test", "dev.#", "1"),
					resource.TestCheckResourceAttr("devops-bootcamp_devops.test", "ops.#", "0"),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_devops.test", "dev.*", map[string]string{
						"name": "bengal",
					}),

					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_devops.test", "dev.*.engineers.*", map[string]string{
						"name":  "Bob",
						"email": "Bob222@gmail.com",
					}),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig +
					`
resource "devops-bootcamp_engineer" "bob" {
	name  = "Bob"
	email = "Bob222@gmail.com"
}
resource "devops-bootcamp_engineer" "bobby" {
	name  = "bobby2"
	email = "bobby@gmail.com"
}
				  
resource "devops-bootcamp_dev" "test" {
	name      = "bengal2"
	engineers = [{ id = devops-bootcamp_engineer.bob.id }]
}

resource "devops-bootcamp_ops" "test" {
	name      = "bengalnew"
	engineers = [{ id = devops-bootcamp_engineer.bobby.id }]
}
resource "devops-bootcamp_devops" "test" {
	dev = [{ id = devops-bootcamp_dev.test.id }]
	ops = [{ id = devops-bootcamp_ops.test.id }]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name
					resource.TestCheckResourceAttrSet("devops-bootcamp_devops.test", "id"),
					resource.TestCheckResourceAttr("devops-bootcamp_devops.test", "dev.#", "1"),
					resource.TestCheckResourceAttr("devops-bootcamp_devops.test", "ops.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_devops.test", "dev.*", map[string]string{
						"name": "bengal2",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_devops.test", "ops.*", map[string]string{
						"name": "bengalnew",
					}),

					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_devops.test", "dev.*.engineers.*", map[string]string{
						"name":  "Bob",
						"email": "Bob222@gmail.com",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("devops-bootcamp_devops.test", "ops.*.engineers.*", map[string]string{
						"name":  "bobby2",
						"email": "bobby@gmail.com",
					}),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
