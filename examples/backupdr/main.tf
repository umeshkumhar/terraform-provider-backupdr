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
  endpoint     = "https://bmc-393422088677-hi2uw6tg-dot-us-central1.backupdr.googleusercontent.com"
  access_token = data.google_client_config.default.access_token
}

###### DiskPool ##############
data "backupdr_diskpool" "example" {
  id = "62448"
}

resource "backupdr_diskpool" "name" {
  name     = "test01234"
  pooltype = "vault"
  cluster = {
    clusterid = "143045097332"
  }
  properties = [
    { key = "accessId", value = "backup-recovery-appliance001@drip-site-02.iam.gserviceaccount.com" },
    { key = "bucket", value = "drip-demo" },
    { key = "compression", value = "true" },
    { key = "vaulttype", value = "GoogleNative" }
  ]
}

###### profile ##################
data "backupdr_profile" "example1" {
  id = "21032"
}

data "backupdr_profiles" "name" {}

resource "backupdr_profile" "name" {
  cid             = "4197"
  description     = "test profile0123"
  localnode       = "backup-recovery-appliance001-frkr"
  name            = "test0123"
  performancepool = "act_per_pool000"
  remotenode      = "None"
  vaultpool = {
    id = backupdr_diskpool.name.id
  }
}
#### SLT #####################
data "backupdr_template" "example" {
  id = "18123"
}

resource "backupdr_template" "name" {
  name        = "um3210"
  description = "from TF"
  policies = [{
    "op"             = "cloud",
    "name"           = "onvault-pol",
    "iscontinuous"   = false,
    "starttime"      = 68400,
    "endtime"        = 67800,
    "rpo"            = "24",
    "rpom"           = "hours",
    "retention"      = "14",
    "retentionm"     = "days",
    "priority"       = "medium",
    "sourcevault"    = 0,
    "targetvault"    = 1,
    "temp-id"        = "431",
    "scheduletype"   = "daily",
    "repeatinterval" = "1",
    # "scheduling"        = "windowed",
    "exclusiontype"     = "none",
    "exclusioninterval" = "1",
    "exclusion"         = "none",
    "selection"         = "none"
    }, {
    endtime           = 25200,
    exclusion         = "none",
    exclusioninterval = "1",
    exclusiontype     = "none",
    iscontinuous      = false,
    name              = "pol1",
    op                = "snap",
    priority          = "medium",
    repeatinterval    = "1",
    retention         = "2",
    retentionm        = "days",
    rpo               = "24",
    rpom              = "hours",
    scheduletype      = "daily",
    # scheduling        = "windowed",
    selection = "none",
    starttime = 68400,
    # temp-id="362",
  }]
}

#### SLA #####################
data "backupdr_plan" "example" {
  id = "64274"
}

resource "backupdr_plan" "name" {
  description = "test sla"
  scheduleoff = "true"
  application = {
    id = 22207
  }
  slt = {
    id = 18123
  }
  slp = {
    id = 4208
  }
}

#####  Appliance  #####################

data "backupdr_appliance" "name" {
  id = "86122"
}

data "backupdr_appliances" "name" {}

#####  CloudCredential  #####################
data "backupdr_cloudcredential" "name" {
  id = "49385"
}

data "backupdr_cloudcredential_all" "name" {}

#####  Cloud VMs  #####################
resource "backupdr_application_compute_vm" "name" {
  cloudcredential = "86139"
  cluster = {
    clusterid = "145353943664"
  }
  region    = "us-central1-a"
  projectid = "drip-site-02"
  vmids     = ["745278443586790556", "8242374309568738021", "5589298639124382135"]
}

#####  vCenter  #####################
resource "backupdr_host" "example" {
  friendlypath = "vcsa-303836.fecaf039.asia-northeast1.gve.goog"
  hostname     = "vcsa-303836.fecaf039.asia-northeast1.gve.goog"
  hosttype     = "vcenter"
  hypervisoragent = {
    username = "CloudOwner@gve.local"
    password = "ZReJ*c0NBpVCYvFL"
  }
  ipaddress = "10.10.0.2"
  sources   = [{ clusterid = "145353943664" }]
}

resource "backupdr_application_vmware_vm" "name" {
  cluster_name = "bcdr"
  cluster      = "86122"
  vcenter_id   = backupdr_host.example.id
  vms          = ["502309d6-8a3a-410e-2d3c-4573f35300d3"]
}

############### Update is failing ************

