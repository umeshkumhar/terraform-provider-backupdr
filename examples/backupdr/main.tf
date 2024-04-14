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
    { key : "accessId", value : "backup-recovery-appliance001@drip-site-02.iam.gserviceaccount.com" },
    { key : "bucket", value : "drip-demo" },
    { key : "compression", value : "true" },
    { key : "vaulttype", value : "GoogleNative" }
  ]
}
###### SLP ##################
data "backupdr_slp" "example1" {
  id = "21032"
}

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
    "op"                = "cloud",
    "name"              = "onvault-pol",
    "iscontinuous"      = false,
    "starttime"         = 68400,
    "endtime"           = 67800,
    "rpo"               = "24",
    "rpom"              = "hours",
    "retention"         = "14",
    "retentionm"        = "days",
    "priority"          = "medium",
    "sourcevault"       = 0,
    "targetvault"       = 1,
    "temp-id"           = "431",
    "scheduletype"      = "daily",
    "repeatinterval"    = "1",
    "scheduling"        = "windowed",
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
    scheduling        = "windowed",
    selection         = "none",
    starttime         = 68400,
    # temp-id="362",
  }]
}

############### Update is failing ************

