// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/antihax/optional"
	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &diskpoolResource{}
	_ resource.ResourceWithConfigure   = &diskpoolResource{}
	_ resource.ResourceWithImportState = &diskpoolResource{}
)

// NewDiskpoolResource to create DiskPool
func NewDiskpoolResource() resource.Resource {
	return &diskpoolResource{}
}

// diskpoolResource is the resource implementation.
type diskpoolResource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// Metadata returns the resource type name.
func (r *diskpoolResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_diskpool"
}

// Schema defines the schema for the resource.
func (r *diskpoolResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an DiskPool.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"href": schema.StringAttribute{
				Computed: true,
			},
			"stale": schema.BoolAttribute{
				Computed: true,
			},
			"modifydate": schema.Int64Attribute{
				Computed: true,
			},
			"syncdate": schema.Int64Attribute{
				Computed: true,
			},
			"pooltype": schema.StringAttribute{
				Required: true,
			},
			"state": schema.StringAttribute{
				Computed: true,
			},
			"srcid": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"mdiskgrp": schema.StringAttribute{
				Computed: true,
			},
			"pooltypedisplayname": schema.StringAttribute{
				Computed: true,
			},

			"warnpct": schema.Int64Attribute{
				Computed: true,
			},
			"safepct": schema.Int64Attribute{
				Computed: true,
			},
			"udsuid": schema.Int64Attribute{
				Computed: true,
			},
			"free_mb": schema.Int64Attribute{
				Computed: true,
			},
			"usage_mb": schema.Int64Attribute{
				Computed: true,
			},
			"capacity_mb": schema.Int64Attribute{
				Computed: true,
			},
			"pct": schema.Float64Attribute{
				Computed: true,
			},

			"usedefaultsa": schema.BoolAttribute{
				Computed: true,
			},
			"immutable": schema.BoolAttribute{
				Computed: true,
			},
			"metadataonly": schema.BoolAttribute{
				Computed: true,
			},

			"properties": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"appliance_clusterid": schema.StringAttribute{
				Required: true,
			},
			"cluster": schema.SingleNestedAttribute{
				Computed: true,
				Default: objectdefault.StaticValue(
					types.ObjectValueMust(
						map[string]attr.Type{
							"id": types.StringType,
						},
						map[string]attr.Value{
							"id": types.StringValue(""),
						},
					),
				),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
					"clusterid": schema.StringAttribute{
						Computed: true,
					},
					"serviceaccount": schema.StringAttribute{
						Computed: true,
					},
					"zone": schema.StringAttribute{
						Computed: true,
					},
					"region": schema.StringAttribute{
						Computed: true,
					},
					"projectid": schema.StringAttribute{
						Computed: true,
					},
					"version": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Computed: true,
					},
					"ipaddress": schema.StringAttribute{
						Computed: true,
					},
					"publicip": schema.StringAttribute{
						Computed: true,
					},
					"supportstatus": schema.StringAttribute{
						Computed: true,
					},
					"secureconnect": schema.BoolAttribute{
						Computed: true,
					},
					"pkibootstrapped": schema.BoolAttribute{
						Computed: true,
					},
					"stale": schema.BoolAttribute{
						Computed: true,
					},
					"syncdate": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"vaultprops": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"bucket": schema.StringAttribute{
						Computed: true,
					},
					"href": schema.StringAttribute{
						Computed: true,
					},
					"region": schema.StringAttribute{
						Computed: true,
					},

					"compression": schema.BoolAttribute{
						Computed: true,
					},
					"stale": schema.BoolAttribute{
						Computed: true,
					},
					"syncdate": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *diskpoolResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*backupdrProvider).client
	r.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

