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
	_ datasource.DataSource              = &applianceAllDataSource{}
	_ datasource.DataSourceWithConfigure = &applianceAllDataSource{}
)

// applianceAllDataSource is the data source implementation.
type applianceAllDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type allAppliancesResourceModel struct {
	Items []appliancesResourceModel `tfsdk:"items"`
}

// NewApplianceAllDataSource - Datasource for SLA Profile
func NewApplianceAllDataSource() datasource.DataSource {
	return &applianceAllDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *applianceAllDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *applianceAllDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliances"
}

func (d *applianceAllDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source can be used to read information about all backup/recovery Appliances. It displays the backup/recovery appliance ID as shown in the **Management console** > **Manage** > **Appliances** page.",
		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the ID of the appliance.",
						},
						"href": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the URL to access the storage pools in the management console.",
						},
						"stale": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the possible values true or false.",
						},
						"clusterid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the backup/recovery appliance ID as shown in the **Management console** > **Manage** > **Appliances** page.",
						},
						"serviceaccount": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the GCP service account used for backup/recovery appliances.",
						},
						"zone": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the zone where the appliance is located.",
						},
						"region": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the region where the backup/recovery appliance is created.",
						},
						"projectid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the project ID of the backup/recovery appliance ID.",
						},
						"version": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the version of the backup appliance.",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the name of the backup/recovery appliance ID.",
						},
						"type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the appliance type.",
						},
						"ipaddress": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the IP address of the backup/recovery appliance ID.",
						},
						"publicip": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the public IP of the backup/recovery appliance ID.",
						},
						"supportstatus": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the appliance up to date with latest patches or updates status. It can be true or false.",
						},
						"secureconnect": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the possible values for secure connect as true or false.",
						},
						"pkibootstrapped": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "It displays if the PKI boot strap is enabled or not.",
						},
					},
				},
			},
		},
	}
}

func (d *applianceAllDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state allAppliancesResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	appliances, res, err := d.client.ApplianceApi.ListClusters(d.authCtx, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR Appliances",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR Appliances",
			res.Status,
		)
	}

	var apps = []appliancesResourceModel{}
	// Map response body to model
	for _, appliance := range appliances.Items {
		applianceState := appliancesResourceModel{
			ID:              types.StringValue(appliance.Id),
			Href:            types.StringValue(appliance.Href),
			Stale:           types.BoolValue(appliance.Stale),
			Clusterid:       types.StringValue(appliance.Clusterid),
			Serviceaccount:  types.StringValue(appliance.Serviceaccount),
			Ipaddress:       types.StringValue(appliance.Ipaddress),
			Projectid:       types.StringValue(appliance.Projectid),
			Region:          types.StringValue(appliance.Region),
			Name:            types.StringValue(appliance.Name),
			Version:         types.StringValue(appliance.Version),
			Publicip:        types.StringValue(appliance.Publicip),
			Type:            types.StringValue(appliance.Type_),
			Zone:            types.StringValue(appliance.Zone),
			Secureconnect:   types.BoolValue(appliance.Secureconnect),
			Supportstatus:   types.StringValue(appliance.Supportstatus),
			PkiBootstrapped: types.BoolValue(appliance.PkiBootstrapped),
		}
		apps = append(apps, applianceState)
	}

	state.Items = apps

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
