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
	_ datasource.DataSource              = &profileAllDataSource{}
	_ datasource.DataSourceWithConfigure = &profileAllDataSource{}
)

// profileAllDataSource is the data source implementation.
type profileAllDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type allProfileResourceModel struct {
	// Count types.Int64        `tfsdk:"count"`
	Items []profileResourceModel `tfsdk:"items"`
}

// NewProfileAllDataSource - Datasource for SLA Profile
func NewProfileAllDataSource() datasource.DataSource {
	return &profileAllDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *profileAllDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *profileAllDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_profiles"
}

func (d *profileAllDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// "count": schema.Int64Attribute{
			// 	Computed: true,
			// },
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
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
				},
			},
		},
	}
}

func (d *profileAllDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state allProfileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	slp, res, err := d.client.SLAProfileApi.ListSlps(d.authCtx, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLA Profiles",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLA Profiles",
			res.Status,
		)
	}

	var slps = []profileResourceModel{}
	// Map response body to model
	for _, v := range slp.Items {
		slpState := profileResourceModel{
			ID:              types.StringValue(v.Id),
			Href:            types.StringValue(v.Href),
			Name:            types.StringValue(v.Name),
			Description:     types.StringValue(v.Description),
			Srcid:           types.StringValue(v.Srcid),
			Clusterid:       types.StringValue(v.Clusterid),
			Cid:             types.StringValue(v.Cid),
			Performancepool: types.StringValue(v.Performancepool),
			Remotenode:      types.StringValue(v.Remotenode),
			Dedupasyncnode:  types.StringValue(v.Dedupasyncnode),
			Localnode:       types.StringValue(v.Localnode),
			Createdate:      types.Int64Value(v.Createdate),
			Modifydate:      types.Int64Value(v.Modifydate),
			Syncdate:        types.Int64Value(v.Syncdate),
			Stale:           types.BoolValue(v.Stale),
		}

		if v.Vaultpool != nil {
			slpState.Vaultpool = &profileDiskPoolResourceModel{
				ID:   types.StringValue(v.Vaultpool.Id),
				Href: types.StringValue(v.Vaultpool.Href),
				Name: types.StringValue(v.Vaultpool.Name),
			}

		}
		if v.Vaultpool2 != nil {
			slpState.Vaultpool2 = &profileDiskPoolResourceModel{
				ID:   types.StringValue(v.Vaultpool2.Id),
				Href: types.StringValue(v.Vaultpool2.Href),
				Name: types.StringValue(v.Vaultpool2.Name),
			}
		}
		if v.Vaultpool3 != nil {
			slpState.Vaultpool3 = &profileDiskPoolResourceModel{
				ID:   types.StringValue(v.Vaultpool3.Id),
				Href: types.StringValue(v.Vaultpool3.Href),
				Name: types.StringValue(v.Vaultpool3.Name),
			}
		}
		if v.Vaultpool4 != nil {
			slpState.Vaultpool4 = &profileDiskPoolResourceModel{
				ID:   types.StringValue(v.Vaultpool4.Id),
				Href: types.StringValue(v.Vaultpool4.Href),
				Name: types.StringValue(v.Vaultpool4.Name),
			}
		}
		slps = append(slps, slpState)
	}

	state.Items = slps

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
