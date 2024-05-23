package provider

import (
	"context"

	backupdr "github.com/umeshkumhar/backupdr-client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &diskpoolDataSource{}
	_ datasource.DataSourceWithConfigure = &diskpoolDataSource{}
)

// diskpoolDataSource is the data source implementation.
type diskpoolDataSource struct {
	client  *backupdr.APIClient
	authCtx context.Context
}

// NewDiskpoolDataSource - Datasource for DiskPool
func NewDiskpoolDataSource() datasource.DataSource {
	return &diskpoolDataSource{}
}

// Configure adds the provider configured client to the data source.
func (d *diskpoolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*backupdrProvider).client
	d.authCtx = req.ProviderData.(*backupdrProvider).authCtx
}

func (d *diskpoolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_diskpool"
}

func (d *diskpoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				Computed: true,
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
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed: true,
						},
						"value": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"appliance_clusterid": schema.StringAttribute{
				Computed: true,
			},
			"cluster": schema.SingleNestedAttribute{
				Computed: true,
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
				Computed: true,
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

func (d *diskpoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state diskPoolResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	diskpool, res, err := d.client.DiskPoolApi.GetDiskPool(d.authCtx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR DiskPool",
			err.Error(),
		)
		return
	}

	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Read BackupDR DiskPool",
			res.Status,
		)
		return
	}

	// Map response body to model
	state = diskPoolResourceModel{
		Name:                types.StringValue(diskpool.Name),
		ID:                  types.StringValue(diskpool.Id),
		Pooltype:            types.StringValue(diskpool.Pooltype),
		State:               types.StringValue(diskpool.State),
		Srcid:               types.StringValue(diskpool.Srcid),
		Status:              types.StringValue(diskpool.Status),
		Mdiskgrp:            types.StringValue(diskpool.Mdiskgrp),
		Pooltypedisplayname: types.StringValue(diskpool.Pooltypedisplayname),
		Href:                types.StringValue(diskpool.Href),
		Usedefaultsa:        types.BoolValue(diskpool.Usedefaultsa),
		Immutable:           types.BoolValue(diskpool.Immutable),
		Metadataonly:        types.BoolValue(diskpool.Metadataonly),
		Stale:               types.BoolValue(diskpool.Stale),
		Modifydate:          types.Int64Value(diskpool.Modifydate),
		Warnpct:             types.Int64Value(int64(diskpool.Warnpct)),
		Safepct:             types.Int64Value(int64(diskpool.Safepct)),
		Udsuid:              types.Int64Value(int64(diskpool.Udsuid)),
		FreeMb:              types.Int64Value(diskpool.FreeMb),
		UsageMb:             types.Int64Value(diskpool.UsageMb),
		CapacityMb:          types.Int64Value(diskpool.CapacityMb),
		Syncdate:            types.Int64Value(diskpool.Syncdate),
	}

	for _, prop := range diskpool.Properties {
		state.Properties = append(state.Properties, keyValueRestModel{
			Key:   types.StringValue(prop.Key),
			Value: types.StringValue(prop.Value),
		})
	}

	if diskpool.Cluster != nil {
		state.ApplianceClusterID = types.StringValue(diskpool.Cluster.Clusterid)
		state.Cluster = &ClusterRest{
			Clusterid:       types.StringValue(diskpool.Cluster.Clusterid),
			Serviceaccount:  types.StringValue(diskpool.Cluster.Serviceaccount),
			Zone:            types.StringValue(diskpool.Cluster.Zone),
			Region:          types.StringValue(diskpool.Cluster.Region),
			Projectid:       types.StringValue(diskpool.Cluster.Projectid),
			Version:         types.StringValue(diskpool.Cluster.Version),
			Name:            types.StringValue(diskpool.Cluster.Name),
			Type:            types.StringValue(diskpool.Cluster.Type_),
			Ipaddress:       types.StringValue(diskpool.Cluster.Ipaddress),
			Publicip:        types.StringValue(diskpool.Cluster.Publicip),
			Secureconnect:   types.BoolValue(diskpool.Cluster.Secureconnect),
			PkiBootstrapped: types.BoolValue(diskpool.Cluster.PkiBootstrapped),
			Supportstatus:   types.StringValue(diskpool.Cluster.Supportstatus),
			ID:              types.StringValue(diskpool.Cluster.Id),
			Href:            types.StringValue(diskpool.Cluster.Href),
			Syncdate:        types.Int64Value(diskpool.Cluster.Syncdate),
			Stale:           types.BoolValue(diskpool.Cluster.Stale),
		}
	}

	if diskpool.Vaultprops != nil {
		state.Vaultprops = &vaultPropsRest{
			Bucket:      types.StringValue(diskpool.Vaultprops.Bucket),
			Compression: types.BoolValue(diskpool.Vaultprops.Compression),
			Region:      types.StringValue(diskpool.Vaultprops.Region),
			ID:          types.StringValue(diskpool.Vaultprops.Id),
			Href:        types.StringValue(diskpool.Vaultprops.Href),
			Syncdate:    types.Int64Value(diskpool.Vaultprops.Syncdate),
			Stale:       types.BoolValue(diskpool.Vaultprops.Stale),
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

}
