resource "backupdr_application_compute_vm" "example" {
  cloudcredential = "<cloud-credential-id>"
  cluster = {
    clusterid = "<appliance-clusterid>"
  }
  region    = "<location>"
  projectid = "<gcp-project>"
  vmids     = ["<gcp-vm-instanceid>"]
}

