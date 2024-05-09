resource "backupdr_profile" "name" {
  cid             = "<appliance-id>"
  name            = "<name>"
  description     = "<profile description>"
  localnode       = "<appliance-node-pool>"
  performancepool = "act_per_pool000"
  remotenode      = "None"
  vaultpool = {
    id = "1234"
  }
}
