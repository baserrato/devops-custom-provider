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
var _ resource.Resource = &OpsResource{}
var _ resource.ResourceWithImportState = &OpsResource{}

func NewOpsResource() resource.Resource {
	return &OpsResource{}
}

// ExampleResource defines the resource implementation.
type OpsResource struct {
	client *Client
}

// ExampleResourceModel describes the resource data model.

func (r *OpsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ops"
}

func (r *OpsResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Ops stuff",
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required:            true,
				MarkdownDescription: "name for a Ops group",
				Type:                types.StringType,
			},
			"id": {
				Computed:            true,
				MarkdownDescription: "identifier for a ops group",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
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
		},
	}, nil
}

func (r *OpsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	r.client = client
}

func (r *OpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan opsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item Ops_Api
	item.Name = string(plan.Name.ValueString())
	for _, eng := range plan.Engineers {
		item.Engineers = append(item.Engineers, Engineer_Api{
			Name:  eng.Name.ValueString(),
			Id:    eng.Id.ValueString(),
			Email: eng.Email.ValueString(),
		})
	}
	newOp, err := r.client.CreateOp(item)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating dev",
			"Could not create dev, unexpected error: "+err.Error(),
		)
		return
	}
	var state opsModel
	state.Name = types.StringValue(newOp.Name)
	state.Id = types.StringValue(newOp.Id)
	for _, eng := range newOp.Engineers {
		state.Engineers = append(state.Engineers, engineersModel{
			Name:  types.StringValue(string(eng.Name)),
			Id:    types.StringValue(string(eng.Id)),
			Email: types.StringValue(string(eng.Email)),
		})
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *OpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state opsModel

	// Read Terraform prior state data into the model
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ops, err := r.client.GetOp(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Dev",
			"Could not read Dev with that Id "+state.Id.ValueString()+": "+err.Error(),
		)
		return
	}
	state.Name = types.StringValue(ops.Name)
	state.Id = types.StringValue(ops.Id)
	for _, eng := range ops.Engineers {
		state.Engineers = append(state.Engineers, engineersModel{
			Name:  types.StringValue(string(eng.Name)),
			Id:    types.StringValue(string(eng.Id)),
			Email: types.StringValue(string(eng.Email)),
		})
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *OpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	/*
		var plan *EngineerModel
		diags := req.Plan.Get(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}


			// Generate API request body from plan
			var items []Engineer_Api
			for _, item := range plan.Engineers {
				items = append(items, Engineer_Api{
					Name:  string(item.Name.ValueString()),
					Id:    string(item.Id.ValueString()),
					Email: string(item.Email.ValueString()),
				})
			}

			// Update existing order
			newEngineer, err := r.client.UpdateEngineer(items[0])
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Updating HashiCups Order",
					"Could not update order, unexpected error: "+err.Error(),
				)
				return
			}

			// Update resource state with updated items and timestamp
			plan.Engineers = []opModel{}
			plan.Engineers = append(plan.Engineers, opModel{
				Name:  types.StringValue(newEngineer.Name),
				Id:    types.StringValue(newEngineer.Id),
				Email: types.StringValue(newEngineer.Email),
			})
			//plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

			diags = resp.State.Set(ctx, plan)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	*/
}

func (r *OpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state opsModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteOps(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Ops",
			"Could not delete op, unexpected error: "+err.Error(),
		)
		return
	}

}

func (r *OpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
