package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &EngineerDataSource{}
var _ datasource.DataSourceWithConfigure = &EngineerDataSource{}

func NewEngineerDataSource() datasource.DataSource {
	return &EngineerDataSource{}
}

// EngineerDataSource defines the data source implementation.
type EngineerDataSource struct {
	client *Client
}

// EngineerDataSourceModel describes the data source data model.
type EngineerDataSourceModel struct {
	Engineers []engineersModel `tfsdk:"engineers"`
}

type engineersModel struct {
	Name  types.String `tfsdk:"name"`
	Id    types.String `tfsdk:"id"`
	Email types.String `tfsdk:"email"`
}

func (d *EngineerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_engineer_data"
}

func (d *EngineerDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"engineers": {
				Computed: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type:     types.StringType,
						Computed: true,
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
		},
	}, nil
}

func (d *EngineerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	d.client = client
}

func (d *EngineerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EngineerDataSourceModel

	engineers, err := d.client.GetEngineers()
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
	for _, engineer := range engineers {
		engineerState := engineersModel{
			Name:  types.StringValue(engineer.Name),
			Id:    types.StringValue(engineer.Id),
			Email: types.StringValue(engineer.Email),
		}

		state.Engineers = append(state.Engineers, engineerState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
