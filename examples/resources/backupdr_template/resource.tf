resource "backupdr_template" "example" {
  name        = "<name>"
  description = "<SLA Template description>"
  policies = [
    {
      name              = "<policy name>",
      op                = "cloud",
      iscontinuous      = false,
      starttime         = 68400,
      endtime           = 67800,
      rpo               = "24",
      rpom              = "hours",
      retention         = "14",
      retentionm        = "days",
      priority          = "medium",
      sourcevault       = 0,
      targetvault       = 1,
      scheduletype      = "daily",
      repeatinterval    = "1",
      exclusiontype     = "none",
      exclusioninterval = "1",
      exclusion         = "none",
      selection         = "none"
    },
    {
      name              = "<policy name>",
      op                = "snap",
      endtime           = 25200,
      exclusion         = "none",
      exclusioninterval = "1",
      exclusiontype     = "none",
      iscontinuous      = false,
      priority          = "medium",
      repeatinterval    = "1",
      retention         = "2",
      retentionm        = "days",
      rpo               = "24",
      rpom              = "hours",
      scheduletype      = "daily",
      selection         = "none",
      starttime         = 68400,
  }]
}
