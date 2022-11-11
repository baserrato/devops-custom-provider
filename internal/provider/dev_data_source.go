package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &DevDataSource{}
var _ datasource.DataSourceWithConfigure = &DevDataSource{}

func NewDevDataSource() datasource.DataSource {
	return &DevDataSource{}
}

// EngineerDataSource defines the data source implementation.
type DevDataSource struct {
	client *Client
}

// EngineerDataSourceModel describes the data source data model.
type DevDataSourceModel struct {
	Devs []devModel `tfsdk:"devs"`
}
type devModel struct {
	Name      types.String     `tfsdk:"name"`
	Id        types.String     `tfsdk:"id"`
	Engineers []engineersModel `tfsdk:"engineers"`
}

func (d *DevDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dev_data"
}

func (d *DevDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
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
					"engineers": {
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
		},
	}, nil
}

func (d *DevDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	d.client = client
}

func (d *DevDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DevDataSourceModel

	devs, err := d.client.GetDevs()
	if devs == nil {
		return
	}
	/*
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Returned value from GetEngineers: %s.", engineers),
		)
	*/
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Engineers",
			err.Error(),
		)
		return
	}

	// Map response body to model

	for _, dev := range devs {
		devsState := devModel{
			Name: types.StringValue(dev.Name),
			Id:   types.StringValue(dev.Id),
		}
		for _, eng := range dev.Engineers {
			devsState.Engineers = append(devsState.Engineers, engineersModel{
				Name:  types.StringValue(string(eng.Name)),
				Id:    types.StringValue(string(eng.Id)),
				Email: types.StringValue(string(eng.Email)),
			})
		}

		state.Devs = append(state.Devs, devsState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
