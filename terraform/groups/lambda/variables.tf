variable "aws_region" {
}
variable "aws_profile" {
  type        = string
  description = "The AWS profile to use for deployment."
}

variable "handler" {
  default = "payment-reconciler"
}

variable "memory_megabytes" {
  default = "320"
}

variable "release_bucket_name" {
  default = "release.ch.gov.uk"
}

variable "runtime" {
  default = "go1.x"
}

variable "timeout_seconds" {
  default = "6"
}

variable "environment" {
}

variable "release_version" {
}

variable "aws_bucket" {
}

variable "service" {
  default = "payment-reconciler"
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "cron_schedule" {
}

variable "workspace_key_prefix" {
}
variable "state_file_name" {
}

# Vault
variable "vault_username" {
  type        = string
  description = "The username used by the Vault provider."
}
variable "vault_password" {
  type        = string
  description = "The password used by the Vault provider."
}
