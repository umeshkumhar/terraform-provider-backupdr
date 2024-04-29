resource "backupdr_vcenter_addvms" "name" {
  cluster_name = "<vcenter-cluster-name>"
  cluster      = "<appliance-id>"
  vcenter_id   = "<vcenter-host-id>"
  vms          = ["<vcenter-vm-uuid>"]
}
