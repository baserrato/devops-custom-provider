package provider

import (
	"context"

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
				MarkdownDescription: "identifier for a dev group",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"dev": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"id": {
						Type:     types.StringType,
						Required: true,
					},
					"engineers": {
						Optional: true,
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
								Computed: true,
							},
						}),
					},
				}),
			},
			"ops": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"id": {
						Type:     types.StringType,
						Required: true,
					},
					"engineers": {
						Optional: true,
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
		var newEngineers []Engineer_Api
		for _, eng := range op.Engineers {
			newEngineers = append(newEngineers, Engineer_Api{
				Name:  eng.Name.ValueString(),
				Id:    eng.Id.ValueString(),
				Email: eng.Email.ValueString(),
			})
		}
		item.Ops = append(item.Ops, Ops_Api{
			Name:      op.Name.ValueString(),
			Id:        op.Id.ValueString(),
			Engineers: newEngineers,
		})
	}
	for _, dev := range plan.Dev {
		var newEngineers []Engineer_Api
		for _, eng := range dev.Engineers {
			newEngineers = append(newEngineers, Engineer_Api{
				Name:  eng.Name.ValueString(),
				Id:    eng.Id.ValueString(),
				Email: eng.Email.ValueString(),
			})
		}
		item.Devs = append(item.Devs, Dev_Api{
			Name:      dev.Name.ValueString(),
			Id:        dev.Id.ValueString(),
			Engineers: newEngineers,
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
		plan.Ops = append(plan.Ops, opsModel{
			Name:      types.StringValue(string(op.Name)),
			Id:        types.StringValue(string(op.Id)),
			Engineers: newEngineers,
		})
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
		plan.Dev = append(plan.Dev, devModel{
			Name:      types.StringValue(string(dev.Name)),
			Id:        types.StringValue(string(dev.Id)),
			Engineers: newEngineers,
		})
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
		resp.Diagnostics.AddError(
			"Error creating dev",
			"Could not create dev, unexpected error: "+err.Error(),
		)
		return
	}

	state.Dev = []devModel{}
	state.Ops = []opsModel{}
	state.Id = types.StringValue(newDevOps.Id)
	for _, op := range newDevOps.Ops {
		var newEngineers []engineersModel
		for _, eng := range op.Engineers {
			newEngineers = append(newEngineers, engineersModel{
				Name:  types.StringValue(string(eng.Name)),
				Id:    types.StringValue(string(eng.Id)),
				Email: types.StringValue(string(eng.Email)),
			})
		}
		state.Ops = append(state.Ops, opsModel{
			Name:      types.StringValue(string(op.Name)),
			Id:        types.StringValue(string(op.Id)),
			Engineers: newEngineers,
		})
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
		state.Dev = append(state.Dev, devModel{
			Name:      types.StringValue(string(dev.Name)),
			Id:        types.StringValue(string(dev.Id)),
			Engineers: newEngineers,
		})
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DevOpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan devModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item Dev_Api
	item.Name = string(plan.Name.ValueString())
	item.Id = string(plan.Id.ValueString())
	for _, eng := range plan.Engineers {
		item.Engineers = append(item.Engineers, Engineer_Api{
			Name:  eng.Name.ValueString(),
			Id:    eng.Id.ValueString(),
			Email: eng.Email.ValueString(),
		})
	}
	newDev, err := r.client.UpdateDev(item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating dev",
			"Could not updating dev, unexpected error: "+err.Error(),
		)
		return
	}
	plan.Engineers = []engineersModel{}
	plan.Name = types.StringValue(newDev.Name)
	plan.Id = types.StringValue(newDev.Id)
	for _, eng := range newDev.Engineers {
		plan.Engineers = append(plan.Engineers, engineersModel{
			Name:  types.StringValue(string(eng.Name)),
			Id:    types.StringValue(string(eng.Id)),
			Email: types.StringValue(string(eng.Email)),
		})
	}
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DevOpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state devModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteDev(state.Id.ValueString())
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
