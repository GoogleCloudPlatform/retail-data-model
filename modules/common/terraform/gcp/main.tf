provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

module "bigquery" {
  source = "./modules/bigquery"
  bq_owner = var.bq_owner
  bq_reader_domain = var.bq_reader_domain
}