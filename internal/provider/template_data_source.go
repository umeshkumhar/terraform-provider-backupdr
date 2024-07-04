package provider

import (
	"context"
	"strconv"

	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &templateDataSource{}
	_ datasource.DataSourceWithConfigure = &templateDataSource{}
)

// templateDataSource is the data source implementation.
type templateDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// NewTemplateDataSource - Datasource for SLA Template
func NewTemplateDataSource() datasource.DataSource {
	return &templateDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *templateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *templateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template"
}

func (d *templateDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source can be used to read information about a BackupDR Template. It displays the backup template ID as shown in the Management console > Manage > Backup Templates page.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				MarkdownDescription: "Provide the backup template ID.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				MarkdownDescription: "Name It displays the name of the backup template.",
			},
			"href": schema.StringAttribute{
				Computed:    true,
				MarkdownDescription: "It displays the API URI for Backup Plan template",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				MarkdownDescription: "It displays the description for the backup template.",
			},
			"managedbyagm": schema.BoolAttribute{
				Computed:    true,
				MarkdownDescription: "Managed by AGM",
			},
			"usedbycloudapp": schema.BoolAttribute{
				Computed:    true,
				MarkdownDescription: "Used by CloudApp It displays if the template is used by applications or not - true/false",
			},
			"option_href": schema.StringAttribute{
				Computed:    true,
				MarkdownDescription: "It displays the API URI for Backup Plan template options",
			},
			"policy_href": schema.StringAttribute{
				Computed:    true,
				MarkdownDescription: "It displays the backup policy ID.",
			},
			"sourcename": schema.StringAttribute{
				Computed:    true,
				MarkdownDescription: "Source Name It displays the source name. It should match the name value.",
			},
			"override": schema.StringAttribute{
				Computed:    true,
				MarkdownDescription: "Override options It displays the template override settings. Setting “Yes” will allow the policies set in this template to be overridden per-application. Setting “No” will enforce the policies as configured in this template without allowing any per-application overrides.",
			},
			"policies": schema.ListNestedAttribute{
				Computed:    true,
				MarkdownDescription: "List of policies (see below for nested schema)",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "Unique ID for this object",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the description for the backup policy.",
						},
						"encrypt": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the encryption identifier.",
						},
						"endtime": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the end time for the backup plan.",
						},
						"exclusion": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays specific days, days of week, month and days of month excluded for backup snapshots.",
						},
						"exclusioninterval": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the exclusion interval for the template. Normally set to 1.",
						},
						"exclusiontype": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the exclusion type as daily, weekly, monthly, or yearly.",
						},
						"href": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "Href of the policy",
						},
						"iscontinuous": schema.BoolAttribute{
							Computed:    true,
							MarkdownDescription: "It displays boolean value true or false if the policy setting for continuous mode or windowed",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "Name of the Policy",
						},
						"op": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the operation type. Normally set to snap, DirectOnVault, or stream_snap.",
						},
						"policytype": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the backup policy type. It can be snapshot, direct to OnVault, OnVault replication, mirror, and OnVault policy.",
						},
						"priority": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the application priority. It can be medium, high or low. The default job priority is medium, but you can change the priority to high or low.",
						},
						"remoteretention": schema.Int64Attribute{
							Computed:    true,
							MarkdownDescription: "It displays for mirror policy options.",
						},
						"repeatinterval": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the interval value. Normally set to 1.",
						},
						"reptype": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays for mirror policy options.",
						},
						"retention": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays how long the image is set for retention.",
						},
						"retentionm": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the retention in days, weeks, months, or years.",
						},
						"rpo": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays how often to run policy again. 24 is once per day.",
						},
						"rpom": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the PRP in hours. You can also set the RPO in minutes.",
						},
						"scheduletype": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the schedule type as daily, weekly, monthly or yearly.",
						},
						"selection": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the days to run the scheduled job. For example, weekly jobs on Sunday - days of week as sun.",
						},
						"sourcevault": schema.Int64Attribute{
							Computed:    true,
							MarkdownDescription: "It displays the OnVault disk pool id. You can get the from the management console > Manage > Storage Pools, then enabling visibility of the ID column.",
						},
						"starttime": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the start time for the backup plan in decimal format: total seconds = (hours x 3600) + (minutes + 60) + seconds",
						},
						"targetvault": schema.Int64Attribute{
							Computed:    true,
							MarkdownDescription: "It displays the OnVault disk pool id. You can get the from the management console > Manage > Storage Pools, then enabling visibility of the ID column.",
						},
						"truncatelog": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the Enable log truncation options. This may not work as required in advanced options.",
						},
						"verification": schema.BoolAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the verification values as true or false.",
						},
						"verifychoice": schema.StringAttribute{
							Computed:    true,
							MarkdownDescription: "It displays the empty value by default - to be used in future versions.",
						},
					},
				},
			},
		},
	}
}

