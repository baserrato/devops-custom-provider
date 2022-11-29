package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/liatrio/devops-bootcamp/examples/ch6/devops-resources"
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
type opsModel struct {
	Name      types.String `tfsdk:"name"`
	Id        types.String `tfsdk:"id"`
	Engineers types.List   `tfsdk:"engineers"`
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
				Required: true,
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
						Optional: true,
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

	engin := []engineersModel{}
	diags = plan.Engineers.ElementsAs(ctx, &engin, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var item devops_resource.Ops
	item.Name = string(plan.Name.ValueString())
	for _, eng := range engin {
		item.Engineers = append(item.Engineers, &devops_resource.Engineer{
			Id: eng.Id.ValueString(),
		})
	}
	newOp, err := r.client.CreateOp(item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating ops",
			"Could not create ops, unexpected error: "+err.Error(),
		)
		return
	}

	//plan.Engineers = //[]engineersModel{}
	intermediate := []engineersModel{}
	plan.Name = types.StringValue(newOp.Name)
	plan.Id = types.StringValue(newOp.Id)

	for _, eng := range newOp.Engineers {
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
	}}}, &plan.Engineers)

	diags = resp.State.Set(ctx, plan)
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
	op, err := r.client.GetOp(state.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		return
	}
	intermediate := []engineersModel{}
	state.Name = types.StringValue(op.Name)
	state.Id = types.StringValue(op.Id)

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
	}}}, &state.Engineers)
}

func (r *OpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan opsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	engin := []engineersModel{}
	diags = plan.Engineers.ElementsAs(ctx, &engin, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var item devops_resource.Ops
	item.Engineers = []*devops_resource.Engineer{}
	item.Name = string(plan.Name.ValueString())
	item.Id = string(plan.Id.ValueString())
	for _, eng := range engin {
		item.Engineers = append(item.Engineers, &devops_resource.Engineer{
			Id: eng.Id.ValueString(),
		})
	}
	newOp, err := r.client.UpdateOp(item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating ops",
			"Could not update ops, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Engineers = types.List{}
	intermediate := []engineersModel{}
	plan.Name = types.StringValue(newOp.Name)
	plan.Id = types.StringValue(newOp.Id)

	for _, eng := range newOp.Engineers {
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
	}}}, &plan.Engineers)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
func (r *OpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state opsModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteOp(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting ops",
			"Could not delete ops, unexpected error: "+err.Error(),
		)
		return
	}

}

func (r *OpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
