
## Use client config datasource to fetch access_token
data "google_client_config" "default" {}

provider "backupdr" {
  ## Replace with Backup DR Service Endpoint
  endpoint     = "https://bmc-30000000000-xxxxxxxxx-dot-us-central1.backupdr.googleusercontent.com"
  access_token = data.google_client_config.default.access_token
}
