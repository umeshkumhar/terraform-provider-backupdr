package provider

import (
	"context"
	"strconv"

	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &applianceDataSource{}
	_ datasource.DataSourceWithConfigure = &applianceDataSource{}
)

// applianceDataSource is the data source implementation.
type applianceDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type appliancesResourceModel struct {
	Clusterid       types.String `tfsdk:"clusterid"`
	Serviceaccount  types.String `tfsdk:"serviceaccount"`
	Zone            types.String `tfsdk:"zone"`
	Region          types.String `tfsdk:"region"`
	Projectid       types.String `tfsdk:"projectid"`
	Version         types.String `tfsdk:"version"`
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	Ipaddress       types.String `tfsdk:"ipaddress"`
	Publicip        types.String `tfsdk:"publicip"`
	Secureconnect   types.Bool   `tfsdk:"secureconnect"`
	PkiBootstrapped types.Bool   `tfsdk:"pkibootstrapped"`
	Supportstatus   types.String `tfsdk:"supportstatus"`
	ID              types.String `tfsdk:"id"`
	Href            types.String `tfsdk:"href"`
	Stale           types.Bool   `tfsdk:"stale"`
}

// NewApplianceDataSource - Datasource for Appliance
func NewApplianceDataSource() datasource.DataSource {
	return &applianceDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *applianceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *applianceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance"
}

func (d *applianceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			"clusterid": schema.StringAttribute{
				Computed: true,
			},
			"serviceaccount": schema.StringAttribute{
				Computed: true,
			},
			"zone": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
			},
			"projectid": schema.StringAttribute{
				Computed: true,
			},
			"version": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Computed: true,
			},
			"publicip": schema.StringAttribute{
				Computed: true,
			},
			"supportstatus": schema.StringAttribute{
				Computed: true,
			},
			"secureconnect": schema.BoolAttribute{
				Computed: true,
			},
			"pkibootstrapped": schema.BoolAttribute{
				Computed: true,
			},
		},
	}
}

func (d *applianceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state appliancesResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	applianceID, convErr := strconv.ParseInt(state.ID.ValueString(), 10, 64)
	if convErr != nil {
		tflog.Error(ctx, "Error parsing appliance ID - "+convErr.Error())
	}
	appliance, res, err := d.client.ApplianceApi.GetCluster(d.authCtx, applianceID, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR Appliance",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR Appliance",
			res.Status,
		)
	}

	// Map response body to model
	state = appliancesResourceModel{
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
