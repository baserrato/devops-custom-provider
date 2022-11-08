package provider

import (
	"context"
	"fmt"
	"net/http"

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
	client *http.Client
}

// EngineerDataSourceModel describes the data source data model.
type EngineerDataSourceModel struct {
	Engineers []engineersModel `tfsdk:"engineers"`
}
type engineersModel struct {
	Id    types.Int64  `tfsdk:"Id"`
	Name  types.String `tfsdk:"Name"`
	Email types.String `tfsdk:"Email"`
}

func (d *EngineerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_engineer"
}

func (d *EngineerDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"engineers": {
				Computed: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"Id": {
						Type:     types.Int64Type,
						Computed: true,
					},
					"Name": {
						Type:     types.StringType,
						Computed: true,
					},
					"Email": {
						Type:     types.StringType,
						Computed: true,
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

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *EngineerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	/*
			var state EngineerDataSourceModel

		    engineers, err := d.client.GetCoffees()
		    if err != nil {
		        resp.Diagnostics.AddError(
		            "Unable to Read HashiCups Coffees",
		            err.Error(),
		        )
		        return
		    }

		    // Map response body to model
		    for _, engineer := range engineers {
		        coffeeState := engineersModel{
		            Id:          types.Int64Value(int64(engineer.Id)),
		            Name:        types.StringValue(engineer.Name),
		            Email:      types.StringValue(engineer.Email),
		        }

		        state.Engineers = append(state.Engineers, coffeeState)
		    }

		    // Set state
		    diags := resp.State.Set(ctx, &state)
		    resp.Diagnostics.Append(diags...)
		    if resp.Diagnostics.HasError() {
		        return
		    }
	*/
}
