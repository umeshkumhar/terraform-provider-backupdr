package provider

import (
	"context"

	"github.com/antihax/optional"
	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &hostResource{}
	_ resource.ResourceWithConfigure   = &hostResource{}
	_ resource.ResourceWithImportState = &hostResource{}
)

// NewHostResource to create vCenter Host
func NewHostResource() resource.Resource {
	return &hostResource{}
}

// hostResource is the resource implementation.
type hostResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// Metadata returns the resource type name.
func (r *hostResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_host"
}

// Schema defines the schema for the resource.
func (r *hostResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an vCenter Host.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"href": schema.StringAttribute{
				Computed: true,
			},
			"friendlypath": schema.StringAttribute{
				Optional: true,
			},
			"hostname": schema.StringAttribute{
				Required: true,
			},
			"hosttype": schema.StringAttribute{
				Required: true,
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
			},
			"autoupgrade": schema.StringAttribute{
				Computed: true,
			},
			"cert_revoked": schema.BoolAttribute{
				Computed: true,
			},
			"clusterid": schema.StringAttribute{
				Computed: true,
			},
			"dbauthentication": schema.BoolAttribute{
				Computed: true,
			},
			"diskpref": schema.StringAttribute{
				Computed: true,
			},
			"hasagent": schema.BoolAttribute{
				Computed: true,
			},
			"isclusternode": schema.BoolAttribute{
				Computed: true,
			},

			"isshadowhost": schema.BoolAttribute{
				Computed: true,
			},
			"isclusterhost": schema.BoolAttribute{
				Computed: true,
			},
			"isesxhost": schema.BoolAttribute{
				Computed: true,
			},
			"isproxyhost": schema.BoolAttribute{
				Computed: true,
			},
			"isvcenterhost": schema.BoolAttribute{
				Computed: true,
			},
			"isvm": schema.BoolAttribute{
				Computed: true,
			},

			"maxjobs": schema.Int64Attribute{
				Computed: true,
			},
			"modifydate": schema.Int64Attribute{
				Computed: true,
			},
			"multiregion": schema.StringAttribute{
				Optional: true,
			},

			"name": schema.StringAttribute{
				Computed: true,
			},
			"originalhostid": schema.StringAttribute{
				Computed: true,
			},
			"ostype_special": schema.StringAttribute{
				Optional: true,
			},
			"pki_state": schema.StringAttribute{
				Computed: true,
			},
			"sourcecluster": schema.StringAttribute{
				Computed: true,
			},
			"srcid": schema.StringAttribute{
				Computed: true,
			},
			"svcname": schema.StringAttribute{
				Computed: true,
			},
			"transport": schema.StringAttribute{
				Computed: true,
			},
			"uniquename": schema.StringAttribute{
				Computed: true,
			},
			"zone": schema.StringAttribute{
				Computed: true,
			},
			"alternateip": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},

			"hypervisoragent": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"username": schema.StringAttribute{
						Required: true,
					},
					"password": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
					},
					"haspassword": schema.BoolAttribute{
						Computed: true,
					},
					"hasalternatekey": schema.BoolAttribute{
						Computed: true,
					},
					"agenttype": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"appliance_clusterid": schema.StringAttribute{
				Required: true,
			},
			// "sources": schema.ListNestedAttribute{
			// 	Optional: true,
			// 	NestedObject: schema.NestedAttributeObject{
			// 		Attributes: map[string]schema.Attribute{
			// 			"id": schema.StringAttribute{
			// 				Computed: true,
			// 			},
			// 			"href": schema.StringAttribute{
			// 				Computed: true,
			// 			},
			// 			"clusterid": schema.StringAttribute{
			// 				Required: true,
			// 			},
			// 		},
			// 	},
			// },
			// "agents": schema.ListNestedAttribute{
			// 	Optional: true,
			// 	NestedObject: schema.NestedAttributeObject{
			// 		Attributes: map[string]schema.Attribute{
			// 			"agenttype": schema.StringAttribute{
			// 				Computed: true,
			// 			},
			// 			"hasalternatekey": schema.BoolAttribute{
			// 				Computed: true,
			// 			},
			// 			"haspassword": schema.BoolAttribute{
			// 				Computed: true,
			// 			},
			// 			"password": schema.StringAttribute{
			// 				Computed:  true,
			// 				Optional:  true,
			// 				Sensitive: true,
			// 			},
			// 			"username": schema.StringAttribute{
			// 				Computed: true,
			// 			},
			// 		},
			// 	},
			// },
			// "esxlist": schema.ListNestedAttribute{
			// 	Computed: true,
			// 	NestedObject: schema.NestedAttributeObject{
			// 		Attributes: map[string]schema.Attribute{
			// 			"id": schema.StringAttribute{
			// 				Computed: true,
			// 			},
			// 			"href": schema.StringAttribute{
			// 				Computed: true,
			// 			},
			// 			"clusterid": schema.StringAttribute{
			// 				Computed: true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *hostResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *hostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan vcenterHostRest
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqVcenterHost := backupdr.HostRest{
		Hostname:     plan.Hostname.ValueString(),
		Hosttype:     plan.Hosttype.ValueString(),
		Friendlypath: plan.Friendlypath.ValueString(),
		Ipaddress:    plan.Ipaddress.ValueString(),
	}

	if plan.Hypervisoragent != nil {
		reqVcenterHost.Hypervisoragent = &backupdr.AgentRest{
			Username: plan.Hypervisoragent.Username.ValueString(),
			Password: plan.Hypervisoragent.Password.ValueString(),
		}
	}

	reqVcenterHost.Sources = append(reqVcenterHost.Sources, backupdr.HostRest{
		Clusterid: plan.ApplianceClusterID.ValueString(),
	})

	// Generate API request body from plan
	reqBody := backupdr.HostApiCreateHostOpts{
		Body: optional.NewInterface(reqVcenterHost),
	}

	// Create new vCenter Host
	respObject, _, err := r.client.HostApi.CreateHost(r.authCtx, &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating vCenter Host",
			"Could not create vCenter Host, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(respObject.Id)
	plan.Href = types.StringValue(respObject.Href)
	// plan.Agents = types.BoolValue(respObject.Stale)
	plan.Modifydate = types.Int64Value(respObject.Modifydate)
	plan.Autoupgrade = types.StringValue(respObject.Autoupgrade)
	plan.CertRevoked = types.BoolValue(respObject.CertRevoked)
	plan.Clusterid = types.StringValue(respObject.Clusterid)
	plan.Dbauthentication = types.BoolValue(respObject.Dbauthentication)
	plan.Diskpref = types.StringValue(respObject.Diskpref)
	plan.Friendlypath = types.StringValue(respObject.Friendlypath)
	plan.Hasagent = types.BoolValue(respObject.Hasagent)
	plan.Hostname = types.StringValue(respObject.Hostname)
	plan.Hosttype = types.StringValue(respObject.Hosttype)
	plan.Ipaddress = types.StringValue(respObject.Ipaddress)
	plan.IsClusterNode = types.BoolValue(respObject.IsClusterNode)
	plan.Isclusterhost = types.BoolValue(respObject.Isclusterhost)
	plan.IsShadowHost = types.BoolValue(respObject.IsShadowHost)
	plan.Isesxhost = types.BoolValue(respObject.Isesxhost)
	plan.Isproxyhost = types.BoolValue(respObject.Isproxyhost)
	plan.Isvcenterhost = types.BoolValue(respObject.Isvcenterhost)
	plan.Isvm = types.BoolValue(respObject.Isvm)
	plan.Maxjobs = types.Int64Value(int64(respObject.Maxjobs))
	plan.Name = types.StringValue(respObject.Name)
	plan.Originalhostid = types.StringValue(respObject.Originalhostid)
	plan.PkiState = types.StringValue(respObject.PkiState)
	plan.Sourcecluster = types.StringValue(respObject.Sourcecluster)
	plan.Srcid = types.StringValue(respObject.Srcid)
	plan.Svcname = types.StringValue(respObject.Svcname)
	plan.Transport = types.StringValue(respObject.Transport)
	plan.Uniquename = types.StringValue(respObject.Uniquename)
	plan.Zone = types.StringValue(respObject.Zone)

	plan.Hypervisoragent.Haspassword = types.BoolValue(respObject.Hypervisoragent.Haspassword)
	plan.Hypervisoragent.Agenttype = types.StringValue(respObject.Hypervisoragent.Agenttype)
	plan.Hypervisoragent.Hasalternatekey = types.BoolValue(respObject.Hypervisoragent.Hasalternatekey)

	// plan.Agents = []types.Object{{
	// 	Agenttype:       types.StringValue(respObject.Agents[0].Agenttype),
	// 	Hasalternatekey: types.BoolValue(respObject.Agents[0].Hasalternatekey),
	// 	Haspassword:     types.BoolValue(respObject.Agents[0].Haspassword),
	// 	Username:        types.StringValue(respObject.Agents[0].Username),
	// 	Password:        types.StringNull(),
	// }}

	// plan.Sources[0].ID = types.StringValue(respObject.Sources[0].Id)
	// plan.Sources[0].Href = types.StringValue(respObject.Sources[0].Href)

	// plan.Appliance = &ClusterRestRef{
	// 	// Clusterid: types.StringValue(respObject.Appliance.Clusterid),
	// 	ID:   types.StringValue(respObject.Appliance.Id),
	// 	Href: types.StringValue(respObject.Appliance.Href),
	// }

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *hostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state vcenterHostRest
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed values
	respObject, _, err := r.client.HostApi.GetHost(r.authCtx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading vCenter Host",
			"Could not read vCenter Host with ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	// Map response body to schema and populate Computed attribute values
	state.ID = types.StringValue(respObject.Id)
	state.Href = types.StringValue(respObject.Href)
	// plan.Agents = types.BoolValue(respObject.Stale)
	state.Modifydate = types.Int64Value(respObject.Modifydate)
	state.Autoupgrade = types.StringValue(respObject.Autoupgrade)
	state.CertRevoked = types.BoolValue(respObject.CertRevoked)
	state.Clusterid = types.StringValue(respObject.Clusterid)
	state.Dbauthentication = types.BoolValue(respObject.Dbauthentication)
	state.Diskpref = types.StringValue(respObject.Diskpref)
	state.Friendlypath = types.StringValue(respObject.Friendlypath)
	state.Hasagent = types.BoolValue(respObject.Hasagent)
	state.Hostname = types.StringValue(respObject.Hostname)
	state.Hosttype = types.StringValue(respObject.Hosttype)
	state.Ipaddress = types.StringValue(respObject.Ipaddress)
	state.IsClusterNode = types.BoolValue(respObject.IsClusterNode)
	state.Isclusterhost = types.BoolValue(respObject.Isclusterhost)
	state.IsShadowHost = types.BoolValue(respObject.IsShadowHost)
	state.Isesxhost = types.BoolValue(respObject.Isesxhost)
	state.Isproxyhost = types.BoolValue(respObject.Isproxyhost)
	state.Isvcenterhost = types.BoolValue(respObject.Isvcenterhost)
	state.Isvm = types.BoolValue(respObject.Isvm)
	state.Maxjobs = types.Int64Value(int64(respObject.Maxjobs))
	state.Name = types.StringValue(respObject.Name)
	state.Originalhostid = types.StringValue(respObject.Originalhostid)
	state.PkiState = types.StringValue(respObject.PkiState)
	state.Sourcecluster = types.StringValue(respObject.Sourcecluster)
	state.Srcid = types.StringValue(respObject.Srcid)
	state.Svcname = types.StringValue(respObject.Svcname)
	state.Transport = types.StringValue(respObject.Transport)
	state.Uniquename = types.StringValue(respObject.Uniquename)
	state.Zone = types.StringValue(respObject.Zone)

	state.Hypervisoragent.Haspassword = types.BoolValue(respObject.Hypervisoragent.Haspassword)
	state.Hypervisoragent.Agenttype = types.StringValue(respObject.Hypervisoragent.Agenttype)
	state.Hypervisoragent.Hasalternatekey = types.BoolValue(respObject.Hypervisoragent.Hasalternatekey)

	// state.Agents = []AgentRest{{
	// 	Agenttype:       types.StringValue(respObject.Agents[0].Agenttype),
	// 	Hasalternatekey: types.BoolValue(respObject.Agents[0].Hasalternatekey),
	// 	Haspassword:     types.BoolValue(respObject.Agents[0].Haspassword),
	// 	Username:        types.StringValue(respObject.Agents[0].Username),
	// 	Password:        types.StringNull(),
	// }}

	// state.Sources[0].ID = types.StringValue(respObject.Sources[0].Id)
	// state.Sources[0].Href = types.StringValue(respObject.Sources[0].Href)

	// state.Appliance = &ClusterRestRef{
	// 	// Clusterid: types.StringValue(respObject.Appliance.Clusterid),
	// 	ID:   types.StringValue(respObject.Appliance.Id),
	// 	Href: types.StringValue(respObject.Appliance.Href),
	// }

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *hostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan vcenterHostRest
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqVcenterHost := backupdr.HostRest{
		Hostname:     plan.Hostname.ValueString(),
		Hosttype:     plan.Hosttype.ValueString(),
		Friendlypath: plan.Friendlypath.ValueString(),
		Ipaddress:    plan.Ipaddress.ValueString(),
	}

	if plan.Hypervisoragent != nil {
		reqVcenterHost.Hypervisoragent = &backupdr.AgentRest{
			Username: plan.Hypervisoragent.Username.ValueString(),
			Password: plan.Hypervisoragent.Password.ValueString(),
		}
	}

	reqVcenterHost.Sources = append(reqVcenterHost.Sources, backupdr.HostRest{
		Clusterid: plan.ApplianceClusterID.ValueString(),
	})

	// Generate API request body from plan
	reqBody := backupdr.HostApiUpdateHostOpts{
		Body: optional.NewInterface(reqVcenterHost),
	}

	// Update vCenter Host
	respObject, res, err := r.client.HostApi.UpdateHost(r.authCtx, plan.ID.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating vCenter Host",
			"Could not updating vCenter Host, unexpected error: "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Update vCenter Host ",
			"An unexpected error occurred when creating the BackupDR API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"BackupDR Client Error: "+res.Status,
		)
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(respObject.Id)
	plan.Href = types.StringValue(respObject.Href)
	// plan.Agents = types.BoolValue(respObject.Stale)
	plan.Modifydate = types.Int64Value(respObject.Modifydate)
	plan.Autoupgrade = types.StringValue(respObject.Autoupgrade)
	plan.CertRevoked = types.BoolValue(respObject.CertRevoked)
	plan.Clusterid = types.StringValue(respObject.Clusterid)
	plan.Dbauthentication = types.BoolValue(respObject.Dbauthentication)
	plan.Diskpref = types.StringValue(respObject.Diskpref)
	plan.Friendlypath = types.StringValue(respObject.Friendlypath)
	plan.Hasagent = types.BoolValue(respObject.Hasagent)
	plan.Hostname = types.StringValue(respObject.Hostname)
	plan.Hosttype = types.StringValue(respObject.Hosttype)
	plan.Ipaddress = types.StringValue(respObject.Ipaddress)
	plan.IsClusterNode = types.BoolValue(respObject.IsClusterNode)
	plan.Isclusterhost = types.BoolValue(respObject.Isclusterhost)
	plan.IsShadowHost = types.BoolValue(respObject.IsShadowHost)
	plan.Isesxhost = types.BoolValue(respObject.Isesxhost)
	plan.Isproxyhost = types.BoolValue(respObject.Isproxyhost)
	plan.Isvcenterhost = types.BoolValue(respObject.Isvcenterhost)
	plan.Isvm = types.BoolValue(respObject.Isvm)
	plan.Maxjobs = types.Int64Value(int64(respObject.Maxjobs))
	plan.Name = types.StringValue(respObject.Name)
	plan.Originalhostid = types.StringValue(respObject.Originalhostid)
	plan.PkiState = types.StringValue(respObject.PkiState)
	plan.Sourcecluster = types.StringValue(respObject.Sourcecluster)
	plan.Srcid = types.StringValue(respObject.Srcid)
	plan.Svcname = types.StringValue(respObject.Svcname)
	plan.Transport = types.StringValue(respObject.Transport)
	plan.Uniquename = types.StringValue(respObject.Uniquename)
	plan.Zone = types.StringValue(respObject.Zone)

	plan.Hypervisoragent.Haspassword = types.BoolValue(respObject.Hypervisoragent.Haspassword)
	plan.Hypervisoragent.Agenttype = types.StringValue(respObject.Hypervisoragent.Agenttype)
	plan.Hypervisoragent.Hasalternatekey = types.BoolValue(respObject.Hypervisoragent.Hasalternatekey)

	// plan.Agents = []AgentRest{{
	// 	Agenttype:       types.StringValue(respObject.Agents[0].Agenttype),
	// 	Hasalternatekey: types.BoolValue(respObject.Agents[0].Hasalternatekey),
	// 	Haspassword:     types.BoolValue(respObject.Agents[0].Haspassword),
	// 	Username:        types.StringValue(respObject.Agents[0].Username),
	// 	Password:        types.StringNull(),
	// }}

	// plan.Sources[0].ID = types.StringValue(respObject.Sources[0].Id)
	// plan.Sources[0].Href = types.StringValue(respObject.Sources[0].Href)

	// plan.Appliance, _ = types.ObjectValue(ClusterRestRef{
	// 	// Clusterid: types.StringValue(respObject.Appliance.Clusterid),
	// 	ID:   types.StringValue(respObject.Appliance.Id),
	// 	Href: types.StringValue(respObject.Appliance.Href),
	// })

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *hostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state vcenterHostRest
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing vCenter Host
	_, err := r.client.HostApi.DeleteHost(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting vCenter Host",
			"Could not delete vCenter Host, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *hostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
