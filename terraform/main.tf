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
  service                       = "${var.service}"
  handler                       = "${var.handler}"
  memory_megabytes              = "${var.memory_megabytes}"
  runtime                       = "${var.runtime}"
  timeout_seconds               = "${var.timeout_seconds}" 
  release_version               = "${var.release_version}"
  release_bucket_name           = "${var.release_bucket_name}"
  execution_role                = "${module.lambda-roles.execution_role}"
  application_ids               = "${data.terraform_remote_state.networks.data_ids}"
  security_group_ids            = "${module.security-group.lambda_into_vpc_id}"
  environment                   = "${var.environment}"
}

module "lambda-roles" {
  source                    = "module-lambda-roles"
  service                   = "${var.service}"
  environment               = "${var.environment}"
  app_env_directory         = "${var.app_env_directory}"
}

module "security-group" {
  source                    = "module-security-group"
  vpc_id                    ="${var.vpc_id}"
  environment               = "${var.environment}"
  service                   = "${var.service}"
}

module "cloud-watch" {
  source                        = "module-cloud-watch"
  service                       = "${var.service}"
  arn                           = "${module.lambda.arn}"
  environment                   = "${var.environment}"
}
