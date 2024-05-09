resource "backupdr_application_vmware_vm" "name" {
  cluster_name = "<vcenter-cluster-name>"
  appliance_id = "<appliance-id>"
  vcenter_id   = "<vcenter-host-id>"
  vms          = ["<vcenter-vm-uuid>"]
}
