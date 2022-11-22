package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &DevOpsResource{}
var _ resource.ResourceWithImportState = &DevOpsResource{}

func NewDevOpsResource() resource.Resource {
	return &DevOpsResource{}
}

// ExampleResource defines the resource implementation.
type DevOpsResource struct {
	client *Client
}

// ExampleResourceModel describes the resource data model.
type devOpsModel struct {
	Id  types.String `tfsdk:"id"`
	Ops []opsModel   `tfsdk:"ops"`
	Dev []devModel   `tfsdk:"dev"`
}

func (r *DevOpsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devops"
}

func (r *DevOpsResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "DevOps stuff",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Computed:            true,
				MarkdownDescription: "identifier for a devops group",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"dev": {
				Required: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"id": {
						Type:     types.StringType,
						Optional: true,
					},
					"engineers": {
						Computed: true,
						PlanModifiers: tfsdk.AttributePlanModifiers{
							resource.UseStateForUnknown(),
						},
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
				Required: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"id": {
						Type:     types.StringType,
						Optional: true,
					},
					"engineers": {
						Computed: true,
						PlanModifiers: tfsdk.AttributePlanModifiers{
							resource.UseStateForUnknown(),
						},
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

func (r *DevOpsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	r.client = client
}

func (r *DevOpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan devOpsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item DevOps_Api

	for _, op := range plan.Ops {
		item.Ops = append(item.Ops, Ops_Api{
			Id: op.Id.ValueString(),
		})
	}

	for _, dev := range plan.Dev {
		item.Devs = append(item.Devs, Dev_Api{
			Id: dev.Id.ValueString(),
		})
	}

	newDevOps, err := r.client.CreateDevOps(item)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating dev",
			"Could not create dev, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Dev = []devModel{}
	plan.Ops = []opsModel{}
	plan.Id = types.StringValue(newDevOps.Id)
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

		plan.Ops = append(plan.Ops, newOp)
	}

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

		plan.Dev = append(plan.Dev, newDev)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *DevOpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state devOpsModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newDevOps, err := r.client.GetDevOp(state.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.Dev = []devModel{}
	state.Ops = []opsModel{}

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

		state.Ops = append(state.Ops, newOp)
	}

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

		state.Dev = append(state.Dev, newDev)
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DevOpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan devOpsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item DevOps_Api
	item.Id = string(plan.Id.ValueString())
	for _, op := range plan.Ops {
		item.Ops = append(item.Ops, Ops_Api{
			Id: op.Id.ValueString(),
		})
	}

	for _, dev := range plan.Dev {
		item.Devs = append(item.Devs, Dev_Api{
			Id: dev.Id.ValueString(),
		})
	}

	newDevOps, err := r.client.UpdateDevOps(item)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating devops",
			"Could not updating devops, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Dev = []devModel{}
	plan.Ops = []opsModel{}
	plan.Id = types.StringValue(newDevOps.Id)

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

		plan.Ops = append(plan.Ops, newOp)
	}

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

		plan.Dev = append(plan.Dev, newDev)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *DevOpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state devOpsModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteDevOps(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting HashiCups Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}

}

func (r *DevOpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
