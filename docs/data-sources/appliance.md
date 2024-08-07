---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "backupdr_appliance Data Source - terraform-provider-backupdr"
subcategory: ""
description: |-
  This data source can be used to read information about a backup/recovery Appliance. It displays the backup/recovery appliance ID as shown in the Management console > Manage > Appliances page.
---

# backupdr_appliance (Data Source)

This data source can be used to read information about a backup/recovery Appliance. It displays the backup/recovery appliance ID as shown in the **Management console** > **Manage** > **Appliances** page.

## Example Usage

```terraform
data "backupdr_appliance" "example" {
  ## replace with appliance ID
  id = "86122"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Provide the ID of the appliance.

### Read-Only

- `clusterid` (String) It displays the backup/recovery appliance ID as shown in the **Management console** > **Manage** > **Appliances** page.
- `href` (String) It displays the URL to access the storage pools in the management console.
- `ipaddress` (String) It displays the IP address of the backup/recovery appliance ID.
- `name` (String) It displays the name of the backup/recovery appliance ID.
- `pkibootstrapped` (Boolean) It displays if the PKI boot strap is enabled or not.
- `projectid` (String) It displays the project ID of the backup/recovery appliance ID.
- `publicip` (String) It displays the public IP of the backup/recovery appliance ID.
- `region` (String) It displays the region where the backup/recovery appliance is created.
- `secureconnect` (Boolean) It displays the possible values for secure connect as true or false.
- `serviceaccount` (String) It displays the GCP service account used for backup/recovery appliances.
- `stale` (Boolean) It displays the possible values true or false.
- `supportstatus` (String) It displays the appliance up to date with latest patches or updates status. It can be true or false.
- `type` (String) It displays the appliance type.
- `version` (String) It displays the version of the backup appliance.
- `zone` (String) It displays the zone where the appliance is located.
