
variable "project" {
  type = string
  description = "GCP Project"
  nullable = false
}

variable "bq_owner" {
  type = string
  description = "The email of the database owner"
  nullable = false
}

variable "bq_reader_domain" {
  type = string
  description = "The domain of services allowed to read the table."
  nullable = false
}

variable "region" {
  type = string
  description = "GCP Region"
  default = "us-central1"
}

variable "zone" {
  type = string
  description = "GCP Zone"
  default = "us-central1-c"
}