terraform {
  required_providers {
    backupdr = {
      source = "example.com/google/backupdr"
    }
    google = {
      source = "hashicorp/google"
    }
  }
}


provider "google" {}

data "google_client_config" "default" {}

provider "backupdr" {
  endpoint     = "https://bmc-116926924342-sugw2kk2-dot-us-central1.backupdr.googleusercontent.com"
  access_token = data.google_client_config.default.access_token
}

###### DiskPool ##############x
# data "backupdr_diskpool" "example" {
#   id = "62448"
# }

# resource "backupdr_diskpool" "name" {
#   name                = "test01234qqqqqQq"
#   pooltype            = "vault"
#   appliance_clusterid = "145353943664"

#   properties = [
#     { key = "accessId", value = "backup-recovery-appliance001@drip-site-02.iam.gserviceaccount.com" },
#     { key = "bucket", value = "backupdr-demo" },
#     { key = "compression", value = "true" },
#     { key = "vaulttype", value = "GoogleNative" }
#   ]
# }

###### profile ##################
# data "backupdr_profile" "example1" {
#   id = "21032"
# }

# data "backupdr_profiles" "name" {}

resource "backupdr_profile" "name" {
  cid             = "6917"
  name            = "test012301101wlo"
  description     = "test profile0123"
  localnode       = "site1-appliance-40967"
  performancepool = "act_per_pool000"
  remotenode      = "None"
  # vaultpool2 = {
  #   id = "0"
  # }
}
#### SLT #####################
# data "backupdr_template" "example" {
#   id = "18123"
# }

# resource "backupdr_template" "name" {
#   name        = "um321000111wv101"
#   description = "proper umesh descripition"
#   policies = [
#     # {
#     #   endtime           = 25200,
#     #   exclusion         = "none",
#     #   exclusioninterval = "1",
#     #   exclusiontype     = "none",
#     #   iscontinuous      = false,
#     #   name              = "pol1",
#     #   op                = "snap",
#     #   priority          = "medium",
#     #   repeatinterval    = "1",
#     #   retention         = "99",
#     #   # "sourcevault"     = 0,
#     #   # "targetvault"     = 1,
#     #   retentionm   = "days",
#     #   rpo          = "24",
#     #   rpom         = "hours",
#     #   scheduletype = "daily",
#     #   policytype   = "windowed",
#     #   # scheduling        = "windowed",
#     #   selection = "none",
#     #   starttime = 68400,
#     #   # temp-id="362",
#     # },
#     {
#       endtime           = 25200,
#       exclusion         = "none",
#       exclusioninterval = "1",
#       exclusiontype     = "none",
#       iscontinuous      = false,
#       name              = "pol77",
#       # "sourcevault"     = null,
#       # "targetvault"     = null,
#       op             = "snap",
#       priority       = "medium",
#       repeatinterval = "1",
#       retention      = "78",
#       retentionm     = "days",
#       rpo            = "24",
#       rpom           = "hours",
#       scheduletype   = "daily",
#       policytype     = "windowed",
#       # scheduling        = "windowed",
#       selection = "none",
#       starttime = 68400,
#       # temp-id="362",
#     }
#   ]
# }

#####  Appliance  #####################

# data "backupdr_appliance" "name" {
#   id = "86122"
# }

# data "backupdr_appliances" "name" {}

#####  CloudCredential  #####################
# data "backupdr_cloudcredential" "name" {
#   id = "49385"
# }

# data "backupdr_cloudcredentials" "name" {}

#####  Cloud VMs  #####################
# resource "backupdr_application_compute_vm" "name" {
#   cloudcredential     = "86139"
#   appliance_clusterid = "145353943664"

#   region    = "us-central1-a"
#   projectid = "drip-site-02"
#   vmids     = ["745278411111110556", "824237111111118021", "5589291111114382135"]
# }

# #####  vCenter  #####################
# resource "backupdr_host" "example" {
#   friendlypath = "vcsa-303836.fecaf039.asia-northeast1.gve.goog"
#   hostname     = "vcsa-303836.fecaf039.asia-northeast1.gve.goog"
#   hosttype     = "vcenter"
#   hypervisoragent = {
#     username = "CloudOwner@gve.local"
#     password = "xxxxxxxxxxxx"
#   }
#   ipaddress           = "10.10.0.2"
#   appliance_clusterid = "145353943664"
#   # sources   = [{ clusterid = "145353943664" }]
# }

# resource "backupdr_application_vmware_vm" "name" {
#   cluster_name = "bcdr"
#   appliance_id = "86122"
#   vcenter_id   = backupdr_host.example.id
#   vms          = ["502309d6-8a3a-410e-2d3c-4573f35300d3"]
# }


#### SLA #####################
# data "backupdr_plan" "example" {
#   id = "64274"
# }

## backup cloud_vm 
# resource "backupdr_plan" "cloud_backup" {
#   count       = length(backupdr_application_compute_vm.name.applications)
#   description = "test sla"
#   scheduleoff = "true"
#   application = {
#     id = backupdr_application_compute_vm.name.applications[count.index]
#   }
#   slt = {
#     id = backupdr_template.name.id
#   }
#   slp = {
#     id = backupdr_profile.name.id
#   }
# }

# ## backup vmware_vm 
# resource "backupdr_plan" "vmware_backup" {
#   count       = length(backupdr_application_vmware_vm.name.applications)
#   description = "test vmware sla"
#   scheduleoff = "false"
#   application = {
#     id = backupdr_application_vmware_vm.name.applications[count.index]
#   }
#   slt = {
#     id = backupdr_template.name.id
#   }
#   slp = {
#     id = "168884" #backupdr_profile.name.id
#   }
# }




