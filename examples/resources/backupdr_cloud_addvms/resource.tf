resource "backupdr_cloud_addvms" "name" {
  cloudcredential = "<cloud-credential-id>"
  cluster = {
    clusterid = "<appliance-clusterid>"
  }
  region    = "<location>"
  projectid = "<gcp-project>"
  vmids     = ["<gcp-vm-instanceid>"]
}

