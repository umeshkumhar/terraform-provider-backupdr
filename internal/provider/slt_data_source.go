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
	_ datasource.DataSource              = &sltDataSource{}
	_ datasource.DataSourceWithConfigure = &sltDataSource{}
)

// sltDataSource is the data source implementation.
type sltDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// NewSltDataSource - Datasource for SLA Template
func NewSltDataSource() datasource.DataSource {
	return &sltDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *sltDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *sltDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_slt"
}

func (d *sltDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
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
			"immutable": schema.BoolAttribute{
				Computed: true,
			},
			"stale": schema.BoolAttribute{
				Computed: true,
			},
			"syncdate": schema.Int64Attribute{
				Computed: true,
			},
			"managedbyagm": schema.BoolAttribute{
				Computed: true,
			},
			"usedbycloudapp": schema.BoolAttribute{
				Computed: true,
			},
			"option_href": schema.StringAttribute{
				Computed: true,
			},
			"policy_href": schema.StringAttribute{
				Computed: true,
			},
			"sourcename": schema.StringAttribute{
				Computed: true,
			},
			"override": schema.StringAttribute{
				Computed: true,
			},
			"policies": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"href": schema.StringAttribute{
							Computed: true,
						},
						"stale": schema.BoolAttribute{
							Computed: true,
						},
						"syncdate": schema.Int64Attribute{
							Computed: true,
						},
						"priority": schema.StringAttribute{
							Computed: true,
						},
						"rpo": schema.StringAttribute{
							Computed: true,
						},
						"rpom": schema.StringAttribute{
							Computed: true,
						},
						"exclusiontype": schema.StringAttribute{
							Computed: true,
						},
						"iscontinuous": schema.BoolAttribute{
							Computed: true,
						},
						"targetvault": schema.Int64Attribute{
							Computed: true,
						},
						"sourcevault": schema.Int64Attribute{
							Computed: true,
						},
						"selection": schema.StringAttribute{
							Computed: true,
						},
						"scheduletype": schema.StringAttribute{
							Computed: true,
						},
						"exclusion": schema.StringAttribute{
							Computed: true,
						},
						"reptype": schema.StringAttribute{
							Computed: true,
						},
						"retention": schema.StringAttribute{
							Computed: true,
						},
						"retentionm": schema.StringAttribute{
							Computed: true,
						},
						"encrypt": schema.StringAttribute{
							Computed: true,
						},
						"repeatinterval": schema.StringAttribute{
							Computed: true,
						},
						"exclusioninterval": schema.StringAttribute{
							Computed: true,
						},
						"remoteretention": schema.Int64Attribute{
							Computed: true,
						},
						"policytype": schema.StringAttribute{
							Computed: true,
						},
						"op": schema.StringAttribute{
							Computed: true,
						},
						"verification": schema.BoolAttribute{
							Computed: true,
						},
						"verifychoice": schema.StringAttribute{
							Computed: true,
						},
						"truncatelog": schema.StringAttribute{
							Computed: true,
						},
						"starttime": schema.StringAttribute{
							Computed: true,
						},
						"endtime": schema.StringAttribute{
							Computed: true,
						},
						"scheduling": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *sltDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state sltRestModel
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
	sltState := sltRestModel{
		ID:          types.StringValue(slt.Id),
		Href:        types.StringValue(slt.Href),
		Name:        types.StringValue(slt.Name),
		Description: types.StringValue(slt.Description),
		Immutable:   types.BoolValue(slt.Immutable),
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
			Scheduling:        types.StringValue(pol.Scheduling),
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
