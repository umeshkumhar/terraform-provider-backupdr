// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/antihax/optional"
	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &profileResource{}
	_ resource.ResourceWithConfigure   = &profileResource{}
	_ resource.ResourceWithImportState = &profileResource{}
)

// NewProfileResource to create SLA Profiles
func NewProfileResource() resource.Resource {
	return &profileResource{}
}

// profileResource is the resource implementation.
type profileResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// Metadata returns the resource type name.
func (r *profileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_profile"
}

// Schema defines the schema for the resource.
func (r *profileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an SLA Profile.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"href": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"cid": schema.StringAttribute{
				Optional: true,
			},
			"performancepool": schema.StringAttribute{
				Optional: true,
			},
			"localnode": schema.StringAttribute{
				Optional: true,
			},
			"remotenode": schema.StringAttribute{
				Optional: true,
			},
			"dedupasyncnode": schema.StringAttribute{
				Computed: true,
			},

			// Computed fields
			"srcid": schema.StringAttribute{
				Computed: true,
			},
			"clusterid": schema.StringAttribute{
				Computed: true,
			},
			"modifydate": schema.Int64Attribute{
				Computed: true,
			},
			"createdate": schema.Int64Attribute{
				Computed: true,
			},
			"stale": schema.BoolAttribute{
				Computed: true,
			},
			"syncdate": schema.Int64Attribute{
				Computed: true,
			},

			"vaultpool": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"vaultpool2": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"vaultpool3": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"vaultpool4": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *profileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *profileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan profileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqSlp := backupdr.SlpRest{
		Name:            plan.Name.ValueString(),
		Description:     plan.Description.ValueString(),
		Cid:             plan.Cid.ValueString(),
		Performancepool: plan.Performancepool.ValueString(),
		Remotenode:      plan.Remotenode.ValueString(),
		Localnode:       plan.Localnode.ValueString(),
	}

	if plan.Vaultpool != nil {
		reqSlp.Vaultpool = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool.ID.ValueString(),
		}
	}

	if plan.Vaultpool2 != nil {
		reqSlp.Vaultpool2 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool2.ID.ValueString(),
		}
	}

	if plan.Vaultpool3 != nil {
		reqSlp.Vaultpool3 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool3.ID.ValueString(),
		}
	}

	if plan.Vaultpool4 != nil {
		reqSlp.Vaultpool4 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool4.ID.ValueString(),
		}
	}

	// Generate API request body from plan
	reqBody := backupdr.SLAProfileApiCreateSlpOpts{
		Body: optional.NewInterface(reqSlp),
	}

	// Create new slp
	respObject, _, err := r.client.SLAProfileApi.CreateSlp(r.authCtx, &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SLA Profle",
			"Could not create SLA Profile, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Dedupasyncnode = types.StringValue(respObject.Dedupasyncnode)
	plan.Remotenode = types.StringValue(respObject.Remotenode)
	plan.Localnode = types.StringValue(respObject.Localnode)
	plan.Performancepool = types.StringValue(respObject.Performancepool)

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(respObject.Id)
	plan.Href = types.StringValue(respObject.Href)
	plan.Clusterid = types.StringValue(respObject.Clusterid)
	plan.Srcid = types.StringValue(respObject.Srcid)
	plan.Stale = types.BoolValue(respObject.Stale)
	plan.Createdate = types.Int64Value(respObject.Createdate)
	plan.Modifydate = types.Int64Value(respObject.Modifydate)
	plan.Syncdate = types.Int64Value(respObject.Syncdate)

	if respObject.Vaultpool != nil {
		plan.Vaultpool.Name = types.StringValue(respObject.Vaultpool.Name)
		plan.Vaultpool.Href = types.StringValue(respObject.Vaultpool.Href)
	}

	if respObject.Vaultpool2 != nil {
		plan.Vaultpool2.Name = types.StringValue(respObject.Vaultpool2.Name)
		plan.Vaultpool2.Href = types.StringValue(respObject.Vaultpool2.Href)
	}
	if respObject.Vaultpool3 != nil {
		plan.Vaultpool3.Name = types.StringValue(respObject.Vaultpool3.Name)
		plan.Vaultpool3.Href = types.StringValue(respObject.Vaultpool3.Href)
	}
	if respObject.Vaultpool4 != nil {
		plan.Vaultpool4.Name = types.StringValue(respObject.Vaultpool4.Name)
		plan.Vaultpool4.Href = types.StringValue(respObject.Vaultpool4.Href)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *profileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state profileResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed values
	respObject, _, err := r.client.SLAProfileApi.GetSlp(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SLA Profile",
			"Could not read SLA Profile with ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state.ID = types.StringValue(respObject.Id)
	state.Href = types.StringValue(respObject.Href)
	state.Syncdate = types.Int64Value(respObject.Syncdate)
	state.Modifydate = types.Int64Value(respObject.Modifydate)
	state.Stale = types.BoolValue(respObject.Stale)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *profileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan profileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqSlp := backupdr.SlpRest{
		Name:            plan.Name.ValueString(),
		Description:     plan.Description.ValueString(),
		Cid:             plan.Cid.ValueString(),
		Performancepool: plan.Performancepool.ValueString(),
		Remotenode:      plan.Remotenode.ValueString(),
		Localnode:       plan.Localnode.ValueString(),
	}

	if plan.Vaultpool != nil {
		reqSlp.Vaultpool = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool.ID.ValueString(),
		}
	}

	if plan.Vaultpool2 != nil {
		reqSlp.Vaultpool2 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool2.ID.ValueString(),
		}
	}

	if plan.Vaultpool3 != nil {
		reqSlp.Vaultpool3 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool3.ID.ValueString(),
		}
	}

	if plan.Vaultpool4 != nil {
		reqSlp.Vaultpool4 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool4.ID.ValueString(),
		}
	}

	// Generate API request body from plan
	reqBody := backupdr.SLAProfileApiUpdateSlpOpts{
		Body: optional.NewInterface(reqSlp),
	}

	// Update existing order
	respObject, res, err := r.client.SLAProfileApi.UpdateSlp(r.authCtx, plan.ID.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SLA Profile",
			"An unexpected error occurred when updating the BackupDR SLA Profile, unexpected error: "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Update SLA Profile",
			"An unexpected error occurred when creating the BackupDR API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"BackupDR Client Error: "+res.Status,
		)
	}

	// Update resource state with updated items and timestamp
	plan.Dedupasyncnode = types.StringValue(respObject.Dedupasyncnode)
	plan.Remotenode = types.StringValue(respObject.Remotenode)
	plan.Localnode = types.StringValue(respObject.Localnode)
	plan.Performancepool = types.StringValue(respObject.Performancepool)

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(respObject.Id)
	plan.Href = types.StringValue(respObject.Href)
	plan.Clusterid = types.StringValue(respObject.Clusterid)
	plan.Srcid = types.StringValue(respObject.Srcid)
	plan.Stale = types.BoolValue(respObject.Stale)
	plan.Createdate = types.Int64Value(respObject.Createdate)
	plan.Modifydate = types.Int64Value(respObject.Modifydate)
	plan.Syncdate = types.Int64Value(respObject.Syncdate)

	if respObject.Vaultpool != nil {
		plan.Vaultpool.Name = types.StringValue(respObject.Vaultpool.Name)
		plan.Vaultpool.Href = types.StringValue(respObject.Vaultpool.Href)
	}

	if respObject.Vaultpool2 != nil {
		plan.Vaultpool2.Name = types.StringValue(respObject.Vaultpool2.Name)
		plan.Vaultpool2.Href = types.StringValue(respObject.Vaultpool2.Href)
	}
	if respObject.Vaultpool3 != nil {
		plan.Vaultpool3.Name = types.StringValue(respObject.Vaultpool3.Name)
		plan.Vaultpool3.Href = types.StringValue(respObject.Vaultpool3.Href)
	}
	if respObject.Vaultpool4 != nil {
		plan.Vaultpool4.Name = types.StringValue(respObject.Vaultpool4.Name)
		plan.Vaultpool4.Href = types.StringValue(respObject.Vaultpool4.Href)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *profileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state profileResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing SLA profile
	_, err := r.client.SLAProfileApi.DeleteSlp(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting SLA Profile",
			"Could not delete slp, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *profileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
