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
	_ datasource.DataSource              = &cloudcredentialAllDataSource{}
	_ datasource.DataSourceWithConfigure = &cloudcredentialAllDataSource{}
)

// cloudcredentialAllDataSource is the data source implementation.
type cloudcredentialAllDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type allCloudCredentialsResourceModel struct {
	Items []cloudCredentialResourceModel `tfsdk:"items"`
}

// NewCloudcredentialAllDataSource - Datasource for CloudCredentials
func NewCloudcredentialAllDataSource() datasource.DataSource {
	return &cloudcredentialAllDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *cloudcredentialAllDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *cloudcredentialAllDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudcredentials"
}

func (d *cloudcredentialAllDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"href": schema.StringAttribute{
							Computed: true,
						},
						"stale": schema.BoolAttribute{
							Computed: true,
						},
						"clusterid": schema.Int64Attribute{
							Computed: true,
						},
						"serviceaccount": schema.StringAttribute{
							Computed: true,
						},
						"region": schema.StringAttribute{
							Computed: true,
						},
						"cloudtype": schema.StringAttribute{
							Computed: true,
						},
						"projectid": schema.StringAttribute{
							Computed: true,
						},
						"domain": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"endpoint": schema.StringAttribute{
							Computed: true,
						},
						"clientid": schema.StringAttribute{
							Computed: true,
						},
						"srcid": schema.Int64Attribute{
							Computed: true,
						},

						"usedefaultsa": schema.BoolAttribute{
							Computed: true,
						},
						"immutable": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *cloudcredentialAllDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state allCloudCredentialsResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	ccs, res, err := d.client.DefaultApi.ListCredentials(d.authCtx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR CloudCredentials",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR CloudCredentials",
			res.Status,
		)
	}

	var finalList = []cloudCredentialResourceModel{}
	// Map response body to model
	for _, cc := range ccs.Items {
		ccState := cloudCredentialResourceModel{
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
		finalList = append(finalList, ccState)
	}

	state.Items = finalList

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
