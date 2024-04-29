resource "backupdr_sla" "name" {
  description = "<SLA description>"
  scheduleoff = "true"
  application = {
    id = 1234 ## <application-id>
  }
  slt = {
    id = 18123 ## <sla-template-id>
  }
  slp = {
    id = 4208 ## <sla-profile-id>
  }
}
