variable "aws_region" {
  type        = string
  description = "AWS Region"
}

variable "aws_profile" {
  type        = string
  description = "The AWS profile to use for deployment."
}

variable "handler" {
  type        = string
  default     = "payment-reconciler"
  description = "The entrypoint in the Lambda funtion."
}

variable "memory_megabytes" {
  type        = string
  default     = "320"
  description = "The amount of memory to allocate to the Lambda function"
}

variable "release_bucket_name" {
  type        = string
  default     = "release.ch.gov.uk"
  description = "The S3 release bucket location containing the function code. "
}

variable "release_version" {
  type        = string
  description = "The version of the function code."
}

variable "runtime" {
  type        = string
  default     = "go1.x"
  description = "The Lambda function language / runtime."
}

variable "timeout_seconds" {
  type        = string
  default     = "6"
  description = "The amount of time the Lambda function has to run in seconds."
}

variable "environment" {
  type        = string
  description = "The name of the environment to deploy."
}

variable "service" {
  type        = string
  default     = "payment-reconciler"
  description = "The name of the service being deployed."
}

variable "cron_schedule" {
  type        = string
  description = "CloudWatch cron schedule expression for calling the Lambda function."
}

variable open_lambda_environment_variables {
  type        = map(string)
  description = "Lambda environment variables that do not require encryption."
  default     = {}
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

# Network Remote State
variable "remote_state_bucket" {
  type        = string
  description = "Remote state location for the network to deploy the Lambda to."
}

variable "remote_state_key" {
  type        = string
  description = "Remote state location for the network to deploy the Lambda to."
}
