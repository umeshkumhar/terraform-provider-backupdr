package provider

import (
	"context"
	"fmt"
	"os"

	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &backupdrProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &backupdrProvider{
			version: version,
		}
	}
}

// backupdrProvider is the provider implementation.
type backupdrProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
	client  *backupdr.APIClient
	authCtx context.Context
}

// backupdrProviderModel maps provider schema data to a Go type.
type backupdrProviderModel struct {
	Endpoint       types.String `tfsdk:"endpoint"`
	GcpAccessToken types.String `tfsdk:"access_token"`
}

// Metadata returns the provider type name.
func (p *backupdrProvider) Metadata(ctx context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "backupdr"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *backupdrProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Required: true,
				Optional: false,
			},
			"access_token": schema.StringAttribute{
				Required: true,
				Optional: false,
			},
		},
	}
}

// Configure prepares a BackupDR API client for data sources and resources.
func (p *backupdrProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring BackupDR client")

	// Retrieve provider data from configuration
	var config backupdrProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown BackupDR API Endpoint",
			"The provider cannot create the BackupDR API client as there is an unknown configuration value for the BackupDR API Endpoint. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the BACKUPDR_ENDPOINT environment variable.",
		)
	}

	if config.GcpAccessToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("user"),
			"Unknown BackupDR API AccessToken",
			"The provider cannot create the BackupDR API client as there is an unknown configuration value for the BackupDR API AcessToken. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the BACKUPDR_ACCESS_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	endpoint := os.Getenv("BACKUPDR_ENDPOINT")
	accessToken := os.Getenv("BACKUPDR_ACCESS_TOKEN")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	if !config.GcpAccessToken.IsNull() {
		accessToken = config.GcpAccessToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Missing BackupDR API endpoint",
			"The provider cannot create the BackupDR API client as there is a missing or empty value for the BackupDR API endpoint. "+
				"Set the endpoint value in the configuration or use the BACKUPDR_ENDPOINT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if accessToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Unknown BackupDR API AccessToken",
			"The provider cannot create the BackupDR API client as there is a missing or empty value for the BackupDR API access_token. "+
				"Set the access_token value in the configuration or use the BACKUPDR_ACCESS_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "BACKUPDR_ENDPOINT", endpoint)
	ctx = tflog.SetField(ctx, "BACKUPDR_ACCESS_TOKEN", accessToken)

	tflog.Debug(ctx, "Creating BackupDR client")

	// define auth context with accessToken
	authCtx := context.WithValue(context.Background(), backupdr.ContextAccessToken, accessToken)

	// define client configuration object
	cfg := backupdr.NewConfiguration()
	cfg.Host = endpoint

	// define backupdr client using configuration object
	client := *backupdr.NewAPIClient(cfg)
	sessionObj, res, err := client.UserSessionApi.Login(authCtx)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Session for BackupDR API Client",
			"An unexpected error occurred when creating the BackupDR API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"BackupDR Client Error: "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Create Session for BackupDR API Client",
			"An unexpected error occurred when creating the BackupDR API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"BackupDR Client Error: "+res.Status,
		)
	}

	if sessionObj.SessionId == "" {
		resp.Diagnostics.AddError(
			"Unable to Create SessionID for BackupDR API Client",
			"An unexpected error occurred when creating the BackupDR API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"BackupDR SessionID "+fmt.Sprint(sessionObj),
		)
		return
	}

	p.authCtx = context.WithValue(authCtx, backupdr.ContextAPIKey, backupdr.APIKey{
		Key:    sessionObj.SessionId,
		Prefix: "Actifio",
	})
	p.client = &client

	// // Make the BackupDR client available during DataSource and Resource
	// // type Configure methods.
	resp.DataSourceData = p
	resp.ResourceData = p
	tflog.Info(ctx, "Configured BackupDR client", map[string]any{"success": true})

}

// DataSources defines the data sources implemented in the provider.
func (p *backupdrProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSltDataSource,
		NewSlpDataSource,
		NewDiskpoolDataSource,
		NewSLADataSource,
		NewSlpAllDataSource,
		NewApplianceDataSource,
		NewApplianceAllDataSource,
		NewCloudCredentialDataSource,
		NewCloudcredentialAllDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *backupdrProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSltResource,
		NewSlpResource,
		NewDiskpoolResource,
		NewSlaResource,
		NewVcenterHostResource,
		NewVcenterHostAddVMsResource,
		NewCloudAddVMsResource,
	}
}
