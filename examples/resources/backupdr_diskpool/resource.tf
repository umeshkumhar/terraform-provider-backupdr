resource "backupdr_diskpool" "name" {
  name     = "<name>"
  pooltype = "vault"
  cluster = {
    clusterid = "<appliance-clusterid>"
  }
  properties = [
    {
      key   = "accessId"
      value = "<service-account-email>"
    },
    {
      key   = "bucket"
      value = "<gcs-bucket-name>"
    },
    {
      key   = "compression"
      value = "true"
    },
    {
      key   = "vaulttype"
      value = "GoogleNative"
    }
  ]
}
