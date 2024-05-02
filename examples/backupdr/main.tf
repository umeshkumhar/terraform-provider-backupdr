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
###### SLP ##################
data "backupdr_slp" "example1" {
  id = "21032"
}

data "backupdr_slp_all" "name" {}

resource "backupdr_slp" "name" {
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
data "backupdr_slt" "example" {
  id = "63512"
}

resource "backupdr_slt" "name" {
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
data "backupdr_sla" "example" {
  id = "64274"
}

resource "backupdr_sla" "name" {
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

#####  vCenter  #####################
resource "backupdr_vcenter" "name" {
  friendlypath = "vcsa-303836.fecaf039.asia-northeast1.gve.goog"
  hostname     = "vcsa-303836.fecaf039.asia-northeast1.gve.goog"
  hosttype     = "vcenter"
  hypervisoragent = {
    username = "CloudOwner@gve.local"
    password = "ZReJ*c0NBpVCYvFL"
  }
  ipaddress = "10.10.0.2"
  # orglist   = []
  sources = [{ clusterid = "145353943664" }]

}

resource "backupdr_vcenter_addvms" "name" {
  cluster_name = "bcdr"
  cluster      = "86122"
  vcenter_id   = backupdr_vcenter.name.id
  vms          = ["502309d6-8a3a-410e-2d3c-4573f35300d3"]
}

#####  Appliance  #####################

data "backupdr_appliance" "name" {
  id = "86122"
}

data "backupdr_appliance_all" "name" {}

#####  CloudCredential  #####################
data "backupdr_cloudcredential" "name" {
  id = "49385"
}

data "backupdr_cloudcredential_all" "name" {}

#####  Cloud VMs  #####################
resource "backupdr_cloud_addvms" "name" {
  cloudcredential = "86139"
  cluster = {
    clusterid = "145353943664"
  }
  region    = "us-central1-c"
  projectid = "drip-site-02"
  vmids     = ["745278443586790556"]
}

############### Update is failing ************


