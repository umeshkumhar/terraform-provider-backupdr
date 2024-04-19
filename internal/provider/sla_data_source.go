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
	_ datasource.DataSource              = &slaDataSource{}
	_ datasource.DataSourceWithConfigure = &slaDataSource{}
)

// slaDataSource is the data source implementation.
type slaDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// NewSLADataSource - Datasource for SLA Profile
func NewSLADataSource() datasource.DataSource {
	return &slaDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *slaDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *slaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sla"
}

func (d *slaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"href": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"stale": schema.BoolAttribute{
				Computed: true,
			},
			"immutable": schema.BoolAttribute{
				Computed: true,
			},
			"syncdate": schema.Int64Attribute{
				Computed: true,
			},
			"logexpirationoff": schema.BoolAttribute{
				Computed: true,
			},
			"dedupasyncoff": schema.StringAttribute{
				Computed: true,
			},
			"expirationoff": schema.StringAttribute{
				Computed: true,
			},
			"scheduleoff": schema.StringAttribute{
				Computed: true,
			},

			"modifydate": schema.Int64Attribute{
				Computed: true,
			},

			"application": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"appname": schema.StringAttribute{
						Computed: true,
					},
					"apptype": schema.StringAttribute{
						Computed: true,
					},
					"stale": schema.BoolAttribute{
						Computed: true,
					},
					"syncdate": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"slp": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
					"stale": schema.BoolAttribute{
						Computed: true,
					},
					"syncdate": schema.Int64Attribute{
						Computed: true,
					},
					"cid": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"slt": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
					"override": schema.StringAttribute{
						Computed: true,
					},
					"sourcename": schema.StringAttribute{
						Computed: true,
					},
					"stale": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *slaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state SlaResourceModel
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
	state = SlaResourceModel{
		ID:               types.StringValue(sla.Id),
		Href:             types.StringValue(sla.Href),
		Description:      types.StringValue(sla.Description),
		Modifydate:       types.Int64Value(sla.Modifydate),
		Syncdate:         types.Int64Value(sla.Syncdate),
		Stale:            types.BoolValue(sla.Stale),
		Immutable:        types.BoolValue(sla.Immutable),
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

	state.Slp = &SlpResourceRefModel{
		ID:       types.StringValue(sla.Slp.Id),
		Href:     types.StringValue(sla.Slp.Href),
		Stale:    types.BoolValue(sla.Slp.Stale),
		Name:     types.StringValue(sla.Slp.Name),
		Syncdate: types.Int64Value(sla.Slp.Syncdate),
		Cid:      types.StringValue(sla.Slp.Cid),
	}

	state.Slt = &SltResourceRefModel{
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
