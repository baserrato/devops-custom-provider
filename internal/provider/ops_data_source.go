package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

func (d *OpsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ops_data"
}

func (d *OpsDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Ops stuff",
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required:            true,
				MarkdownDescription: "name for a ops group",
				Type:                types.StringType,
			},
			"id": {
				Computed:            true,
				MarkdownDescription: "identifier for a ops group",
				Type:                types.StringType,
			},
			"engineers": {
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
					"email": {
						Type:     types.StringType,
						Computed: true,
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
	var config opsModel

	// Read Terraform prior config data into the model
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	op, err := d.client.GetOpByName(config.Name.ValueString())
	if err != nil {
		diags = resp.State.Set(ctx, config)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		return
	}
	intermediate := []engineersModel{}
	config.Name = types.StringValue(op.Name)
	config.Id = types.StringValue(op.Id)

	for _, eng := range op.Engineers {
		intermediate = append(intermediate, engineersModel{
			Name:  types.StringValue(string(eng.Name)),
			Id:    types.StringValue(string(eng.Id)),
			Email: types.StringValue(string(eng.Email)),
		})
	}

	_ = tfsdk.ValueFrom(ctx, intermediate, types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
		"email": types.StringType,
		"id":    types.StringType,
		"name":  types.StringType,
	}}}, &config.Engineers)

	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
