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

func NewDevResource()  resource.Resource {
	return &DevResource{}
}

// ExampleResource defines the resource implementation.
type DevResource struct {
	client *Client
}

// ExampleResourceModel describes the resource data model.
type developerModel  struct {
	Name  types.String `tfsdk:"name"`
	Id    types.String `tfsdk:"id"`
	Email types.String `tfsdk:"email"`
}

type DevModel struct {
	Engineers []developerModel `tfsdk:"engineers"`
}

func (r *DevResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_engineer"
}

func (r *DevResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Engineer stuff",

		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Required:            true,
				MarkdownDescription: "name for an Developer",
				Type:                types.StringType,
			},
			"id": {
				Computed:            true,
				MarkdownDescription: "identifier for an Developer",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"email": {
				Required:            true,
				MarkdownDescription: "email for an Developer",
				Type:                types.StringType,
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
	var plan developerModel 
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item Dev_Api
	item.Name = string(plan.Name.ValueString())
	item.Id = string(plan.Id.Value)
	newEngineer, err := r.client.CreateDev(item)
 

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating engineer",
			"Could not create engineer, unexpected error: "+err.Error(),
		)
		return
	}
	plan.Name = types.StringValue(newEngineer.Name)
	plan.Id = types.StringValue(newEngineer.Id)
	plan.Id = types.StringValue(newEngineer.Id)  //ask ryan about this line

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DevResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	/*
		var data *EngineersModel

		// Read Terraform prior state data into the model
		resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

		if resp.Diagnostics.HasError() {
			return
		}
	*/
	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := d.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	//resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DevResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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
			plan.Engineers = []developerModel{}
			plan.Engineers = append(plan.Engineers, developerModel{
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

func (r *DevResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state developerModel 
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
