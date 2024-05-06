resource "backupdr_application_vmware_vm" "name" {
  cluster_name = "<vcenter-cluster-name>"
  cluster      = "<appliance-id>"
  vcenter_id   = "<vcenter-host-id>"
  vms          = ["<vcenter-vm-uuid>"]
}
