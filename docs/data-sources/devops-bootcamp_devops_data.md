---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "devops-bootcamp_devops_data Data Source - devops-custom-provider"
subcategory: ""
description: |-
  Dev stuff
---

# devops-bootcamp_devops_data (Data Source)

Dev stuff



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `dev` (Attributes List) (see [below for nested schema](#nestedatt--dev))
- `id` (String) The ID of this resource.
- `ops` (Attributes List) (see [below for nested schema](#nestedatt--ops))

<a id="nestedatt--dev"></a>
### Nested Schema for `dev`

Read-Only:

- `engineers` (Attributes List) (see [below for nested schema](#nestedatt--dev--engineers))
- `id` (String)
- `name` (String)

<a id="nestedatt--dev--engineers"></a>
### Nested Schema for `dev.engineers`

Read-Only:

- `email` (String)
- `id` (String)
- `name` (String)



<a id="nestedatt--ops"></a>
### Nested Schema for `ops`

Read-Only:

- `engineers` (Attributes List) (see [below for nested schema](#nestedatt--ops--engineers))
- `id` (String)
- `name` (String)

<a id="nestedatt--ops--engineers"></a>
### Nested Schema for `ops.engineers`

Read-Only:

- `email` (String)
- `id` (String)
- `name` (String)


