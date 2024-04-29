package provider

import (
	"context"

	"github.com/antihax/optional"
	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &vcenterHostAddVMsResource{}
	_ resource.ResourceWithConfigure   = &vcenterHostAddVMsResource{}
	_ resource.ResourceWithImportState = &vcenterHostAddVMsResource{}
)

// NewVcenterHostAddVMsResource to create vCenter Host
func NewVcenterHostAddVMsResource() resource.Resource {
	return &vcenterHostAddVMsResource{}
}

// vcenterHostAddVMsResource is the resource implementation.
type vcenterHostAddVMsResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type vcenterHostAddVMsResourceModel struct {
	Cluster     types.String   `tfsdk:"cluster"`
	ClusterName types.String   `tfsdk:"cluster_name"`
	VMs         []types.String `tfsdk:"vms"`
	VcenterID   types.String   `tfsdk:"vcenter_id"`
	Status      types.String   `tfsdk:"status"`
}

// Metadata returns the resource type name.
func (r *vcenterHostAddVMsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vcenter_addvms"
}

// Schema defines the schema for the resource.
func (r *vcenterHostAddVMsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an vCenter Host to add Virtual Machines.",
		Attributes: map[string]schema.Attribute{
			"cluster": schema.StringAttribute{
				Required: true,
			},
			"vcenter_id": schema.StringAttribute{
				Required: true,
			},
			"cluster_name": schema.StringAttribute{
				Required: true,
			},
			"vms": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *vcenterHostAddVMsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *vcenterHostAddVMsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan vcenterHostAddVMsResourceModel
	var listVMs []string
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, vm := range plan.VMs {
		listVMs = append(listVMs, vm.ValueString())
	}

	reqVcenterHostAddVMs := backupdr.VmDiscoveryRest{
		Cluster: plan.Cluster.ValueString(),
		Addvms:  true,
		Vms:     listVMs,
	}

	// Generate API request body from plan
	reqBody := backupdr.HostApiVmAddNewOpts{
		Body: optional.NewInterface(reqVcenterHostAddVMs),
	}

	// Add new VMs of vCenter Host
	respObject, err := r.client.HostApi.VmAddNew(r.authCtx, plan.VcenterID.ValueString(), plan.ClusterName.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error adding VMs of the vCenter Host",
			"Could not add VMs of vCenter Host, unexpected error: "+err.Error(),
		)
		return
	}

	if respObject.StatusCode == 200 || respObject.StatusCode == 204 {
		// Map response body to schema and populate Computed attribute values
		plan.Status = types.StringValue(respObject.Status)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *vcenterHostAddVMsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Not implemented: no method exists to read vCenter VMs
	return
}

func (r *vcenterHostAddVMsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan vcenterHostAddVMsResourceModel
	var listVMs []string
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, vm := range plan.VMs {
		listVMs = append(listVMs, vm.ValueString())
	}

	reqVcenterHostAddVMs := backupdr.VmDiscoveryRest{
		Cluster: plan.Cluster.ValueString(),
		Addvms:  true,
		Vms:     listVMs,
	}

	// Generate API request body from plan
	reqBody := backupdr.HostApiVmAddNewOpts{
		Body: optional.NewInterface(reqVcenterHostAddVMs),
	}

	// Add new VMs of vCenter Host
	respObject, err := r.client.HostApi.VmAddNew(r.authCtx, plan.VcenterID.ValueString(), plan.ClusterName.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating VMs of the vCenter Host",
			"Could not update VMs of vCenter Host, unexpected error: "+err.Error(),
		)
		return
	}

	if respObject.StatusCode == 200 || respObject.StatusCode == 204 {
		// Map response body to schema and populate Computed attribute values
		plan.Status = types.StringValue(respObject.Status)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *vcenterHostAddVMsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Not implemented: no method exists to read vCenter VMs
	return
}

func (r *vcenterHostAddVMsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
