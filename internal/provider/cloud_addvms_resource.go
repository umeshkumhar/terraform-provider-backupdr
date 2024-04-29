package provider

import (
	"context"

	"github.com/antihax/optional"
	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &cloudAddVMsResource{}
	_ resource.ResourceWithConfigure   = &cloudAddVMsResource{}
	_ resource.ResourceWithImportState = &cloudAddVMsResource{}
)

// NewCloudAddVMsResource to create vCenter Host
func NewCloudAddVMsResource() resource.Resource {
	return &cloudAddVMsResource{}
}

// cloudAddVMsResource is the resource implementation.
type cloudAddVMsResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// tf go model
type cloudAddVMsResourceModel struct {
	CloudCredential types.String         `tfsdk:"cloudcredential"`
	Cluster         clusterResourceModel `tfsdk:"cluster"`
	VMIds           []types.String       `tfsdk:"vmids"`
	Region          types.String         `tfsdk:"region"`
	ProjectID       types.String         `tfsdk:"projectid"`
	Status          types.String         `tfsdk:"status"`
}

type clusterResourceModel struct {
	ClusterID types.String `tfsdk:"clusterid"`
}

// Metadata returns the resource type name.
func (r *cloudAddVMsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_addvms"
}

// Schema defines the schema for the resource.
func (r *cloudAddVMsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an GCP Cloud Virtual Machines.",
		Attributes: map[string]schema.Attribute{
			"cloudcredential": schema.StringAttribute{
				Required: true,
			},
			"cluster": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"clusterid": schema.StringAttribute{
						Required: true,
					},
				},
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
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *cloudAddVMsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *cloudAddVMsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan cloudAddVMsResourceModel
	var listVMs []string
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, vm := range plan.VMIds {
		listVMs = append(listVMs, vm.ValueString())
	}

	tflog.Info(ctx, "================== "+plan.ProjectID.ValueString())
	reqCloudAddVMs := backupdr.CloudVmDiscoveryRest{
		Region:    plan.Region.ValueString(),
		ProjectId: plan.ProjectID.ValueString(),
		Vmids:     listVMs,
		ListOnly:  false,
	}
	reqCloudAddVMs.Cluster = &backupdr.ClusterRest{Clusterid: plan.Cluster.ClusterID.ValueString()}

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
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *cloudAddVMsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Does not support Read
	return
}

func (r *cloudAddVMsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan cloudAddVMsResourceModel
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
	reqCloudAddVMs.Cluster = &backupdr.ClusterRest{Clusterid: plan.Cluster.ClusterID.ValueString()}

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
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *cloudAddVMsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Does not support Deletion of VM.. TODO. implement Application deletion.
	return
}

func (r *cloudAddVMsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
