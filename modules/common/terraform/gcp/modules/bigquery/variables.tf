
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

