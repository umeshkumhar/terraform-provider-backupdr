---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "backupdr_host Resource - terraform-provider-backupdr"
subcategory: ""
description: |-
  Manages an vCenter Host.
---

# backupdr_host (Resource)

Manages an vCenter Host.

## Example Usage

```terraform
resource "backupdr_host" "example" {
  friendlypath        = "<friendly-name>"
  appliance_clusterid = "<appliance-clusterid>"
  hostname            = "vcsa-000000.xxxxxxx.asia-northeast1.gve.goog" ## vcenter hostname
  ipaddress           = "10.10.0.2"                                    ## vcenter IP address
  hosttype            = "vcenter"
  hypervisoragent = {
    username = "CloudOwner@gve.local"
    password = "<vcenter-password>"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `appliance_clusterid` (String)
- `hostname` (String)
- `hosttype` (String)
- `ipaddress` (String)

### Optional

- `alternateip` (List of String)
- `friendlypath` (String)
- `hypervisoragent` (Attributes) (see [below for nested schema](#nestedatt--hypervisoragent))
- `multiregion` (String)
- `ostype_special` (String)

### Read-Only

- `autoupgrade` (String)
- `cert_revoked` (Boolean)
- `clusterid` (String)
- `dbauthentication` (Boolean)
- `diskpref` (String)
- `hasagent` (Boolean)
- `href` (String)
- `id` (String) The ID of this resource.
- `isclusterhost` (Boolean)
- `isclusternode` (Boolean)
- `isesxhost` (Boolean)
- `isproxyhost` (Boolean)
- `isshadowhost` (Boolean)
- `isvcenterhost` (Boolean)
- `isvm` (Boolean)
- `maxjobs` (Number)
- `modifydate` (Number)
- `name` (String)
- `originalhostid` (String)
- `pki_state` (String)
- `sourcecluster` (String)
- `srcid` (String)
- `svcname` (String)
- `transport` (String)
- `uniquename` (String)
- `zone` (String)

<a id="nestedatt--hypervisoragent"></a>
### Nested Schema for `hypervisoragent`

Required:

- `username` (String)

Optional:

- `password` (String, Sensitive)

Read-Only:

- `agenttype` (String)
- `hasalternatekey` (Boolean)
- `haspassword` (Boolean)
