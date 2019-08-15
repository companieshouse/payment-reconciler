provider "aws" {
  region = "${var.aws_region}"
}
terraform {
  backend "s3" {
    
  }
}

data "terraform_remote_state" "networks" {
    backend = "s3"
    config {
        bucket = "${var.network_state_bucket}"
        key = "${var.state_prefix}/${var.farm}/${var.farm}.tfstate"
        region = "${var.aws_region}"
    }
}

module "lambda" {
  source                        = "module-lambda"
  project_name                  = "${var.project_name}"
  handler                       = "${var.handler}"
  memory_megabytes              = "${var.memory_megabytes}"
  runtime                       = "${var.runtime}"
  timeout_seconds               = "${var.timeout_seconds}" 
  payment_reconciler_bucket     = "${var.payment_reconciler_bucket}"
  release_version               = "${var.release_version}"
  release_bucket_name           = "${var.release_bucket_name}"
  execution_role                = "${module.lambda-roles.execution_role}"
  application_ids               = "${data.terraform_remote_state.networks.application_ids}"
  security_group_ids            = "${module.security-group.security_group_ids}"
}

module "lambda-roles" {
  source                    = "module-lambda-roles"
  project_name              = "${var.project_name}"
  payment_reconciler_bucket = "${var.payment_reconciler_bucket}"
  config_bucket_name        = "${var.config_bucket_name}"
  env                       = "${var.env}"
  app_env_directory         = "${var.app_env_directory}"
}

module "security-group" {
  source                    = "module-security-group"
}

module "cloud-watch" {
  source                        = "module-cloud-watch"
  project_name                  = "${var.project_name}"
  payment_reconciler_bucket     = "${var.payment_reconciler_bucket}"
  arn                           = "${module.lambda.arn}"
  env                           = "${var.env}"
}