// Create a new resource.
func (r *diskpoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from state
	var state diskPoolResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqDiskpool := backupdr.DiskPoolRest{
		Name:     state.Name.ValueString(),
		Pooltype: state.Pooltype.ValueString(),
		Cluster:  &backupdr.ClusterRest{Clusterid: state.ApplianceClusterID.ValueString()},
	}

	for _, prop := range state.Properties {
		reqDiskpool.Properties = append(reqDiskpool.Properties, backupdr.KeyValueRest{
			Key:   prop.Key.ValueString(),
			Value: prop.Value.ValueString(),
		})
	}

	// Generate API request body from state
	reqBody := backupdr.DiskPoolApiCreateDiskPoolOpts{
		Body: optional.NewInterface(reqDiskpool),
	}

	// Create new diskpool
	respObject, res, err := r.client.DiskPoolApi.CreateDiskPool(r.authCtx, &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Diskpool",
			"Could not create Diskpool, unexpected error: "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Create BackupDR DiskPool",
			res.Status,
		)
	}

	// Map response body to schema and populate Computed attribute values
	state.ID = types.StringValue(respObject.Id)
	state.Href = types.StringValue(respObject.Href)
	state.Syncdate = types.Int64Value(respObject.Syncdate)
	state.Stale = types.BoolValue(respObject.Stale)
	state.Usedefaultsa = types.BoolValue(respObject.Usedefaultsa)
	state.Immutable = types.BoolValue(respObject.Immutable)
	state.Metadataonly = types.BoolValue(respObject.Metadataonly)
	state.State = types.StringValue(respObject.State)
	state.Srcid = types.StringValue(respObject.Srcid)
	state.Status = types.StringValue(respObject.Status)
	state.Mdiskgrp = types.StringValue(respObject.Mdiskgrp)
	state.Pooltypedisplayname = types.StringValue(respObject.Pooltypedisplayname)
	state.Srcid = types.StringValue(respObject.Srcid)
	state.Warnpct = types.Int64Value(int64(respObject.Warnpct))
	state.Modifydate = types.Int64Value(respObject.Modifydate)
	state.Safepct = types.Int64Value(int64(respObject.Safepct))
	state.Udsuid = types.Int64Value(int64(respObject.Udsuid))
	state.FreeMb = types.Int64Value(respObject.FreeMb)
	state.UsageMb = types.Int64Value(respObject.UsageMb)
	state.CapacityMb = types.Int64Value(respObject.CapacityMb)
	state.Pct = types.Float64Value(respObject.Pct)

	state.ApplianceClusterID = types.StringValue(respObject.Cluster.Clusterid)
	// Set state to fully populated data
	// state.Cluster = &ClusterRest{
	// 	ID:              types.StringValue(respObject.Cluster.Id),
	// 	Href:            types.StringValue(respObject.Cluster.Href),
	// 	Serviceaccount:  types.StringValue(respObject.Cluster.Serviceaccount),
	// 	Zone:            types.StringValue(respObject.Cluster.Zone),
	// 	Region:          types.StringValue(respObject.Cluster.Region),
	// 	Projectid:       types.StringValue(respObject.Cluster.Projectid),
	// 	Version:         types.StringValue(respObject.Cluster.Version),
	// 	Name:            types.StringValue(respObject.Cluster.Name),
	// 	Type:            types.StringValue(respObject.Cluster.Type_),
	// 	Ipaddress:       types.StringValue(respObject.Cluster.Ipaddress),
	// 	Publicip:        types.StringValue(respObject.Cluster.Publicip),
	// 	Secureconnect:   types.BoolValue(respObject.Cluster.Secureconnect),
	// 	PkiBootstrapped: types.BoolValue(respObject.Cluster.PkiBootstrapped),
	// 	Supportstatus:   types.StringValue(respObject.Cluster.Supportstatus),
	// 	Syncdate:        types.Int64Value(respObject.Cluster.Syncdate),
	// 	Stale:           types.BoolValue(respObject.Cluster.Stale),
	// }

	// // state.Cluster.Href = types.StringValue(respObject.Cluster.Href)
	// // state.Cluster.ID = types.StringValue(respObject.Cluster.Id)
	// // state.Cluster.Serviceaccount = types.StringValue(respObject.Cluster.Serviceaccount)
	// // state.Cluster.Zone = types.StringValue(respObject.Cluster.Zone)
	// // state.Cluster.Region = types.StringValue(respObject.Cluster.Region)
	// // state.Cluster.Projectid = types.StringValue(respObject.Cluster.Projectid)
	// // state.Cluster.Version = types.StringValue(respObject.Cluster.Version)
	// // state.Cluster.Name = types.StringValue(respObject.Cluster.Name)
	// // state.Cluster.Type = types.StringValue(respObject.Cluster.Type_)
	// // state.Cluster.Ipaddress = types.StringValue(respObject.Cluster.Ipaddress)
	// // state.Cluster.Publicip = types.StringValue(respObject.Cluster.Publicip)
	// // state.Cluster.Secureconnect = types.BoolValue(respObject.Cluster.Secureconnect)
	// // state.Cluster.PkiBootstrapped = types.BoolValue(respObject.Cluster.PkiBootstrapped)
	// // state.Cluster.Supportstatus = types.StringValue(respObject.Cluster.Supportstatus)
	// // state.Cluster.Syncdate = types.Int64Value(respObject.Cluster.Syncdate)
	// // state.Cluster.Stale = types.BoolValue(respObject.Cluster.Stale)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *diskpoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state diskPoolResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed values
	respObject, res, err := r.client.DiskPoolApi.GetDiskPool(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading DiskPool",
			"Could not read DiskPool with ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR DiskPool",
			res.Status,
		)
	}

	// Overwrite items with refreshed state
	state.ID = types.StringValue(respObject.Id)
	state.Href = types.StringValue(respObject.Href)
	state.Syncdate = types.Int64Value(respObject.Syncdate)
	state.Stale = types.BoolValue(respObject.Stale)
	state.Usedefaultsa = types.BoolValue(respObject.Usedefaultsa)
	state.Immutable = types.BoolValue(respObject.Immutable)
	state.Metadataonly = types.BoolValue(respObject.Metadataonly)
	state.State = types.StringValue(respObject.State)
	state.Srcid = types.StringValue(respObject.Srcid)
	state.Status = types.StringValue(respObject.Status)
	state.Mdiskgrp = types.StringValue(respObject.Mdiskgrp)
	state.Pooltypedisplayname = types.StringValue(respObject.Pooltypedisplayname)
	state.Srcid = types.StringValue(respObject.Srcid)
	state.Warnpct = types.Int64Value(int64(respObject.Warnpct))
	state.Modifydate = types.Int64Value(respObject.Modifydate)
	state.Safepct = types.Int64Value(int64(respObject.Safepct))
	state.Udsuid = types.Int64Value(int64(respObject.Udsuid))
	state.FreeMb = types.Int64Value(respObject.FreeMb)
	state.UsageMb = types.Int64Value(respObject.UsageMb)
	state.CapacityMb = types.Int64Value(respObject.CapacityMb)
	state.Pct = types.Float64Value(respObject.Pct)

	state.ApplianceClusterID = types.StringValue(respObject.Cluster.Clusterid)

	// Set state to fully populated data
	// state.Cluster = &ClusterRest{
	// 	ID:              types.StringValue(respObject.Cluster.Id),
	// 	Href:            types.StringValue(respObject.Cluster.Href),
	// 	Serviceaccount:  types.StringValue(respObject.Cluster.Serviceaccount),
	// 	Zone:            types.StringValue(respObject.Cluster.Zone),
	// 	Region:          types.StringValue(respObject.Cluster.Region),
	// 	Projectid:       types.StringValue(respObject.Cluster.Projectid),
	// 	Version:         types.StringValue(respObject.Cluster.Version),
	// 	Name:            types.StringValue(respObject.Cluster.Name),
	// 	Type:            types.StringValue(respObject.Cluster.Type_),
	// 	Ipaddress:       types.StringValue(respObject.Cluster.Ipaddress),
	// 	Publicip:        types.StringValue(respObject.Cluster.Publicip),
	// 	Secureconnect:   types.BoolValue(respObject.Cluster.Secureconnect),
	// 	PkiBootstrapped: types.BoolValue(respObject.Cluster.PkiBootstrapped),
	// 	Supportstatus:   types.StringValue(respObject.Cluster.Supportstatus),
	// 	Syncdate:        types.Int64Value(respObject.Cluster.Syncdate),
	// 	Stale:           types.BoolValue(respObject.Cluster.Stale),
	// }

	// state.Vaultprops = &VaultPropsRest{
	// 	Bucket:      types.StringValue(respObject.Vaultprops.Bucket),
	// 	Compression: types.BoolValue(respObject.Vaultprops.Compression),
	// 	Region:      types.StringValue(respObject.Vaultprops.Region),
	// 	ID:          types.StringValue(respObject.Vaultprops.Id),
	// 	Href:        types.StringValue(respObject.Vaultprops.Href),
	// 	Syncdate:    types.Int64Value(respObject.Vaultprops.Syncdate),
	// 	Stale:       types.BoolValue(respObject.Vaultprops.Stale),
	// }

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *diskpoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from state
	var state diskPoolResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqDiskpool := backupdr.DiskPoolRest{
		Name:     state.Name.ValueString(),
		Pooltype: state.Pooltype.ValueString(),
		Cluster:  &backupdr.ClusterRest{Clusterid: state.ApplianceClusterID.ValueString()},
	}

	for _, prop := range state.Properties {
		reqDiskpool.Properties = append(reqDiskpool.Properties, backupdr.KeyValueRest{
			Key:   prop.Key.ValueString(),
			Value: prop.Value.ValueString(),
		})
	}

	// Generate API request body from state
	reqBody := backupdr.DiskPoolApiUpdateDiskPoolOpts{
		Body: optional.NewInterface(reqDiskpool),
	}

	// Update existing order
	respObject, res, err := r.client.DiskPoolApi.UpdateDiskPool(r.authCtx, state.ID.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating DiskPool",
			"An unexpected error occurred when updating the BackupDR DiskPool, unexpected error: "+err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Update Diskpool",
			"An unexpected error occurred when updating the BackupDR DiskPool. "+
				"BackupDR Client Error: "+res.Status,
		)
	}

	// Map response body to schema and populate Computed attribute values
	state.ID = types.StringValue(respObject.Id)
	state.Href = types.StringValue(respObject.Href)
	state.Syncdate = types.Int64Value(respObject.Syncdate)
	state.Stale = types.BoolValue(respObject.Stale)
	state.Usedefaultsa = types.BoolValue(respObject.Usedefaultsa)
	state.Immutable = types.BoolValue(respObject.Immutable)
	state.Metadataonly = types.BoolValue(respObject.Metadataonly)
	state.State = types.StringValue(respObject.State)
	state.Srcid = types.StringValue(respObject.Srcid)
	state.Status = types.StringValue(respObject.Status)
	state.Mdiskgrp = types.StringValue(respObject.Mdiskgrp)
	state.Pooltypedisplayname = types.StringValue(respObject.Pooltypedisplayname)
	state.Srcid = types.StringValue(respObject.Srcid)
	state.Warnpct = types.Int64Value(int64(respObject.Warnpct))
	state.Modifydate = types.Int64Value(respObject.Modifydate)
	state.Safepct = types.Int64Value(int64(respObject.Safepct))
	state.Udsuid = types.Int64Value(int64(respObject.Udsuid))
	state.FreeMb = types.Int64Value(respObject.FreeMb)
	state.UsageMb = types.Int64Value(respObject.UsageMb)
	state.CapacityMb = types.Int64Value(respObject.CapacityMb)
	state.Pct = types.Float64Value(respObject.Pct)

	state.ApplianceClusterID = types.StringValue(respObject.Cluster.Clusterid)

	// state.Cluster.Href = types.StringValue(respObject.Cluster.Href)
	// state.Cluster.ID = types.StringValue(respObject.Cluster.Id)
	// state.Cluster.Clusterid = types.StringValue(respObject.Cluster.Clusterid)
	// state.Cluster.Serviceaccount = types.StringValue(respObject.Cluster.Serviceaccount)
	// state.Cluster.Zone = types.StringValue(respObject.Cluster.Zone)
	// state.Cluster.Region = types.StringValue(respObject.Cluster.Region)
	// state.Cluster.Projectid = types.StringValue(respObject.Cluster.Projectid)
	// state.Cluster.Version = types.StringValue(respObject.Cluster.Version)
	// state.Cluster.Name = types.StringValue(respObject.Cluster.Name)
	// state.Cluster.Type = types.StringValue(respObject.Cluster.Type_)
	// state.Cluster.Ipaddress = types.StringValue(respObject.Cluster.Ipaddress)
	// state.Cluster.Publicip = types.StringValue(respObject.Cluster.Publicip)
	// state.Cluster.Secureconnect = types.BoolValue(respObject.Cluster.Secureconnect)
	// state.Cluster.PkiBootstrapped = types.BoolValue(respObject.Cluster.PkiBootstrapped)
	// state.Cluster.Supportstatus = types.StringValue(respObject.Cluster.Supportstatus)
	// state.Cluster.Syncdate = types.Int64Value(respObject.Cluster.Syncdate)
	// state.Cluster.Stale = types.BoolValue(respObject.Cluster.Stale)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *diskpoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state diskPoolResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing Diskpool
	_, err := r.client.DiskPoolApi.DeleteDiskPool(r.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting DiskPool",
			"Could not delete diskpool, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *diskpoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
