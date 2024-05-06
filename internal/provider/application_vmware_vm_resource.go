package provider

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &applicationVmwareVMsResource{}
	_ resource.ResourceWithConfigure   = &applicationVmwareVMsResource{}
	_ resource.ResourceWithImportState = &applicationVmwareVMsResource{}
)

// NewApplicationVmwareVMsResource to create vCenter Host
func NewApplicationVmwareVMsResource() resource.Resource {
	return &applicationVmwareVMsResource{}
}

// applicationVmwareVMsResource is the resource implementation.
type applicationVmwareVMsResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type applicationVmwareVMsResourceModel struct {
	Cluster      types.String   `tfsdk:"cluster"`
	ClusterName  types.String   `tfsdk:"cluster_name"`
	VMs          []types.String `tfsdk:"vms"`
	VcenterID    types.String   `tfsdk:"vcenter_id"`
	Status       types.String   `tfsdk:"status"`
	Applications types.List     `tfsdk:"applications"`
}

// Metadata returns the resource type name.
func (r *applicationVmwareVMsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_vmware_vm"
}

// Schema defines the schema for the resource.
func (r *applicationVmwareVMsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"applications": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *applicationVmwareVMsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *applicationVmwareVMsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan applicationVmwareVMsResourceModel
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
		applicationIDs := make([]attr.Value, 0)

		for _, application := range listVMs {

			filter := backupdr.ApplicationApiListApplicationsOpts{
				Filter: optional.NewString(fmt.Sprintf("uniquename:==%s", application)),
			}

			// fetch application ID with filter
			lsApps, res, err := r.client.ApplicationApi.ListApplications(r.authCtx, &filter)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error listing applications",
					"Could not list applications, unexpected error: "+res.Status+" : "+err.Error(),
				)
				return
			}
			// capture the id of filtered application
			if lsApps.Items != nil && len(lsApps.Items) > 0 {
				applicationIDs = append(applicationIDs, types.StringValue(lsApps.Items[0].Id))
			}
		}
		if len(applicationIDs) > 0 {
			plan.Applications, diags = types.ListValue(types.StringType, applicationIDs)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
	} else {
		plan.Status = types.StringValue(respObject.Status)
		resp.Diagnostics.AddError(
			"Error adding VMs of the vCenter Host",
			"Could not add VMs of vCenter Host, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *applicationVmwareVMsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Not implemented: no method exists to read vCenter VMs
	return
}

func (r *applicationVmwareVMsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan applicationVmwareVMsResourceModel
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
		applicationIDs := make([]attr.Value, 0)
		applicationIDs = append(applicationIDs, plan.Applications.Elements()...)

		for _, application := range listVMs {
			filter := backupdr.ApplicationApiListApplicationsOpts{
				Filter: optional.NewString(fmt.Sprintf("uniquename:==%s", application)),
			}
			// fetch application ID with filter
			lsApps, res, err := r.client.ApplicationApi.ListApplications(r.authCtx, &filter)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error listing applications",
					"Could not list applications, unexpected error: "+res.Status+" : "+err.Error(),
				)
				return
			}
			// capture the id of filtered application
			if lsApps.Items != nil && len(lsApps.Items) > 0 {
				applicationIDs = append(applicationIDs, types.StringValue(lsApps.Items[0].Id))
			}
		}
		if len(applicationIDs) > 0 {
			plan.Applications, diags = types.ListValue(types.StringType, applicationIDs)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
	} else {
		plan.Status = types.StringValue(respObject.Status)
		resp.Diagnostics.AddError(
			"Error updating VMs of the vCenter Host",
			"Could not update VMs of vCenter Host, unexpected error: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *applicationVmwareVMsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Not implemented: no method exists to read vCenter VMs
	return
}

func (r *applicationVmwareVMsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
