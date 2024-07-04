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
	_ datasource.DataSource              = &profileDataSource{}
	_ datasource.DataSourceWithConfigure = &profileDataSource{}
)

// profileDataSource is the data source implementation.
type profileDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// NewProfileDataSource - Datasource for SLA Profile
func NewProfileDataSource() datasource.DataSource {
	return &profileDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *profileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *profileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_profile"
}

func (d *profileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source can be used to read information about a backup profile. It displays the resource profile ID as shown in the **Management console** > **Backup Plans** > **Profiles** page.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Provide the ID of the resource.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the name of the OnVault pool used for resource profile.",
			},
			"href": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the API URI for backup plan profile.",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the description for the resource profile.",
			},
			"stale": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the possible values true or false.",
			},
			"syncdate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the last sync date.",
			},
			"srcid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the source ID on the appliance.",
			},
			"cid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the ID of the cluster - It is not the same as cluster ID.",
			},
			"clusterid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the backup/recovery appliance ID.",
			},
			"performancepool": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the name of the snapshot (performance) pool. The default is act_per_pool000.",
			},
			"remotenode": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the remote backup/recovery appliance name, when two appliances are to be configured to replicate snapshot data between them.",
			},
			"localnode": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the primary backup/recovery appliance name.",
			},
			"modifydate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the date when the resource profile details are modified.",
			},
			"createdate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the date when the resource profile was created.",
			},
			"vaultpool": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the ID of the OnVault pool.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the OnVault pool used for resource profile.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for OnVault storage pool.",
					},
				},
			},
			"vaultpool2": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
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
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
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
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the ID of the OnVault pool 4.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the OnVault pool 4 used for resource profile.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for OnVault storage pool.",
					},
				},
			},
		},
	}
}

func (d *profileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state profileResourceModel
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
	slpState := profileResourceModel{
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
		slpState.Vaultpool = &profileDiskPoolResourceModel{
			ID:   types.StringValue(slp.Vaultpool.Id),
			Href: types.StringValue(slp.Vaultpool.Href),
			Name: types.StringValue(slp.Vaultpool.Name),
		}

	}
	if slp.Vaultpool2 != nil {
		slpState.Vaultpool2 = &profileDiskPoolResourceModel{
			ID:   types.StringValue(slp.Vaultpool2.Id),
			Href: types.StringValue(slp.Vaultpool2.Href),
			Name: types.StringValue(slp.Vaultpool2.Name),
		}
	}
	if slp.Vaultpool3 != nil {
		slpState.Vaultpool3 = &profileDiskPoolResourceModel{
			ID:   types.StringValue(slp.Vaultpool3.Id),
			Href: types.StringValue(slp.Vaultpool3.Href),
			Name: types.StringValue(slp.Vaultpool3.Name),
		}
	}
	if slp.Vaultpool4 != nil {
		slpState.Vaultpool4 = &profileDiskPoolResourceModel{
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
