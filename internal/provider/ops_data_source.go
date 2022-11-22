package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &OpsDataSource{}
var _ datasource.DataSourceWithConfigure = &OpsDataSource{}

func NewOpsDataSource() datasource.DataSource {
	return &OpsDataSource{}
}

// EngineerDataSource defines the data source implementation.
type OpsDataSource struct {
	client *Client
}

// EngineerDataSourceModel describes the data source data model.
type OpsDataSourceModel struct {
	Ops []opsModel `tfsdk:"ops"`
}

func (d *OpsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ops_data"
}

func (d *OpsDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"ops": {
				Computed: true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"id": {
						Type:     types.StringType,
						Computed: true,
					},
					"engineers": {
						Required: true,
						Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
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
		},
	}, nil
}

func (d *OpsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	d.client = client
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func (d *OpsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	/*
		var state OpsDataSourceModel

		ops, err := d.client.GetOps()
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Engineers",
				err.Error(),
			)
			return
		}

		// Map response body to model
		for _, op := range ops {
			opsState := opsModel{
				Name: types.StringValue(op.Name),
				Id:   types.StringValue(op.Id),
			}
			for _, eng := range op.Engineers {
				opsState.Engineers = append(opsState.Engineers, engineersModel{
					Name:  types.StringValue(string(eng.Name)),
					Id:    types.StringValue(string(eng.Id)),
					Email: types.StringValue(string(eng.Email)),
				})
			}

			state.Ops = append(state.Ops, opsState)
		}

		// Set state
		diags := resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	*/
}
