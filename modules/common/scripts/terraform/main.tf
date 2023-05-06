terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.51.0"
    }
  }
}

provider "google" {
  project = var.project
  region = var.region
  zone = var.zone
}

module "bigquery" {
  source = "./bigquery"
}