func (d *templateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state templateResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	slt, res, err := d.client.SLATemplateApi.GetSlt(d.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLATemplate",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLATemplate",
			res.Status,
		)
		return
	}

	// Map response body to model
	sltState := templateResourceModel{
		ID:          types.StringValue(slt.Id),
		Href:        types.StringValue(slt.Href),
		Name:        types.StringValue(slt.Name),
		Description: types.StringValue(slt.Description),
		OptionHref:  types.StringValue(slt.OptionHref),
		PolicyHref:  types.StringValue(slt.PolicyHref),
		Sourcename:  types.StringValue(slt.Sourcename),
		Override:    types.StringValue(slt.Override),
	}

	// Fetch Policies for the given SLT
	sltID, _ := strconv.Atoi(state.ID.ValueString())
	sltPolicies, res, err := d.client.SLATemplateApi.ListPolicies(d.authCtx, int64(sltID))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLATemplate Policies",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR SLATemplate Policies",
			res.Status,
		)
		return
	}

	for _, pol := range sltPolicies.Items {
		sltState.Policies = append(sltState.Policies, policyRestModel{
			ID:                types.StringValue(pol.Id),
			Name:              types.StringValue(pol.Name),
			Description:       types.StringValue(pol.Description),
			Priority:          types.StringValue(pol.Priority),
			Rpom:              types.StringValue(pol.Rpom),
			Rpo:               types.StringValue(pol.Rpo),
			Exclusiontype:     types.StringValue(pol.Exclusiontype),
			Starttime:         types.StringValue(pol.Starttime),
			Endtime:           types.StringValue(pol.Endtime),
			Selection:         types.StringValue(pol.Selection),
			Scheduletype:      types.StringValue(pol.Scheduletype),
			Exclusion:         types.StringValue(pol.Exclusion),
			Reptype:           types.StringValue(pol.Reptype),
			Retention:         types.StringValue(pol.Retention),
			Retentionm:        types.StringValue(pol.Retentionm),
			Encrypt:           types.StringValue(pol.Encrypt),
			Repeatinterval:    types.StringValue(pol.Repeatinterval),
			Exclusioninterval: types.StringValue(pol.Exclusioninterval),
			PolicyType:        types.StringValue(pol.PolicyType),
			Truncatelog:       types.StringValue(pol.Truncatelog),
			Verifychoice:      types.StringValue(pol.Verifychoice),
			Op:                types.StringValue(pol.Op),
			Href:              types.StringValue(pol.Href),

			Remoteretention: types.Int64Value(int64(pol.Remoteretention)),
			Targetvault:     types.Int64Value(int64(pol.Targetvault)),
			Sourcevault:     types.Int64Value(int64(pol.Sourcevault)),

			Iscontinuous: types.BoolValue(pol.Iscontinuous),
			Verification: types.BoolValue(pol.Verification),
		})
	}

	sltState.Managedbyagm = types.BoolValue(slt.Managedbyagm)
	sltState.Usedbycloudapp = types.BoolValue(slt.Usedbycloudapp)

	state = sltState

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

}
