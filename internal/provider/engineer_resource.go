package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/liatrio/devops-bootcamp/examples/ch6/devops-resources"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &EngineerResource{}
var _ resource.ResourceWithImportState = &EngineerResource{}

func NewEngineerResource() resource.Resource {
	return &EngineerResource{}
}

// ExampleResource defines the resource implementation.
type EngineerResource struct {
	client *Client
}

// ExampleResourceModel describes the resource data model.
type engineersModel struct {
	Name  types.String `tfsdk:"name"`
	Id    types.String `tfsdk:"id"`
	Email types.String `tfsdk:"email"`
}

func (r *EngineerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_engineer"
}

func (r *EngineerResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Engineer stuff",
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required:            true,
				MarkdownDescription: "name for an Engineer",
				Type:                types.StringType,
			},
			"id": {
				Computed:            true,
				MarkdownDescription: "identifier for an Engineer",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"email": {
				Required:            true,
				MarkdownDescription: "email for an Engineer",
				Type:                types.StringType,
			},
		},
	}, nil
}

func (r *EngineerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	r.client = client
}

func (r *EngineerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan engineersModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item devops_resource.Engineer
	item.Name = string(plan.Name.ValueString())
	item.Email = string(plan.Email.ValueString())
	newEngineer, err := r.client.CreateEngineer(item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating engineer",
			"Could not create engineer, unexpected error: "+err.Error(),
		)
		return
	}
	plan.Name = types.StringValue(newEngineer.Name)
	plan.Id = types.StringValue(newEngineer.Id)
	plan.Email = types.StringValue(newEngineer.Email)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *EngineerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state engineersModel

	// Read Terraform prior state data into the model
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	engineer, err := r.client.GetEngineer(state.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.Name = types.StringValue(engineer.Name)
	state.Id = types.StringValue(engineer.Id)
	state.Email = types.StringValue(engineer.Email)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *EngineerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan engineersModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item devops_resource.Engineer 
	item.Name = string(plan.Name.ValueString())
	item.Id = string(plan.Id.ValueString())
	item.Email = string(plan.Email.ValueString())
	engineer, err := r.client.UpdateEngineer(item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating engineer",
			"Could not update engineer, unexpected error: "+err.Error(),
		)
		return
	}
	plan.Name = types.StringValue(engineer.Name)
	plan.Id = types.StringValue(engineer.Id)
	plan.Email = types.StringValue(engineer.Email)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *EngineerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state engineersModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteEngineer(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting HashiCups Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *EngineerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
