resource "backupdr_vcenter" "name" {
  friendlypath = "<friendly-name>"
  hostname     = "vcsa-000000.xxxxxxx.asia-northeast1.gve.goog" ## vcenter hostname
  ipaddress    = "10.10.0.2"                                    ## vcenter IP address
  hosttype     = "vcenter"
  hypervisoragent = {
    username = "CloudOwner@gve.local"
    password = "<vcenter-password>"
  }
  sources = [{ clusterid = "<appliance-clusterid>" }]
}
