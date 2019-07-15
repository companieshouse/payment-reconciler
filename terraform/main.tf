provider "aws" {
  region = "${var.aws_region}"
}
terraform {
  backend "s3" {}
}
module "lambda" {
  source                        = "module-lambda"
  project_name                  = "${var.project_name}"
  handler                       = "${var.handler}"
  memory_megabytes              = "${var.memory_megabytes}"
  runtime                       = "${var.runtime}"
  timeout_seconds               = "${var.timeout_seconds}" 
  payment_reconciler_bucket     = "${var.payment_reconciler_bucket}"
  log_level                     = "${var.log_level}"
  release_version               = "${var.release_version}"
  release_bucket_name           = "${var.release_bucket_name}"
  execution_role                = "${module.lambda-roles.execution_role}"
}

module "lambda-roles" {
  source                   = "module-lambda-roles"
  project_name             = "${var.project_name}"
  file_transfer_bucket     = "${var.payment_reconciler_bucket}"
}