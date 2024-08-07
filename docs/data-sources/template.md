---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "backupdr_template Data Source - terraform-provider-backupdr"
subcategory: ""
description: |-
  This data source can be used to read information about a backup template. It displays the backup template ID as shown in the Management console > Backup Plans > Templates page.
---

# backupdr_template (Data Source)

This data source can be used to read information about a backup template. It displays the backup template ID as shown in the **Management console** > **Backup Plans** > **Templates** page.

## Example Usage

```terraform
data "backupdr_template" "example" {
  ## Replace with any existing SLA Template ID 
  id = "63512"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Provide the backup template ID.

### Read-Only

- `description` (String) It displays the description for the backup template.
- `href` (String) It displays the API URI for Backup Plan template.
- `managedbyagm` (Boolean)
- `name` (String) It displays the name of the backup template.
- `option_href` (String) It displays the API URI for Backup Plan template options.
- `override` (String) It displays the template override settings. Setting “Yes” will allow the policies set in this template to be overridden per-application. Setting “No” will enforce the policies as configured in this template without allowing any per-application overrides.
- `policies` (Attributes List) It displays the policy details. (see [below for nested schema](#nestedatt--policies))
- `policy_href` (String) It displays the backup policy ID.
- `sourcename` (String) It displays the source name. It should match the name value.
- `usedbycloudapp` (Boolean) It displays if the template is used by applications or not - true/false.

<a id="nestedatt--policies"></a>
### Nested Schema for `policies`

Read-Only:

- `description` (String) It displays the description for the backup policy.
- `encrypt` (String) It displays the encryption identifier.
- `endtime` (String) It displays the end time for the backup plan.
- `exclusion` (String) It displays specific days, days of week, month and days of month excluded for backup snapshots.
- `exclusioninterval` (String) It displays the exclusion interval for the template. Normally set to 1.
- `exclusiontype` (String) It displays the exclusion type as daily, weekly, monthly, or yearly.
- `href` (String) It displays the url to access the backup plan template href of the policy.
- `id` (String) It displays the backup plan policy ID.
- `iscontinuous` (Boolean) It displays boolean value true or false if the policy setting for continuous mode or windowed.
- `name` (String) It displays the name of the policy.
- `op` (String) It displays the operation type. Normally set to snap, DirectOnVault, or stream_snap.
- `policytype` (String) It displays the backup policy type. It can be snapshot, direct to OnVault, OnVault replication, mirror, and OnVault policy.
- `priority` (String) It displays the application priority. It can be medium, high or low. The default job priority is medium, but you can change the priority to high or low.
- `remoteretention` (Number) It displays for mirror policy options.
- `repeatinterval` (String) It displays the interval value. Normally set to 1.
- `reptype` (String) It displays for mirror policy options.
- `retention` (String) It displays how long the image is set for retention.
- `retentionm` (String) It displays the retention in days, weeks, months, or years.
- `rpo` (String) It displays how often to run policy again. 24 is once per day.
- `rpom` (String) It displays the PRP in hours. You can also set the RPO in minutes.
- `scheduletype` (String) It displays the schedule type as daily, weekly, monthly or yearly.
- `selection` (String) It displays the days to run the scheduled job. For example, weekly jobs on Sunday - days of week as sun.
- `sourcevault` (Number) It displays the OnVault disk pool id. You can get the from the **Management console** > **Manage** > **Storage Pools**, then enabling visibility of the ID column.
- `starttime` (String) It displays the start time for the backup plan in decimal format: total seconds = (hours x 3600) + (minutes + 60) + seconds.
- `targetvault` (Number) It displays the OnVault disk pool id. You can get the from the **Management console** > **Manage** > **Storage Pools**, then enabling visibility of the ID column.
- `truncatelog` (String) It displays the log truncation options. This may not work as required in advanced options.
- `verification` (Boolean) It displays the verification values as true or false.
- `verifychoice` (String) It displays the empty value by default - to be used in future versions.
