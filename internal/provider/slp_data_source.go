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
	_ datasource.DataSource              = &slpDataSource{}
	_ datasource.DataSourceWithConfigure = &slpDataSource{}
)

// slpDataSource is the data source implementation.
type slpDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// NewSlpDataSource - Datasource for SLA Profile
func NewSlpDataSource() datasource.DataSource {
	return &slpDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *slpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *slpDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_slp"
}

func (d *slpDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
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
			"stale": schema.BoolAttribute{
				Computed: true,
			},
			"syncdate": schema.Int64Attribute{
				Computed: true,
			},
			"srcid": schema.StringAttribute{
				Computed: true,
			},
			"cid": schema.StringAttribute{
				Computed: true,
			},
			"clusterid": schema.StringAttribute{
				Computed: true,
			},
			"performancepool": schema.StringAttribute{
				Computed: true,
			},
			"remotenode": schema.StringAttribute{
				Computed: true,
			},
			"dedupasyncnode": schema.StringAttribute{
				Computed: true,
			},
			"localnode": schema.StringAttribute{
				Computed: true,
			},
			"modifydate": schema.Int64Attribute{
				Computed: true,
			},
			"createdate": schema.Int64Attribute{
				Computed: true,
			},

			"vaultpool": schema.SingleNestedAttribute{
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
				},
			},
			"vaultpool2": schema.SingleNestedAttribute{
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
				},
			},
			"vaultpool3": schema.SingleNestedAttribute{
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
				},
			},
			"vaultpool4": schema.SingleNestedAttribute{
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
				},
			},
		},
	}
}

func (d *slpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state SlpResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	slp, res, err := d.client.SLAProfileApi.GetSlp(d.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLA Profile",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLA Profile",
			res.Status,
		)
	}

	// Map response body to model
	slpState := SlpResourceModel{
		ID:              types.StringValue(slp.Id),
		Href:            types.StringValue(slp.Href),
		Name:            types.StringValue(slp.Name),
		Description:     types.StringValue(slp.Description),
		Srcid:           types.StringValue(slp.Srcid),
		Clusterid:       types.StringValue(slp.Clusterid),
		Cid:             types.StringValue(slp.Cid),
		Performancepool: types.StringValue(slp.Performancepool),
		Remotenode:      types.StringValue(slp.Remotenode),
		Dedupasyncnode:  types.StringValue(slp.Dedupasyncnode),
		Localnode:       types.StringValue(slp.Localnode),
		Createdate:      types.Int64Value(slp.Createdate),
		Modifydate:      types.Int64Value(slp.Modifydate),
		Syncdate:        types.Int64Value(slp.Syncdate),
		Stale:           types.BoolValue(slp.Stale),
	}

	if slp.Vaultpool != nil {
		slpState.Vaultpool = &slpDiskPoolResourceModel{
			ID:   types.StringValue(slp.Vaultpool.Id),
			Href: types.StringValue(slp.Vaultpool.Href),
			Name: types.StringValue(slp.Vaultpool.Name),
		}

	}
	if slp.Vaultpool2 != nil {
		slpState.Vaultpool2 = &slpDiskPoolResourceModel{
			ID:   types.StringValue(slp.Vaultpool2.Id),
			Href: types.StringValue(slp.Vaultpool2.Href),
			Name: types.StringValue(slp.Vaultpool2.Name),
		}
	}
	if slp.Vaultpool3 != nil {
		slpState.Vaultpool3 = &slpDiskPoolResourceModel{
			ID:   types.StringValue(slp.Vaultpool3.Id),
			Href: types.StringValue(slp.Vaultpool3.Href),
			Name: types.StringValue(slp.Vaultpool3.Name),
		}
	}
	if slp.Vaultpool4 != nil {
		slpState.Vaultpool4 = &slpDiskPoolResourceModel{
			ID:   types.StringValue(slp.Vaultpool4.Id),
			Href: types.StringValue(slp.Vaultpool4.Href),
			Name: types.StringValue(slp.Vaultpool4.Name),
		}
	}

	state = slpState

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

}
