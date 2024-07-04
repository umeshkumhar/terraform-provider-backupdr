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
	_ datasource.DataSource              = &cloudCredentialDataSource{}
	_ datasource.DataSourceWithConfigure = &cloudCredentialDataSource{}
)

// cloudCredentialDataSource is the data source implementation.
type cloudCredentialDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type cloudCredentialResourceModel struct {
	Usedefaultsa   types.Bool   `tfsdk:"usedefaultsa"`
	Immutable      types.Bool   `tfsdk:"immutable"`
	Cloudtype      types.String `tfsdk:"cloudtype"`
	Region         types.String `tfsdk:"region"`
	Serviceaccount types.String `tfsdk:"serviceaccount"`
	Domain         types.String `tfsdk:"domain"`
	Projectid      types.String `tfsdk:"projectid"`
	SrcID          types.Int64  `tfsdk:"srcid"`
	ClusterID      types.Int64  `tfsdk:"clusterid"`
	Endpoint       types.String `tfsdk:"endpoint"`
	Clientid       types.String `tfsdk:"clientid"`
	Name           types.String `tfsdk:"name"`
	ID             types.String `tfsdk:"id"`
	Href           types.String `tfsdk:"href"`
	Stale          types.Bool   `tfsdk:"stale"`
}

// NewCloudCredentialDataSource - Datasource for CloudCredential
func NewCloudCredentialDataSource() datasource.DataSource {
	return &cloudCredentialDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *cloudCredentialDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *cloudCredentialDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudcredential"
}

func (d *cloudCredentialDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source can be used to read information about a BackupDR Cloud Credential. It displays the cloud credential ID as shown in the Management console > Manage > Cloud Credentials page.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Provide the ID of the cloud credentials.",
			},
			"href": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the URL to access the storage pools in the management console.",
			},
			"stale": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: "It displays true or false if the data is synchronized with the management console or not.",
			},
			"clusterid": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: "It displays the backup/recovery appliance ID as shown in the Management console > Manage > Appliances page.",
			},
			"serviceaccount": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the service account associated with the cloud credential.",
			},
			"region": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the region where the cloud credential is created.",
			},
			"cloudtype": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the cloud type associated with the cloud credential.",
			},
			"projectid": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the project ID associated with the cloud credential.",
			},
			"domain": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the domain associated with the cloud credential.",
			},
			"name": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the name of the cloud credential.",
			},
			"endpoint": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the endpoint associated with the cloud credential.",
			},
			"clientid": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "It displays the client ID associated with the cloud credential.",
			},
			"srcid": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: "It displays the source ID on the appliance.",
			},
			"usedefaultsa": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: "It displays true or false.",
			},
			"immutable": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: "It displays the immutable values - true or false.",
			},
		},
	}
}

func (d *cloudCredentialDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state cloudCredentialResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	cc, res, err := d.client.DefaultApi.GetCredential(d.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read CloudCredential",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR CloudCredential",
			res.Status,
		)
	}

	// Map response body to model
	state = cloudCredentialResourceModel{
		ID:             types.StringValue(cc.Id),
		Href:           types.StringValue(cc.Href),
		Stale:          types.BoolValue(cc.Stale),
		ClusterID:      types.Int64Value(cc.ClusterId),
		Serviceaccount: types.StringValue(cc.Serviceaccount),
		Projectid:      types.StringValue(cc.Projectid),
		Region:         types.StringValue(cc.Region),
		Name:           types.StringValue(cc.Name),
		Usedefaultsa:   types.BoolValue(cc.Usedefaultsa),
		Immutable:      types.BoolValue(cc.Immutable),
		Cloudtype:      types.StringValue(cc.Cloudtype),
		Domain:         types.StringValue(cc.Domain),
		SrcID:          types.Int64Value(cc.SrcId),
		Endpoint:       types.StringValue(cc.Endpoint),
		Clientid:       types.StringValue(cc.Clientid),
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
