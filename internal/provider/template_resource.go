// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/antihax/optional"
	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &templateResource{}
	_ resource.ResourceWithConfigure   = &templateResource{}
	_ resource.ResourceWithImportState = &templateResource{}
)

// NewTemplateResource to create SLA Template
func NewTemplateResource() resource.Resource {
	return &templateResource{}
}

// templateResource is the resource implementation.
type templateResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// Metadata returns the resource type name.
func (r *templateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template"
}

// Schema defines the schema for the resource.
func (r *templateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an SLA Template.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"href": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"option_href": schema.StringAttribute{
				Computed: true,
			},
			"policy_href": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			// "immutable": schema.BoolAttribute{
			// 	Optional: true,
			// },
			"managedbyagm": schema.BoolAttribute{
				Optional: true,
			},
			"usedbycloudapp": schema.BoolAttribute{
				Optional: true,
			},
			"sourcename": schema.StringAttribute{
				Optional: true,
			},
			"override": schema.StringAttribute{
				Optional: true,
			},
			"policies": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"description": schema.StringAttribute{
							Optional: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
						"href": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"starttime": schema.StringAttribute{
							Optional: true,
						},
						"endtime": schema.StringAttribute{
							Optional: true,
						},
						"priority": schema.StringAttribute{
							Optional: true,
						},
						"rpo": schema.StringAttribute{
							Optional: true,
						},
						"rpom": schema.StringAttribute{
							Optional: true,
						},
						"exclusiontype": schema.StringAttribute{
							Optional: true,
						},
						"iscontinuous": schema.BoolAttribute{
							Optional: true,
						},
						"targetvault": schema.Int64Attribute{
							Optional: true,
						},
						"sourcevault": schema.Int64Attribute{
							Optional: true,
						},
						"selection": schema.StringAttribute{
							Optional: true,
						},
						"scheduletype": schema.StringAttribute{
							Optional: true,
						},
						// "scheduling": schema.StringAttribute{
						// 	Optional: true,
						// },
						"exclusion": schema.StringAttribute{
							Optional: true,
						},
						"reptype": schema.StringAttribute{
							Optional: true,
						},
						"retention": schema.StringAttribute{
							Optional: true,
						},
						"retentionm": schema.StringAttribute{
							Optional: true,
						},
						"encrypt": schema.StringAttribute{
							Optional: true,
						},
						"repeatinterval": schema.StringAttribute{
							Optional: true,
						},
						"exclusioninterval": schema.StringAttribute{
							Optional: true,
						},
						"remoteretention": schema.Int64Attribute{
							Optional: true,
						},
						"policytype": schema.StringAttribute{
							Optional: true,
						},
						"op": schema.StringAttribute{
							Optional: true,
						},
						"verification": schema.BoolAttribute{
							Optional: true,
						},
						"verifychoice": schema.StringAttribute{
							Optional: true,
						},
						"truncatelog": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *templateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *templateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan templateResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqSlt := backupdr.SltRest{
		Name: plan.Name.ValueString(),
		// Immutable:   plan.Immutable.ValueBool(),
		Description: plan.Description.ValueString(),
		Sourcename:  plan.Sourcename.ValueString(),
		Override:    plan.Override.ValueString(),
	}

	for _, pol := range plan.Policies {
		reqSlt.Policies = append(reqSlt.Policies, backupdr.PolicyRest{
			Name:          pol.Name.ValueString(),
			Description:   pol.Description.ValueString(),
			Priority:      pol.Priority.ValueString(),
			Exclusiontype: pol.Exclusiontype.ValueString(),
			Iscontinuous:  pol.Iscontinuous.ValueBool(),
			Rpo:           pol.Rpo.ValueString(),
			Rpom:          pol.Rpom.ValueString(),
			Starttime:     pol.Starttime.ValueString(),
			Endtime:       pol.Endtime.ValueString(),
			Scheduletype:  pol.Scheduletype.ValueString(),
			// Scheduling:        pol.Scheduling.ValueString(),
			Targetvault:       int32(pol.Targetvault.ValueInt64()),
			Sourcevault:       int32(pol.Sourcevault.ValueInt64()),
			Selection:         pol.Selection.ValueString(),
			Exclusion:         pol.Exclusion.ValueString(),
			Exclusioninterval: pol.Exclusioninterval.ValueString(),
			Retention:         pol.Retention.ValueString(),
			Retentionm:        pol.Retentionm.ValueString(),
			Remoteretention:   int32(pol.Remoteretention.ValueInt64()),
			PolicyType:        pol.PolicyType.ValueString(),
			Op:                pol.Op.ValueString(),
			Verification:      pol.Verification.ValueBool(),
			Repeatinterval:    pol.Repeatinterval.ValueString(),
			Encrypt:           pol.Encrypt.ValueString(),
			Reptype:           pol.Reptype.ValueString(),
			Verifychoice:      pol.Verifychoice.ValueString(),
		})
	}

	// Generate API request body from plan
	reqBody := backupdr.SLATemplateApiCreateSltOpts{
		Body: optional.NewInterface(reqSlt),
	}

	// Create new entity
	respObject, _, err := r.client.SLATemplateApi.CreateSlt(r.authCtx, &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SLT",
			"Could not create SLA Template, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(respObject.Id)
	plan.Href = types.StringValue(respObject.Href)
	plan.OptionHref = types.StringValue(respObject.OptionHref)
	plan.PolicyHref = types.StringValue(respObject.PolicyHref)

	// response doesnot show policy details
	sltID, _ := strconv.Atoi(respObject.Id)
	respObjectPolicies, _, err := r.client.SLATemplateApi.ListPolicies(r.authCtx, int64(sltID))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SLT Policy",
			"Could not read SLT policies ID: "+respObject.Id+": "+err.Error(),
		)
		return
	}
	for i, respPol := range respObjectPolicies.Items {
		plan.Policies[i].ID = types.StringValue(respPol.Id)
		plan.Policies[i].Href = types.StringValue(respPol.Href)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *templateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state templateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed values for SLT
	respObject, _, err := r.client.SLATemplateApi.GetSlt(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SLT Template",
			"Could not read SLT Template with ID: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state.ID = types.StringValue(respObject.Id)
	state.Href = types.StringValue(respObject.Href)
	state.OptionHref = types.StringValue(respObject.OptionHref)
	state.PolicyHref = types.StringValue(respObject.PolicyHref)

	// Get refreshed values for SLT Policy
	sltID, _ := strconv.Atoi(respObject.Id)
	respObjectPolicies, _, err := r.client.SLATemplateApi.ListPolicies(r.authCtx, int64(sltID))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SLT Policy",
			"Could not read SLT policies ID: "+respObject.Id+": "+err.Error(),
		)
		return
	}
	if len(respObjectPolicies.Items) > 0 {
		for i, respPol := range respObjectPolicies.Items {
			if i < len(state.Policies) {
				state.Policies[i].ID = types.StringValue(respPol.Id)
				state.Policies[i].Href = types.StringValue(respPol.Href)
			}
		}
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *templateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan templateResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state templateResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// update SLT template
	reqSlt := backupdr.SltRest{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueString(),
		Override:    plan.Override.ValueString(),
		Href:        plan.Href.ValueString(),
		PolicyHref:  plan.PolicyHref.ValueString(),
	}

	// Generate API request body from plan
	reqBody := backupdr.SLATemplateApiUpdateSltOpts{
		Body: optional.NewInterface(reqSlt),
	}

	respObject, res, err := r.client.SLATemplateApi.UpdateSlt(r.authCtx, plan.ID.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SLA Template",
			"Could not update template, unexpected error: "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Update SLA Template",
			"An unexpected error occurred when creating the BackupDR API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"BackupDR Client Error: "+res.Status,
		)
	}

	// Update resource state with updated items and timestamp
	// Map response body to schema and populate Computed attribute values
	plan.Href = types.StringValue(respObject.Href)
	plan.OptionHref = types.StringValue(respObject.OptionHref)
	plan.PolicyHref = types.StringValue(respObject.PolicyHref)

	// create SLT template policy
	if len(plan.Policies) > len(state.Policies) {

		for i, pol := range plan.Policies {
			tflog.Info(ctx, "------------------ Update Method: Check Create policy : "+pol.ID.ValueString())
			if pol.ID.ValueString() == "" {
				tflog.Info(ctx, "------------------ Update Method: Created policy : "+pol.ID.ValueString())

				reqPol := backupdr.PolicyRest{
					Id:                pol.ID.ValueString(),
					Name:              pol.Name.ValueString(),
					Description:       pol.Description.ValueString(),
					Priority:          pol.Priority.ValueString(),
					Exclusiontype:     pol.Exclusiontype.ValueString(),
					Iscontinuous:      pol.Iscontinuous.ValueBool(),
					Rpo:               pol.Rpo.ValueString(),
					Rpom:              pol.Rpom.ValueString(),
					Starttime:         pol.Starttime.ValueString(),
					Endtime:           pol.Endtime.ValueString(),
					Targetvault:       int32(pol.Targetvault.ValueInt64()),
					Sourcevault:       int32(pol.Targetvault.ValueInt64()),
					Scheduletype:      pol.Scheduletype.ValueString(),
					Selection:         pol.Selection.ValueString(),
					Exclusion:         pol.Exclusion.ValueString(),
					Exclusioninterval: pol.Exclusioninterval.ValueString(),
					Retention:         pol.Retention.ValueString(),
					Retentionm:        pol.Retentionm.ValueString(),
					Remoteretention:   int32(pol.Remoteretention.ValueInt64()),
					PolicyType:        pol.PolicyType.ValueString(),
					Op:                pol.Op.ValueString(),
					Verification:      pol.Verification.ValueBool(),
					Repeatinterval:    pol.Repeatinterval.ValueString(),
					Encrypt:           pol.Encrypt.ValueString(),
					Reptype:           pol.Reptype.ValueString(),
					Verifychoice:      pol.Verifychoice.ValueString(),
				}
				// Generate API request body from plan
				reqPolBody := backupdr.SLATemplateApiCreatePolicyOpts{
					Body: optional.NewInterface(reqPol),
				}
				respPol, _, err := r.client.SLATemplateApi.CreatePolicy(r.authCtx, respObject.Id, &reqPolBody)
				if err != nil {
					resp.Diagnostics.AddError(
						"Error Creating SLT Policy",
						"Could not Create SLT policies ID: "+respObject.Id+": "+err.Error(),
					)
					return
				}
				plan.Policies[i].ID = types.StringValue(respPol.Id)
				plan.Policies[i].Href = types.StringValue(respPol.Href)
			}
		}

	}

	// delete SLT template policy
	if len(plan.Policies) < len(state.Policies) {
		missingPolicies := findMissingPolicies(plan.Policies, state.Policies)
		tflog.Info(ctx, "------------------ Update Method: Check Delete policy "+fmt.Sprint(len(missingPolicies)))
		for _, pol := range missingPolicies {
			tflog.Info(ctx, "------------------ Update Method: Deleted policy : "+pol.ID.ValueString())
			_, err := r.client.SLATemplateApi.DeletePolicy(r.authCtx, respObject.Id, pol.ID.ValueString())
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Creating SLT Policy",
					"Could not Create SLT policies ID: "+respObject.Id+": "+err.Error(),
				)
				return
			}
		}

	}

	// update SLT template policy
	for i, pol := range plan.Policies {
		reqPol := backupdr.PolicyRest{
			Id:                pol.ID.ValueString(),
			Name:              pol.Name.ValueString(),
			Description:       pol.Description.ValueString(),
			Priority:          pol.Priority.ValueString(),
			Exclusiontype:     pol.Exclusiontype.ValueString(),
			Iscontinuous:      pol.Iscontinuous.ValueBool(),
			Rpo:               pol.Rpo.ValueString(),
			Rpom:              pol.Rpom.ValueString(),
			Starttime:         pol.Starttime.ValueString(),
			Endtime:           pol.Endtime.ValueString(),
			Targetvault:       int32(pol.Targetvault.ValueInt64()),
			Sourcevault:       int32(pol.Targetvault.ValueInt64()),
			Scheduletype:      pol.Scheduletype.ValueString(),
			Selection:         pol.Selection.ValueString(),
			Exclusion:         pol.Exclusion.ValueString(),
			Exclusioninterval: pol.Exclusioninterval.ValueString(),
			Retention:         pol.Retention.ValueString(),
			Retentionm:        pol.Retentionm.ValueString(),
			Remoteretention:   int32(pol.Remoteretention.ValueInt64()),
			PolicyType:        pol.PolicyType.ValueString(),
			Op:                pol.Op.ValueString(),
			Verification:      pol.Verification.ValueBool(),
			Repeatinterval:    pol.Repeatinterval.ValueString(),
			Encrypt:           pol.Encrypt.ValueString(),
			Reptype:           pol.Reptype.ValueString(),
			Verifychoice:      pol.Verifychoice.ValueString(),
		}
		// Generate API request body from plan
		if pol.ID.ValueString() != "" {
			tflog.Info(ctx, "------------------ Update Method: Update policy : "+pol.ID.ValueString())
			reqPolBody := backupdr.SLATemplateApiUpdatePolicyOpts{
				Body: optional.NewInterface(reqPol),
			}

			// ignore if there is no diff

			respPol, _, err := r.client.SLATemplateApi.UpdatePolicy(r.authCtx, respObject.Id, pol.ID.ValueString(), &reqPolBody)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Updating SLT Policy",
					"Could not Update SLT policies ID: "+pol.ID.ValueString()+": "+err.Error(),
				)
				return
			}
			plan.Policies[i].Href = types.StringValue(respPol.Href)
		}
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *templateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state templateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	_, err := r.client.SLATemplateApi.DeleteSlt(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting HashiCups Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *templateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func findMissingPolicies(list1Policies, list2Policies []policyRestModel) []policyRestModel {
	var missingPolicies []policyRestModel

	// Create a map of state policies for efficient lookup
	list1PolicyMap := make(map[string]bool)
	for _, statePolicy := range list1Policies {
		list1PolicyMap[statePolicy.ID.ValueString()] = true
	}

	// Iterate through plan policies and check if they exist in the state map
	for _, list2Policy := range list2Policies {
		if _, ok := list1PolicyMap[list2Policy.ID.ValueString()]; !ok {
			missingPolicies = append(missingPolicies, list2Policy)
		}
	}

	return missingPolicies
}
