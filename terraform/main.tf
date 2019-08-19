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
        bucket = "${var.aws_bucket}"
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
  release_version               = "${var.release_version}"
  release_bucket_name           = "${var.release_bucket_name}"
  execution_role                = "${module.lambda-roles.execution_role}"
  application_ids               = "${data.terraform_remote_state.networks.application_ids}"
  security_group_ids            = "${module.security-group.lambda_into_vpc_id}"
}

module "lambda-roles" {
  source                    = "module-lambda-roles"
  project_name              = "${var.project_name}"
  env                       = "${var.env}"
  app_env_directory         = "${var.app_env_directory}"
}

module "security-group" {
  source                    = "module-security-group"
  vpc_id                    ="${var.vpc_id}"
  env                       = "${var.env}"
  service                   = "${var.service}"
}

module "cloud-watch" {
  source                        = "module-cloud-watch"
  project_name                  = "${var.project_name}"
  arn                           = "${module.lambda.arn}"
  env                           = "${var.env}"
}
