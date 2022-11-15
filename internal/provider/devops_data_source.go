package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &DevOpsDataSource{}
var _ datasource.DataSourceWithConfigure = &DevOpsDataSource{}

func NewDevOpsDataSource() datasource.DataSource {
	return &DevOpsDataSource{}
}

// EngineerDataSource defines the data source implementation.
type DevOpsDataSource struct {
	client *Client
}

// EngineerDataSourceModel describes the data source data model.
type DevOpsDataSourceModel struct {
	DevOps []devOpsModel `tfsdk:"devops"`
}
type devOpsModel struct {
	Id  types.String `tfsdk:"id"`
	Dev []devModel   `tfsdk:"dev"`
	Ops []opsModel   `tfsdk:"ops"`
}

func (d *DevOpsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devops_data"
}

func (d *DevOpsDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"devops": {
				Computed: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Type:     types.StringType,
						Computed: true,
					},
					"devs": {
						Computed: true,
						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
							"name": {
								Type:     types.StringType,
								Computed: true,
							},
							"id": {
								Type:     types.StringType,
								Computed: true,
							},
							"dev_engineers": {
								Required: true,
								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
									"name": {
										Type:     types.StringType,
										Required: true,
									},
									"id": {
										Type:     types.StringType,
										Required: true,
									},
									"email": {
										Type:     types.StringType,
										Required: true,
									},
								}),
							},
						}),
					},
					"ops": {
						Computed: true,
						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
							"name": {
								Type:     types.StringType,
								Computed: true,
							},
							"id": {
								Type:     types.StringType,
								Computed: true,
							},
							"ops_engineers": {
								Required: true,
								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
									"name": {
										Type:     types.StringType,
										Required: true,
									},
									"id": {
										Type:     types.StringType,
										Required: true,
									},
									"email": {
										Type:     types.StringType,
										Required: true,
									},
								}),
							},
						}),
					},
				}),
			},
		},
	}, nil
}

func (d *DevOpsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	d.client = client
}

func (d *DevOpsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	//var state DevOpsDataSourceModel

	devops, _ := d.client.GetDevOps()
	struc, _ := PrettyStruct(devops)
	resp.Diagnostics.AddError(
		"Unexpected Resource Configure Type",
		fmt.Sprintf("got: %s.", struc),
	)
	/*
	   	if err != nil {
	   		resp.Diagnostics.AddError(
	   			"Unable to Read Engineers",
	   			err.Error(),
	   		)
	   		return
	   	}

	   // Map response body to model

	   	for _, devop := range devops {
	   		devOpsState := devOpsModel{
	   			Id: types.StringValue(devop.Id),
	   		}
	   		for _, dev := range devop.Devs {
	   			devOpsState.Dev = append(devOpsState.Dev, devModel{
	   				Name: types.StringValue(string(dev.Name)),
	   				Id:   types.StringValue(string(dev.Id)),
	   			})
	   		}

	   		state.DevOps = append(state.DevOps, devOpsState)
	   	}

	   // Set state

	   diags := resp.State.Set(ctx, &state)
	   resp.Diagnostics.Append(diags...)

	   	if resp.Diagnostics.HasError() {
	   		return
	   	}
	*/
}
