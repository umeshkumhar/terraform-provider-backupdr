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
		MarkdownDescription: "This data source can be used to read information about all BackupDR Profiles. It displays the resource profile ID as shown in the Management console > Manage > Resource Profiles page.",
		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the ID of the resource.",
						},
						"name": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the name of the OnVault pool used for resource profile.",
						},
						"href": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the API URI for backup plan profile.",
						},
						"description": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the description for the resource profile",
						},
						"stale": schema.BoolAttribute{
							Computed: true,
							MarkdownDescription: "It displays the possible values true or false.",
						},
						"syncdate": schema.Int64Attribute{
							Computed: true,
							MarkdownDescription: "It displays the last sync date.",
						},
						"srcid": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the source ID on the appliance.",
						},
						"cid": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the ID of the cluster - It is not the same as cluster ID.",
						},
						"clusterid": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the backup/recovery appliance ID.",
						},
						"performancepool": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the name of the snapshot (performance) pool. The default is act_per_pool000.",
						},
						"remotenode": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the remote backup/recovery appliance name, when two appliances are to be configured to replicate snapshot data between them.",
						},
						"dedupasyncnode": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the dedupe async node name.",
						},
						"localnode": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: "It displays the primary backup/recovery appliance name.",
						},
						"modifydate": schema.Int64Attribute{
							Computed: true,
							MarkdownDescription: "It displays the date when the resource profile details are modified.",
						},
						"createdate": schema.Int64Attribute{
							Computed: true,
							MarkdownDescription: "It displays the date when the resource profile was created.",
						},

						"vaultpool": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the ID of the OnVault pool.",
								},
								"name": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the name of the OnVault pool used for resource profile.",
								},
								"href": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the API URI for OnVault storage pool",
								},
							},
						},
						"vaultpool2": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the ID of the OnVault pool 2.",
								},
								"name": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the name of the OnVault pool 2 used for resource profile.",
								},
								"href": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the API URI for OnVault storage pool.",
								},
							},
						},
						"vaultpool3": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the ID of the OnVault pool 3.",
								},
								"name": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the name of the OnVault pool 3 used for resource profile.",
								},
								"href": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the API URI for OnVault storage pool.",
								},
							},
						},
						"vaultpool4": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the ID of the OnVault pool 4.",
								},
								"name": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the name of the OnVault pool 4 used for resource profile.",
								},
								"href": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: "It displays the API URI for OnVault storage pool.",
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
