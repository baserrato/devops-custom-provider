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
var _ resource.Resource = &DevResource{}
var _ resource.ResourceWithImportState = &DevResource{}

func NewDevResource() resource.Resource {
	return &DevResource{}
}

// ExampleResource defines the resource implementation.
type DevResource struct {
	client *Client
}

// ExampleResourceModel describes the resource data model.
type devModel struct {
	Name      types.String     `tfsdk:"name"`
	Id        types.String     `tfsdk:"id"`
	Engineers []engineersModel `tfsdk:"engineers"`
}

func (r *DevResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dev"
}

func (r *DevResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Dev stuff",
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required:            true,
				MarkdownDescription: "name for a dev group",
				Type:                types.StringType,
			},
			"id": {
				Computed:            true,
				MarkdownDescription: "identifier for a dev group",
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

func (r *DevResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*Client)

	r.client = client
}

func (r *DevResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan devModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item Dev_Api
	item.Name = string(plan.Name.ValueString())
	for _, eng := range plan.Engineers {
		item.Engineers = append(item.Engineers, Engineer_Api{
			Name:  eng.Name.ValueString(),
			Id:    eng.Id.ValueString(),
			Email: eng.Email.ValueString(),
		})
	}
	newDev, err := r.client.CreateDev(item)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating dev",
			"Could not create dev, unexpected error: "+err.Error(),
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

func (r *DevResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state devModel

	// Read Terraform prior state data into the model
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	dev, err := r.client.GetDev(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Dev",
			"Could not read Dev with that Id "+state.Id.ValueString()+": "+err.Error(),
		)
		return
	}
	state.Name = types.StringValue(dev.Name)
	state.Id = types.StringValue(dev.Id)
	state.Engineers = []engineersModel{}
	for _, eng := range dev.Engineers {
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

func (r *DevResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

func (r *DevResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *DevResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
