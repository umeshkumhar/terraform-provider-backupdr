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
		Description: "Read details about SLA Template",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique ID for this object",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Name",
			},
			"href": schema.StringAttribute{
				Computed:    true,
				Description: "URL to access this object format",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "description",
			},
			"stale": schema.BoolAttribute{
				Computed:    true,
				Description: "Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.",
			},
			"syncdate": schema.Int64Attribute{
				Computed:    true,
				Description: "When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. format",
			},
			"managedbyagm": schema.BoolAttribute{
				Computed:    true,
				Description: "Managed by AGM",
			},
			"usedbycloudapp": schema.BoolAttribute{
				Computed:    true,
				Description: "Used by CloudApp",
			},
			"option_href": schema.StringAttribute{
				Computed:    true,
				Description: "Href for options",
			},
			"policy_href": schema.StringAttribute{
				Computed:    true,
				Description: "Href for policy",
			},
			"sourcename": schema.StringAttribute{
				Computed:    true,
				Description: "Source Name",
			},
			"override": schema.StringAttribute{
				Computed:    true,
				Description: "Override options",
			},
			"policies": schema.ListNestedAttribute{
				Description: "List of policies",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique ID for this object",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: "description",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the Policy",
						},
						"href": schema.StringAttribute{
							Computed:    true,
							Description: "Href of the policy",
						},
						"stale": schema.BoolAttribute{
							Computed:    true,
							Description: "Stale",
						},
						"syncdate": schema.Int64Attribute{
							Computed:    true,
							Description: "Last sync date",
						},
						"priority": schema.StringAttribute{
							Computed:    true,
							Description: "priority",
						},
						"rpo": schema.StringAttribute{
							Computed:    true,
							Description: "RPO",
						},
						"rpom": schema.StringAttribute{
							Computed:    true,
							Description: "RPO months",
						},
						"exclusiontype": schema.StringAttribute{
							Computed:    true,
							Description: "Exclusion Type",
						},
						"iscontinuous": schema.BoolAttribute{
							Computed:    true,
							Description: "Is Continous",
						},
						"targetvault": schema.Int64Attribute{
							Computed:    true,
							Description: "Target Vault",
						},
						"sourcevault": schema.Int64Attribute{
							Computed:    true,
							Description: "Source Vault",
						},
						"selection": schema.StringAttribute{
							Computed:    true,
							Description: "Selection",
						},
						"scheduletype": schema.StringAttribute{
							Computed:    true,
							Description: "Schedule Type",
						},
						"exclusion": schema.StringAttribute{
							Computed:    true,
							Description: "Exclusion",
						},
						"reptype": schema.StringAttribute{
							Computed:    true,
							Description: "Rep Type",
						},
						"retention": schema.StringAttribute{
							Computed:    true,
							Description: "Retention",
						},
						"retentionm": schema.StringAttribute{
							Computed:    true,
							Description: "Retention Type",
						},
						"encrypt": schema.StringAttribute{
							Computed:    true,
							Description: "Encrypt",
						},
						"repeatinterval": schema.StringAttribute{
							Computed:    true,
							Description: "Repeat Interval",
						},
						"exclusioninterval": schema.StringAttribute{
							Computed:    true,
							Description: "Exclusion Interval",
						},
						"remoteretention": schema.Int64Attribute{
							Computed:    true,
							Description: "Remote Retention",
						},
						"policytype": schema.StringAttribute{
							Computed:    true,
							Description: "Policy Type",
						},
						"op": schema.StringAttribute{
							Computed:    true,
							Description: "Operation",
						},
						"verification": schema.BoolAttribute{
							Computed:    true,
							Description: "Verification",
						},
						"verifychoice": schema.StringAttribute{
							Computed:    true,
							Description: "Verify Choice",
						},
						"truncatelog": schema.StringAttribute{
							Computed:    true,
							Description: "Truncate Log",
						},
						"starttime": schema.StringAttribute{
							Computed:    true,
							Description: "Start Time",
						},
						"endtime": schema.StringAttribute{
							Computed:    true,
							Description: "End Time",
						},
						// "scheduling": schema.StringAttribute{
						// 	Computed:    true,
						// 	Description: "Scheduling",
						// },
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
			ID:            types.StringValue(pol.Id),
			Name:          types.StringValue(pol.Name),
			Description:   types.StringValue(pol.Description),
			Priority:      types.StringValue(pol.Priority),
			Rpom:          types.StringValue(pol.Rpom),
			Rpo:           types.StringValue(pol.Rpo),
			Exclusiontype: types.StringValue(pol.Exclusiontype),
			Starttime:     types.StringValue(pol.Starttime),
			Endtime:       types.StringValue(pol.Endtime),
			Selection:     types.StringValue(pol.Selection),
			Scheduletype:  types.StringValue(pol.Scheduletype),
			// Scheduling:        types.StringValue(pol.Scheduling),
			Exclusion:         types.StringValue(pol.Exclusion),
			Reptype:           types.StringValue(pol.Reptype),
			Retention:         types.StringValue(pol.Retention),
			Encrypt:           types.StringValue(pol.Encrypt),
			Repeatinterval:    types.StringValue(pol.Repeatinterval),
			Exclusioninterval: types.StringValue(pol.Exclusioninterval),
			PolicyType:        types.StringValue(pol.PolicyType),
			Truncatelog:       types.StringValue(pol.Truncatelog),
			Verifychoice:      types.StringValue(pol.Verifychoice),
			Op:                types.StringValue(pol.Op),
			Href:              types.StringValue(pol.Href),

			Syncdate:        types.Int64Value(pol.Syncdate),
			Remoteretention: types.Int64Value(int64(pol.Remoteretention)),
			Targetvault:     types.Int64Value(int64(pol.Targetvault)),
			Sourcevault:     types.Int64Value(int64(pol.Sourcevault)),

			Iscontinuous: types.BoolValue(pol.Iscontinuous),
			Verification: types.BoolValue(pol.Verification),
			Stale:        types.BoolValue(pol.Stale),
		})
	}

	state = sltState

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

}
