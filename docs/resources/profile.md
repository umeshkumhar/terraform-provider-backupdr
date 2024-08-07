---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "backupdr_profile Resource - terraform-provider-backupdr"
subcategory: ""
description: |-
  A resource profile specifies the storage media for backups of application and VM data. The template and the resource profile that make up the backup plan dictate the type of application data policies to perform and where to store the application data backups (which storage pool is used). Resource Profiles define which snapshot pool (if needed) is used and which remote appliance data is replicated. In addition to templates, you also create resource profiles in the backup plans menu. Profiles define where to store data. Data can be stored in the following:
   - Primary Appliance: The backup/recovery appliance that the resource profile is created for. This includes selecting which appliance snapshot pool will be used.
   - Remote Appliance: The backup/recovery appliance used for remote replication. This remote appliance must be an appliance that is already paired to the selected local appliance. You can configure the remote appliance field only when one or more remote appliances are configured on the selected local appliance.
   - OnVault: Up to four object storage buckets defined by an OnVault storage pool. The OnVault pools store compressed and encrypted backups of application data on Google Cloud Storage.
  For more information, see Resource profile https://cloud.google.com/backup-disaster-recovery/docs/create-plan/create-resource-profiles.
---

# backupdr_profile (Resource)

A resource profile specifies the storage media for backups of application and VM data. The template and the resource profile that make up the backup plan dictate the type of application data policies to perform and where to store the application data backups (which storage pool is used). Resource Profiles define which snapshot pool (if needed) is used and which remote appliance data is replicated. In addition to templates, you also create resource profiles in the backup plans menu. Profiles define where to store data. Data can be stored in the following: 
 - Primary Appliance: The backup/recovery appliance that the resource profile is created for. This includes selecting which appliance snapshot pool will be used. 
 - Remote Appliance: The backup/recovery appliance used for remote replication. This remote appliance must be an appliance that is already paired to the selected local appliance. You can configure the remote appliance field only when one or more remote appliances are configured on the selected local appliance. 
 - OnVault: Up to four object storage buckets defined by an OnVault storage pool. The OnVault pools store compressed and encrypted backups of application data on Google Cloud Storage. 
For more information, see [Resource profile](https://cloud.google.com/backup-disaster-recovery/docs/create-plan/create-resource-profiles).

## Example Usage

```terraform
resource "backupdr_profile" "name" {
  cid             = "<appliance-id>"
  name            = "<name>"
  description     = "<profile description>"
  localnode       = "<appliance-node-pool>"
  performancepool = "act_per_pool000"
  remotenode      = "None"
  vaultpool = {
    id = "1234"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Provide a name for the resource profile.

### Optional

- `cid` (String) Provide the ID of the cluster - It is not the same as cluster ID.
- `description` (String) Provide a description for the resource profile.
- `localnode` (String) Provide the primary backup/recovery appliance name.
- `performancepool` (String) Provide a name of the snapshot (performance) pool. The default is act_per_pool000.
- `remotenode` (String) Provide the remote backup/recovery appliance name, when two appliances are to be configured to replicate snapshot data between them.
- `vaultpool` (Attributes) (see [below for nested schema](#nestedatt--vaultpool))
- `vaultpool2` (Attributes) (see [below for nested schema](#nestedatt--vaultpool2))
- `vaultpool3` (Attributes) (see [below for nested schema](#nestedatt--vaultpool3))
- `vaultpool4` (Attributes) (see [below for nested schema](#nestedatt--vaultpool4))

### Read-Only

- `clusterid` (String) It displays the backup/recovery appliance ID.
- `createdate` (Number) It displays the date when the resource profile was created.
- `dedupasyncnode` (String)
- `href` (String) It displays the API URI for backup plan profile.
- `id` (String) The ID of this resource.
- `modifydate` (Number) It displays the date when the resource profile details are modified.
- `srcid` (String) It displays the source ID on the appliance.
- `stale` (Boolean) It displays the possible values true or false.
- `syncdate` (Number) It displays the last sync date.

<a id="nestedatt--vaultpool"></a>
### Nested Schema for `vaultpool`

Optional:

- `id` (String) It displays the ID of the OnVault pool.

Read-Only:

- `href` (String) It displays the API URI for OnVault storage pool
- `name` (String) It displays the name of the OnVault pool used for resource profile.


<a id="nestedatt--vaultpool2"></a>
### Nested Schema for `vaultpool2`

Optional:

- `id` (String) It displays the ID of the OnVault pool 2.

Read-Only:

- `href` (String) It displays the API URI for OnVault storage pool.
- `name` (String) It displays the name of the OnVault pool 2 used for resource profile.


<a id="nestedatt--vaultpool3"></a>
### Nested Schema for `vaultpool3`

Optional:

- `id` (String) It displays the ID of the OnVault pool 3.

Read-Only:

- `href` (String) It displays the API URI for OnVault storage pool.
- `name` (String) It displays the name of the OnVault pool 3 used for resource profile.


<a id="nestedatt--vaultpool4"></a>
### Nested Schema for `vaultpool4`

Optional:

- `id` (String) It displays the ID of the OnVault pool 4.

Read-Only:

- `href` (String) It displays gthe API URI for OnVault storage pool.
- `name` (String) It displays the name of the OnVault pool 4 used for resource profile.
