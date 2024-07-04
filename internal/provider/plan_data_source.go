package provider

import (
	"context"

	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &planDataSource{}
	_ datasource.DataSourceWithConfigure = &planDataSource{}
)

// planDataSource is the data source implementation.
type planDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// NewPlanDataSource - Datasource for SLA Profile
func NewPlanDataSource() datasource.DataSource {
	return &planDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *planDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *planDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plan"
}

func (d *planDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source can be used to read information about a backup plan. It displays the backup plan ID as shown in the **Management console** > **Manage** > **Backup Plans** page.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The unique ID of this resource backup plan ID can also be referred as sla ID’s.",
			},
			"href": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the API URI for backup plan.",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the backup plan description.",
			},
			"stale": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays true or false if the data is synchronized with the management console or not.",
			},
			"syncdate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the last sync date.",
			},
			"modifydate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the date when the backup plan was last modified.",
			},
			"logexpirationoff": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays true or false for log expirations. The default value is false.",
			},
			"dedupasyncoff": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the dedup async schedule for application.",
			},
			"expirationoff": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the expiration schedule for application.",
			},
			"scheduleoff": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the schedule for application.",
			},
			"application": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the application details for the backup plan.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the ID of the application.",
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
						MarkdownDescription: "It displays the application name used for the backup plan.",
					},
					"apptype": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the type of application used for the backup plan.",
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
				Computed:            true,
				MarkdownDescription: "It displays the profile details for the backup plan.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the ID of the backup plan.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the resource profile name.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for backup plan profile.",
					},
					"stale": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible values true or false.",
					},
					"syncdate": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "It displays the last sync date.",
					},
					"cid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the ID of the cluster. It is not the same as cluster ID.",
					},
				},
			},
			"slt": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the template details for the backup plan.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the ID of the backup template.",
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
						MarkdownDescription: "It displays the source name. It normally matches the name string.",
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

func (d *planDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state planResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	sla, res, err := d.client.SLAApi.GetSla(d.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLA",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLA",
			res.Status,
		)
	}

	// Map response body to model
	state = planResourceModel{
		ID:          types.StringValue(sla.Id),
		Href:        types.StringValue(sla.Href),
		Description: types.StringValue(sla.Description),
		Modifydate:  types.Int64Value(sla.Modifydate),
		Syncdate:    types.Int64Value(sla.Syncdate),
		Stale:       types.BoolValue(sla.Stale),
		// Immutable:        types.BoolValue(sla.Immutable),
		Dedupasyncoff:    types.StringValue(sla.Dedupasyncoff),
		Expirationoff:    types.StringValue(sla.Expirationoff),
		Scheduleoff:      types.StringValue(sla.Scheduleoff),
		Logexpirationoff: types.BoolValue(sla.Logexpirationoff),
	}
	state.Application = &ApplicationResourceModel{
		ID:          types.StringValue(sla.Application.Id),
		Href:        types.StringValue(sla.Application.Href),
		Description: types.StringValue(sla.Application.Description),
		Appname:     types.StringValue(sla.Application.Appname),
		Syncdate:    types.Int64Value(sla.Application.Syncdate),
		Apptype:     types.StringValue(sla.Application.Apptype),
		Stale:       types.BoolValue(sla.Application.Stale),
		Name:        types.StringValue(sla.Application.Name),
	}

	state.Slp = &profileResourceRefModel{
		ID:       types.StringValue(sla.Slp.Id),
		Href:     types.StringValue(sla.Slp.Href),
		Stale:    types.BoolValue(sla.Slp.Stale),
		Name:     types.StringValue(sla.Slp.Name),
		Syncdate: types.Int64Value(sla.Slp.Syncdate),
		Cid:      types.StringValue(sla.Slp.Cid),
	}

	state.Slt = &templateResourceRefModel{
		ID:         types.StringValue(sla.Slt.Id),
		Href:       types.StringValue(sla.Slt.Href),
		Sourcename: types.StringValue(sla.Slt.Sourcename),
		Override:   types.StringValue(sla.Slt.Override),
		Stale:      types.BoolValue(sla.Slt.Stale),
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
