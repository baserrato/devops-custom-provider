package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

type devOpsDataModel struct {
	Id  types.String `tfsdk:"id"`
	Ops types.List   `tfsdk:"ops"`
	Dev types.List   `tfsdk:"dev"`
}

func (d *DevOpsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devops_data"
}

func (d *DevOpsDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		MarkdownDescription: "Dev stuff",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Required: true,
			},
			"dev": {
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
	var config devOpsDataModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newDevOps, err := d.client.GetDevOp(config.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	intermediateDev := []devModel{}
	intermediateOps := []opsModel{}

	for _, op := range newDevOps.Ops {
		var newEngineers []engineersModel
		for _, eng := range op.Engineers {
			newEngineers = append(newEngineers, engineersModel{
				Name:  types.StringValue(string(eng.Name)),
				Id:    types.StringValue(string(eng.Id)),
				Email: types.StringValue(string(eng.Email)),
			})
		}
		newOp := opsModel{
			Name: types.StringValue(string(op.Name)),
			Id:   types.StringValue(string(op.Id)),
		}

		_ = tfsdk.ValueFrom(ctx, newEngineers, types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"email": types.StringType,
			"id":    types.StringType,
			"name":  types.StringType,
		}}}, &newOp.Engineers)

		intermediateOps = append(intermediateOps, newOp)
	}

	_ = tfsdk.ValueFrom(ctx, intermediateOps, types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
		"id":   types.StringType,
		"name": types.StringType,
		"engineers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"email": types.StringType,
			"id":    types.StringType,
			"name":  types.StringType,
		}}},
	}}}, &config.Ops)

	for _, dev := range newDevOps.Devs {
		var newEngineers []engineersModel
		for _, eng := range dev.Engineers {
			newEngineers = append(newEngineers, engineersModel{
				Name:  types.StringValue(string(eng.Name)),
				Id:    types.StringValue(string(eng.Id)),
				Email: types.StringValue(string(eng.Email)),
			})
		}
		newDev := devModel{
			Name: types.StringValue(string(dev.Name)),
			Id:   types.StringValue(string(dev.Id)),
		}
		_ = tfsdk.ValueFrom(ctx, newEngineers, types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"email": types.StringType,
			"id":    types.StringType,
			"name":  types.StringType,
		}}}, &newDev.Engineers)

		intermediateDev = append(intermediateDev, newDev)
	}

	_ = tfsdk.ValueFrom(ctx, intermediateDev, types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
		"id":   types.StringType,
		"name": types.StringType,
		"engineers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"email": types.StringType,
			"id":    types.StringType,
			"name":  types.StringType,
		}}},
	}}}, &config.Dev)

	/*
		_ = tfsdk.ValueFrom(ctx, newDevOps.Devs, types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"id":   types.StringType,
			"name": types.StringType,
			"engineers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"email": types.StringType,
				"id":    types.StringType,
				"name":  types.StringType,
			}}},
		}}}, &config.Dev)
		_ = tfsdk.ValueFrom(ctx, newDevOps.Ops, types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"id":   types.StringType,
			"name": types.StringType,
			"engineers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"email": types.StringType,
				"id":    types.StringType,
				"name":  types.StringType,
			}}},
		}}}, &config.Ops)
	*/
	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
