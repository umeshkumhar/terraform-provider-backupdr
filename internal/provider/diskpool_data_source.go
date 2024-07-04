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
		MarkdownDescription: "This data source can be used to read information about a Backup and DR service diskpool.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Provide the ID of the storage pool.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the name of the storage pool.",
			},
			"href": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the URL to access the storage pools in the management console.",
			},
			"stale": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the state of the disk pool. Ok indicates the disk pool is healthy.",
			},
			"modifydate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the modified date in epoch time or date conversion.",
			},
			"syncdate": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the last sync date.",
			},
			"pooltype": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the type of storage pool (cloud/perf/primary/vault), where perf = snapshot type.",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the state of the disk pool. Ok indicates the disk pool is healthy.",
			},
			"srcid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the source ID on the appliance.",
			},
			"status": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the status of the disk pool. The green indicates the disk pool has available space.",
			},
			"mdiskgrp": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the storage pool name.",
			},
			"pooltypedisplayname": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the type of storage pool (cloud/perf/primary/vault), where perf = snapshot type.",
			},
			"warnpct": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the warn percent number, where alerts are generated once this threshold is met. Backup jobs and mounts can continue in this warning state.",
			},
			"safepct": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the safe percent number, where alerts are generated once this threshold is met. Backup jobs or mounts will not be possible where this value is met.",
			},
			"udsuid": schema.Int64Attribute{
				Computed: true,
			},
			"free_mb": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the free pool space in Megabytes.",
			},
			"usage_mb": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the current consumption of the pool in Megabytes.",
			},
			"capacity_mb": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the current pool capacity in Megabytes.",
			},
			"pct": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "It displays the percentage of the pool used.",
			},
			"usedefaultsa": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays true or false.",
			},
			"immutable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the immutable values - true or false.",
			},
			"metadataonly": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Identifies if this Storage pool is used for PD snapshot metadata or as a backup data storage pool. It displays true or false.",
			},
			"properties": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the key-value pair for the diskpool.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the storage pool attributes. It can be object size, use ssl, bucket name, or ID.",
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "It displays the storage pool values.",
						},
					},
				},
			},
			"appliance_clusterid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the backup/recovery appliance ID.",
			},
			"cluster": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the properties of the cluster.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the unique cluster id used in api call.",
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the name of the storage pool.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for disk pool.",
					},
					"clusterid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the backup/recovery appliance ID as shown in the **Management console** > **Manage** > **Appliances** page.",
					},
					"serviceaccount": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the GCP service account used for OnVault pool access.",
					},
					"zone": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the zone where the appliance is located.",
					},
					"region": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the region where the OnVault pool is created.",
					},
					"projectid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the project ID used to create the OnVault pool.",
					},
					"version": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the version of the backup appliance.",
					},
					"type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the appliance type.",
					},
					"ipaddress": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the IP address of the backup/recovery appliance ID.",
					},
					"publicip": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the public IP of the backup/recovery appliance ID.",
					},
					"supportstatus": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the appliance up to date with latest patches or updates status. It can be true or false.",
					},
					"secureconnect": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible values for secure connect as true or false.",
					},
					"pkibootstrapped": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays if the PKI boot strap is enabled or not.",
					},
					"stale": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible values true or false.",
					},
					"syncdate": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "It displays the last sync date between appliance and management console.",
					},
				},
			},
			"vaultprops": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "It displays the properties of OnVault.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the unique ID for objects.",
					},
					"bucket": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the OnVault pool bucket ID.",
					},
					"href": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the API URI for disk pool.",
					},
					"region": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the region where the OnVault pool is created.",
					},
					"compression": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible compression values true or false.",
					},
					"stale": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "It displays the possible values true or false.",
					},
					"syncdate": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "It displays the last sync date in epoch converted format.",
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
