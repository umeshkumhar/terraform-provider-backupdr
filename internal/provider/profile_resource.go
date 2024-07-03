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
		MarkdownDescription: "A resource profile specifies the storage media for backups of application and VM data. The template and the resource profile that make up the backup plan dictate the type of application data policies to perform and where to store the application data backups (which storage pool is used). Resource Profiles define which snapshot pool (if needed) is used and which remote appliance data is replicated. " +
			"In addition to templates, you also create resource profiles in the backup plans menu. Profiles define where to store data. Data can be stored in the following: " +
			"Primary Appliance: The backup/recovery appliance that the resource profile is created for. This includes selecting which appliance snapshot pool will be used. " +
			"Remote Appliance: The backup/recovery appliance used for remote replication. This remote appliance must be an appliance that is already paired to the selected local appliance. You can configure the remote appliance field only when one or more remote appliances are configured on the selected local appliance. See [Join appliance in non-sharing mode](https://cloud.google.com/backup-disaster-recovery/docs/concepts/join-appliance). " +
			"OnVault: Up to four object storage buckets defined by an OnVault storage pool. The OnVault pools store compressed and encrypted backups of application data on Google Cloud Storage. " +
			"For more information, see [Resource profile](https://cloud.google.com/backup-disaster-recovery/docs/concepts/resource-profile).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The ID of this resource.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Provide a name for the resource profile.",
			},
			"href": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the API URI for backup plan profile.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provide a description for the resource profile",
			},
			"cid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provide the ID of the cluster - It is not the same as cluster ID.",
			},
			"performancepool": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provide a name of the snapshot (performance) pool. The default is act_per_pool000.",
			},
			"localnode": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provide the primary backup/recovery appliance name.",
			},
			"remotenode": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provide the remote backup/recovery appliance name, when two appliances are to be configured to replicate snapshot data between them.",
			},
			"dedupasyncnode": schema.StringAttribute{
				Computed: true,
			},

			// Computed fields
			"srcid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the source ID on the appliance.",
			},
			"clusterid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the backup/recovery appliance ID.",
			},
			"modifydate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the date when the resource profile details are modified.",
			},
			"createdate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the date when the resource profile was created.",
			},
			"stale": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the possible values true or false.",
			},
			"syncdate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the last sync date.",
			},

			"vaultpool": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "It displays the ID of the OnVault pool.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the OnVault pool used for resource profile.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for OnVault storage pool",
					},
				},
			},
			"vaultpool2": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "It displays the ID of the OnVault pool 2.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the OnVault pool 2 used for resource profile.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for OnVault storage pool.",
					},
				},
			},
			"vaultpool3": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "It displays the ID of the OnVault pool 3.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the OnVault pool 3 used for resource profile.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for OnVault storage pool.",
					},
				},
			},
			"vaultpool4": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "It displays the ID of the OnVault pool 4.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the OnVault pool 4 used for resource profile.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays gthe API URI for OnVault storage pool.",
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
	} else {
		reqSlp.Vaultpool = &backupdr.DiskPoolRest{
			Id: "0",
		}
	}

	if plan.Vaultpool2 != nil {
		reqSlp.Vaultpool2 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool2.ID.ValueString(),
		}
	} else {
		reqSlp.Vaultpool2 = &backupdr.DiskPoolRest{
			Id: "0",
		}
	}

	if plan.Vaultpool3 != nil {
		reqSlp.Vaultpool3 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool3.ID.ValueString(),
		}
	} else {
		reqSlp.Vaultpool3 = &backupdr.DiskPoolRest{
			Id: "0",
		}
	}

	if plan.Vaultpool4 != nil {
		reqSlp.Vaultpool4 = &backupdr.DiskPoolRest{
			Id: plan.Vaultpool4.ID.ValueString(),
		}
	} else {
		reqSlp.Vaultpool4 = &backupdr.DiskPoolRest{
			Id: "0",
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

	if respObject.Vaultpool != nil || plan.Vaultpool != nil {
		if plan.Vaultpool.ID.ValueString() != "0" {
			plan.Vaultpool.Name = types.StringValue(respObject.Vaultpool.Name)
			plan.Vaultpool.Href = types.StringValue(respObject.Vaultpool.Href)
		} else {
			plan.Vaultpool.Name = types.StringNull()
			plan.Vaultpool.Href = types.StringNull()
		}
	}

	if respObject.Vaultpool2 != nil || plan.Vaultpool2 != nil {
		if plan.Vaultpool2.ID.ValueString() != "0" {
			plan.Vaultpool2.Name = types.StringValue(respObject.Vaultpool2.Name)
			plan.Vaultpool2.Href = types.StringValue(respObject.Vaultpool2.Href)
		} else {
			plan.Vaultpool2.Name = types.StringNull()
			plan.Vaultpool2.Href = types.StringNull()
		}
	}

	if respObject.Vaultpool3 != nil || plan.Vaultpool3 != nil {
		if plan.Vaultpool3.ID.ValueString() != "0" {
			plan.Vaultpool3.Name = types.StringValue(respObject.Vaultpool3.Name)
			plan.Vaultpool3.Href = types.StringValue(respObject.Vaultpool3.Href)
		} else {
			plan.Vaultpool3.Name = types.StringNull()
			plan.Vaultpool3.Href = types.StringNull()
		}
	}

	if respObject.Vaultpool4 != nil || plan.Vaultpool4 != nil {
		if plan.Vaultpool4.ID.ValueString() != "0" {
			plan.Vaultpool4.Name = types.StringValue(respObject.Vaultpool4.Name)
			plan.Vaultpool4.Href = types.StringValue(respObject.Vaultpool4.Href)
		} else {
			plan.Vaultpool4.Name = types.StringNull()
			plan.Vaultpool4.Href = types.StringNull()
		}
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
