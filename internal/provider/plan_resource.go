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
	_ resource.Resource                = &planResource{}
	_ resource.ResourceWithConfigure   = &planResource{}
	_ resource.ResourceWithImportState = &planResource{}
)

// NewPlanResource to create SLA Profiles
func NewPlanResource() resource.Resource {
	return &planResource{}
}

// planResource is the resource implementation.
type planResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// Metadata returns the resource type name.
func (r *planResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plan"
}

// Schema defines the schema for the resource.
func (r *planResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Backup plans are the binding of templates ids, profiles ids to application ids management console shows when an application (VM, Filesystem or Database) is protected. Templates define how often to back up application data, how long to retain the application data backups, and where and how to replicate the application's data backups. Use the backup templates  to create policies and resource profiles to specify which backup/recovery appliance is used to manage the backup plan. A backup plans violation occurs when data is not being backed up according to the boundaries you have set in a templates policy. For more information, see [Backup plan](https://cloud.google.com/backup-disaster-recovery/docs/concepts/backup-plan).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The unique ID of this resource backup plan id can also be referred as sla ID’s.",
			},
			"href": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the API URI for backup plan.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provide a description for the backup plan.",
			},
			"stale": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays true or false if the data is synchronized with the management console or not.",
			},
			"syncdate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the last sync date.",
			},
			"logexpirationoff": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays true or false for log expirations. The default value is false.",
			},
			"dedupasyncoff": schema.StringAttribute{
				Computed: true,
			},
			"expirationoff": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the expiration schedule for application.",
			},
			"scheduleoff": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provide true or false values - to disable the backup plan set to true, else leave to false to ensure backups are enabled for the application on the defined schedule in the template.",
			},
			"modifydate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the date when the backup plan was last modified.",
			},
			"application": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Provide application details for the backup plan.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Provide the backup plan ID. It is also referred to as SLA ID.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the backup plan.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for backup plan.",
					},
					"description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the description of the backup plan.",
					},
					"appname": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the application name used for backup plan.",
					},
					"apptype": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the type of application used for backup plan.",
					},
					"stale": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible values true or false.",
					},
					"syncdate": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "It displays the last sync date.",
					},
				},
			},
			"slp": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Provide profile details for the backup plan.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Provide the resource backup plan profile ID.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the resource profile name.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for backup plan profile.",
					},
					"cid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the ID of the cluster. It is not the same as cluster ID.",
					},
					"stale": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible values true or false.",
					},
					"syncdate": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "It displays the last sync date.",
					},
				},
			},
			"slt": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Provide template details for the backup plan.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Provide the backup plan template ID.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the backup template name.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for backup plan template.",
					},
					"override": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays if you can override the backup plan settings or not. It can be true or false.",
					},
					"sourcename": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "it displays the source name. It normally matches the name string.",
					},
					"stale": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible values true or false.",
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *planResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *planResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan planResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqSla := backupdr.SlaRest{
		Description:      plan.Description.ValueString(),
		Logexpirationoff: plan.Logexpirationoff.ValueBool(),
		Dedupasyncoff:    plan.Dedupasyncoff.ValueString(),
		Expirationoff:    plan.Expirationoff.ValueString(),
		Scheduleoff:      plan.Scheduleoff.ValueString(),
	}

	if plan.Application != nil {
		reqSla.Application = &backupdr.ApplicationRest{
			Id: plan.Application.ID.ValueString(),
		}
	}

	if plan.Slp != nil {
		reqSla.Slp = &backupdr.SlpRest{
			Id: plan.Slp.ID.ValueString(),
		}
	}

	if plan.Slt != nil {
		reqSla.Slt = &backupdr.SltRest{
			Id: plan.Slt.ID.ValueString(),
		}
	}

	// Generate API request body from plan
	reqBody := backupdr.SLAApiCreateSlaOpts{
		Body: optional.NewInterface(reqSla),
	}

	// Create new sla
	respObject, _, err := r.client.SLAApi.CreateSla(r.authCtx, &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SLA",
			"Could not create SLA, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(respObject.Id)
	plan.Href = types.StringValue(respObject.Href)
	plan.Stale = types.BoolValue(respObject.Stale)
	plan.Modifydate = types.Int64Value(respObject.Modifydate)
	plan.Syncdate = types.Int64Value(respObject.Syncdate)
	// plan.Immutable = types.BoolValue(respObject.Immutable)

	plan.Expirationoff = types.StringValue(respObject.Expirationoff)
	plan.Dedupasyncoff = types.StringValue(respObject.Dedupasyncoff)
	plan.Logexpirationoff = types.BoolValue(respObject.Logexpirationoff)
	plan.Scheduleoff = types.StringValue(respObject.Scheduleoff)

	plan.Application.Appname = types.StringValue(respObject.Application.Appname)
	plan.Application.Apptype = types.StringValue(respObject.Application.Apptype)
	plan.Application.Description = types.StringValue(respObject.Application.Description)
	plan.Application.Href = types.StringValue(respObject.Application.Href)
	plan.Application.Name = types.StringValue(respObject.Application.Name)
	plan.Application.Stale = types.BoolValue(respObject.Application.Stale)
	plan.Application.Syncdate = types.Int64Value(respObject.Application.Syncdate)
	plan.Application.Href = types.StringValue(respObject.Application.Href)
	plan.Application.Href = types.StringValue(respObject.Application.Href)

	plan.Slp.Href = types.StringValue(respObject.Slp.Href)
	plan.Slp.Cid = types.StringValue(respObject.Slp.Cid)
	plan.Slp.Name = types.StringValue(respObject.Slp.Name)
	plan.Slp.Stale = types.BoolValue(respObject.Slp.Stale)
	plan.Slp.Syncdate = types.Int64Value(respObject.Slp.Syncdate)

	plan.Slt.Href = types.StringValue(respObject.Slt.Href)
	plan.Slt.Name = types.StringValue(respObject.Slt.Name)
	plan.Slt.Override = types.StringValue(respObject.Slt.Override)
	plan.Slt.Sourcename = types.StringValue(respObject.Slt.Sourcename)
	plan.Slt.Stale = types.BoolValue(respObject.Slt.Stale)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *planResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state planResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed values
	respObject, _, err := r.client.SLAApi.GetSla(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SLA",
			"Could not read SLA with ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state.ID = types.StringValue(respObject.Id)
	state.Href = types.StringValue(respObject.Href)
	state.Syncdate = types.Int64Value(respObject.Syncdate)
	state.Modifydate = types.Int64Value(respObject.Modifydate)
	state.Stale = types.BoolValue(respObject.Stale)

	state.Expirationoff = types.StringValue(respObject.Expirationoff)
	state.Dedupasyncoff = types.StringValue(respObject.Dedupasyncoff)
	state.Logexpirationoff = types.BoolValue(respObject.Logexpirationoff)
	state.Scheduleoff = types.StringValue(respObject.Scheduleoff)

	state.Application.Appname = types.StringValue(respObject.Application.Appname)
	state.Application.Apptype = types.StringValue(respObject.Application.Apptype)
	state.Application.Description = types.StringValue(respObject.Application.Description)
	state.Application.Href = types.StringValue(respObject.Application.Href)
	state.Application.Name = types.StringValue(respObject.Application.Name)
	state.Application.Stale = types.BoolValue(respObject.Application.Stale)
	state.Application.Syncdate = types.Int64Value(respObject.Application.Syncdate)
	state.Application.Href = types.StringValue(respObject.Application.Href)
	state.Application.Href = types.StringValue(respObject.Application.Href)

	state.Slp.Href = types.StringValue(respObject.Slp.Href)
	state.Slp.Cid = types.StringValue(respObject.Slp.Cid)
	state.Slp.Name = types.StringValue(respObject.Slp.Name)
	state.Slp.Stale = types.BoolValue(respObject.Slp.Stale)
	state.Slp.Syncdate = types.Int64Value(respObject.Slp.Syncdate)

	state.Slt.Href = types.StringValue(respObject.Slt.Href)
	state.Slt.Name = types.StringValue(respObject.Slt.Name)
	state.Slt.Override = types.StringValue(respObject.Slt.Override)
	state.Slt.Sourcename = types.StringValue(respObject.Slt.Sourcename)
	state.Slt.Stale = types.BoolValue(respObject.Slt.Stale)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *planResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan planResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqSla := backupdr.SlaRest{
		Id:               plan.ID.ValueString(),
		Description:      plan.Description.ValueString(),
		Logexpirationoff: plan.Logexpirationoff.ValueBool(),
		Dedupasyncoff:    plan.Dedupasyncoff.ValueString(),
		Expirationoff:    plan.Expirationoff.ValueString(),
	}

	if plan.Slp != nil {
		reqSla.Slp = &backupdr.SlpRest{
			Id: plan.Slp.ID.ValueString(),
		}
	}

	if plan.Slt != nil {
		reqSla.Slt = &backupdr.SltRest{
			Id: plan.Slt.ID.ValueString(),
		}
	}

	// Generate API request body from plan
	reqBody := backupdr.SLAApiUpdateSlaOpts{
		Body: optional.NewInterface(reqSla),
	}

	// Update existing order
	respObject, res, err := r.client.SLAApi.UpdateSla(r.authCtx, plan.ID.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SLA ",
			"An unexpected error occurred when updating the BackupDR SLA , unexpected error: "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Update SLA ",
			"An unexpected error occurred when creating the BackupDR API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"BackupDR Client Error: "+res.Status,
		)
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(respObject.Id)
	plan.Href = types.StringValue(respObject.Href)
	plan.Stale = types.BoolValue(respObject.Stale)
	plan.Modifydate = types.Int64Value(respObject.Modifydate)
	plan.Syncdate = types.Int64Value(respObject.Syncdate)

	plan.Expirationoff = types.StringValue(respObject.Expirationoff)
	plan.Dedupasyncoff = types.StringValue(respObject.Dedupasyncoff)
	plan.Logexpirationoff = types.BoolValue(respObject.Logexpirationoff)
	plan.Scheduleoff = types.StringValue(respObject.Scheduleoff)

	plan.Application.Appname = types.StringValue(respObject.Application.Appname)
	plan.Application.Apptype = types.StringValue(respObject.Application.Apptype)
	plan.Application.Description = types.StringValue(respObject.Application.Description)
	plan.Application.Href = types.StringValue(respObject.Application.Href)
	plan.Application.Name = types.StringValue(respObject.Application.Name)
	plan.Application.Stale = types.BoolValue(respObject.Application.Stale)
	plan.Application.Syncdate = types.Int64Value(respObject.Application.Syncdate)
	plan.Application.Href = types.StringValue(respObject.Application.Href)
	plan.Application.Href = types.StringValue(respObject.Application.Href)

	plan.Slp.Href = types.StringValue(respObject.Slp.Href)
	plan.Slp.Cid = types.StringValue(respObject.Slp.Cid)
	plan.Slp.Name = types.StringValue(respObject.Slp.Name)
	plan.Slp.Stale = types.BoolValue(respObject.Slp.Stale)
	plan.Slp.Syncdate = types.Int64Value(respObject.Slp.Syncdate)

	plan.Slt.Href = types.StringValue(respObject.Slt.Href)
	plan.Slt.Name = types.StringValue(respObject.Slt.Name)
	plan.Slt.Override = types.StringValue(respObject.Slt.Override)
	plan.Slt.Sourcename = types.StringValue(respObject.Slt.Sourcename)
	plan.Slt.Stale = types.BoolValue(respObject.Slt.Stale)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *planResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state planResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing SLA profile
	_, err := r.client.SLAApi.DeleteSla(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting SLA",
			"Could not delete sla, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *planResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
