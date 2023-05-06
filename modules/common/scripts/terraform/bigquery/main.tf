
resource "google_bigquery_dataset" "ds_common" {
  dataset_id = var.dataset
  friendly_name = "Common Dataset"
  description = "Google Retail Data Model Common Dataset"
  location = var.location

  labels = {
    generated_by = "google"
    asset_type = "google.rdm.common.ds"
  }
}

