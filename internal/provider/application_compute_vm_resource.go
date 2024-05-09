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
	_ resource.Resource                = &applicationComputeVMsResource{}
	_ resource.ResourceWithConfigure   = &applicationComputeVMsResource{}
	_ resource.ResourceWithImportState = &applicationComputeVMsResource{}
)

// NewApplicationComputeVMsResource to create vCenter Host
func NewApplicationComputeVMsResource() resource.Resource {
	return &applicationComputeVMsResource{}
}

// applicationComputeVMsResource is the resource implementation.
type applicationComputeVMsResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type applicationComputeVMsResourceModel struct {
	CloudCredential    types.String   `tfsdk:"cloudcredential"`
	ApplianceClusterID types.String   `tfsdk:"appliance_clusterid"`
	VMIds              []types.String `tfsdk:"vmids"`
	Region             types.String   `tfsdk:"region"`
	ProjectID          types.String   `tfsdk:"projectid"`
	Status             types.String   `tfsdk:"status"`
	Applications       types.List     `tfsdk:"applications"`
}

// Metadata returns the resource type name.
func (r *applicationComputeVMsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_compute_vm"
}

// Schema defines the schema for the resource.
func (r *applicationComputeVMsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an GCP Cloud Virtual Machines.",
		Attributes: map[string]schema.Attribute{
			"cloudcredential": schema.StringAttribute{
				Required: true,
			},
			"appliance_clusterid": schema.StringAttribute{
				Required: true,
			},
			"projectid": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Required: true,
			},
			"vmids": schema.ListAttribute{
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
func (r *applicationComputeVMsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *applicationComputeVMsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan applicationComputeVMsResourceModel
	var listVMs []string
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, vm := range plan.VMIds {
		listVMs = append(listVMs, vm.ValueString())
	}

	reqCloudAddVMs := backupdr.CloudVmDiscoveryRest{
		Region:    plan.Region.ValueString(),
		ProjectId: plan.ProjectID.ValueString(),
		Vmids:     listVMs,
		ListOnly:  false,
	}
	reqCloudAddVMs.Cluster = &backupdr.ClusterRest{Clusterid: plan.ApplianceClusterID.ValueString()}

	// Generate API request body from plan
	reqBody := backupdr.DefaultApiAddVmOpts{
		Body: optional.NewInterface(reqCloudAddVMs),
	}

	// Add new Cloud VMs
	respObject, err := r.client.DefaultApi.AddVm(r.authCtx, plan.CloudCredential.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error adding Cloud VM",
			"Could not add Cloud VMs, unexpected error: "+err.Error(),
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
			"Error adding Cloud VM",
			"Could not add Cloud VMs, unexpected error: "+err.Error(),
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
func (r *applicationComputeVMsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Does not support Read
	return
}

func (r *applicationComputeVMsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan applicationComputeVMsResourceModel
	var listVMs []string
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, vm := range plan.VMIds {
		listVMs = append(listVMs, vm.ValueString())
	}

	reqCloudAddVMs := backupdr.CloudVmDiscoveryRest{
		Region:    plan.Region.ValueString(),
		ProjectId: plan.ProjectID.ValueString(),
		Vmids:     listVMs,
	}
	reqCloudAddVMs.Cluster = &backupdr.ClusterRest{Clusterid: plan.ApplianceClusterID.ValueString()}

	// Generate API request body from plan
	reqBody := backupdr.DefaultApiAddVmOpts{
		Body: optional.NewInterface(reqCloudAddVMs),
	}

	// Add new Cloud VMs
	respObject, err := r.client.DefaultApi.AddVm(r.authCtx, plan.CloudCredential.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Cloud VM",
			"Could not update Cloud VMs, unexpected error: "+err.Error(),
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
			"Error updating Cloud VM",
			"Could not update Cloud VMs, unexpected error: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *applicationComputeVMsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Does not support Deletion of VM.. TODO. implement Application deletion.
	return
}

func (r *applicationComputeVMsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
