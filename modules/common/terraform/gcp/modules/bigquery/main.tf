
resource "google_bigquery_dataset" "common_ds" {
  dataset_id = "common_ds"
  friendly_name = "Common Dataset"
  description = "The dataset used by common services."
  location = "US"

  labels = {
    provider = "google_cloud_retail"
  }

  access {
    role = "OWNER"
    user_by_email = var.bq_owner

  }

  access {
    role = "reader"
    domain = var.bq_reader_domain
  }
}

resource "google_bigquery_table" "audit_record" {
  dataset_id = google_bigquery_dataset.common_ds.dataset_id
  table_id   = "audit_record"
  time_partitioning {
    type = "DAY"
  }
  labels = {
    provider = "google_cloud_retail"
  }
  schema = file("${path.module}/audit_record.json")
}

resource "google_bigquery_table" "countries" {
  dataset_id = google_bigquery_dataset.common_ds.dataset_id
  table_id   = "countries"
  time_partitioning {
    type = "DAY"
  }
  labels = {
    provider = "google_cloud_retail"
  }
  schema = file("${path.module}/countries.json")
}

resource "google_bigquery_table" "country_subdivisions" {
  dataset_id = google_bigquery_dataset.common_ds.dataset_id
  table_id   = "country_subdivisions"
  time_partitioning {
    type = "DAY"
  }
  labels = {
    provider = "google_cloud_retail"
  }
  schema = file("${path.module}/country_subdivision.json")
}

resource "google_bigquery_table" "icao_codes" {
  dataset_id = google_bigquery_dataset.common_ds.dataset_id
  table_id   = "icao_codes"
  time_partitioning {
    type = "DAY"
  }
  labels = {
    provider = "google_cloud_retail"
  }
  schema = file("${path.module}/icao_codes.json")
}

resource "google_service_account" "bqowner" {
  account_id = "bqowner"
}