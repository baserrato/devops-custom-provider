
### Create.Read.Update.Delete (CRUD)
Terraform providers work closely with the idea of CRUD operations. Meaning that the providers
try to classify API functionality into 4 groups, here are examples related to our API:
- C[reate]
- - Creates new resource: dev, ops, devops, engineers
- R[ead]
- - Read in an existing resource
- U[pdate]
- - Update an existing resource, however we can not affect the sub-resources (engineers) inside it
- D[elete]
- - Deletes an existing resource, however we can not affect those resource's children

### Schema
The terraform provider framework manages data through its schema. The schema is the structure of data
internal to the provider. This schema should loosely mirror the API data structure.[schema docs](https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas)

### File Hierarchy
![boiler plate file tree](file_tree.png)
### Look inside the internal directory
``` Go
    type ExampleDataSourceModel struct {
        ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
        Id                    types.String `tfsdk:"id"`
    }
```
- this code can be found [example_data_source.go](https://github.com/hashicorp/terraform-provider-scaffolding-framework/blob/95358ccbb8c897150c9cbe973c237448bdefcd4f/internal/provider/example_data_source.go)
- this is how data should be modled and stored in the provider
#### [example_resource.go](https://github.com/hashicorp/terraform-provider-scaffolding-framework/blob/95358ccbb8c897150c9cbe973c237448bdefcd4f/internal/provider/example_resource.go)
- This files shows use of crud operations
- As well as metadata and GetSchema which are important for documentation purposes
