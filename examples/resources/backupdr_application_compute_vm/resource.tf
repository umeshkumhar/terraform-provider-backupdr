resource "backupdr_application_compute_vm" "example" {
  cloudcredential     = "<cloud-credential-id>"
  appliance_clusterid = "<appliance-clusterid>"
  region              = "<gcp-region/zone>"
  projectid           = "<gcp-project>"
  vmids               = ["<gcp-vm-instanceid>"]
}

