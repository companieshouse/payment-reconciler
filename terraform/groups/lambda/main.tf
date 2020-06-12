provider "aws" {
  region = var.aws_region
}

terraform {
  backend "s3" {
    encrypt = true
  }
}

provider "vault" {
  auth_login {
    path = "auth/userpass/login/${var.vault_username}"
    parameters = {
      password = var.vault_password
    }
  }
}

data "terraform_remote_state" "network_remote_state" {
  backend = "s3"
  config = {
    bucket = var.remote_state_bucket
    key    = var.remote_state_key
    region = var.aws_region
  }
}

locals {
  test_and_development_vpc_id     = data.terraform_remote_state.network_remote_state.outputs.vpc_id
  test_and_development_subnet_ids = split(",", data.terraform_remote_state.network_remote_state.outputs.application_ids)
}

module "lambda" {
  source               = "./module-lambda"
  service              = var.service
  handler              = var.handler
  memory_megabytes     = var.memory_megabytes
  runtime              = var.runtime
  timeout_seconds      = var.timeout_seconds
  release_version      = var.release_version
  release_bucket_name  = var.release_bucket_name
  execution_role       = module.lambda-roles.execution_role
  open_lambda_environment_variables = var.open_lambda_environment_variables
  aws_profile          = var.aws_profile
  subnet_ids           = local.test_and_development_subnet_ids
  security_group_ids   = [module.security-group.lambda_into_vpc_id]
  environment          = var.environment
}

module "lambda-roles" {
  source      = "./module-lambda-roles"
  service     = var.service
  environment = var.environment
}

module "security-group" {
  source      = "./module-security-group"
  vpc_id      = local.test_and_development_vpc_id
  environment = var.environment
  service     = var.service
}

module "cloud-watch" {
  source        = "./module-cloud-watch"
  service       = var.service
  lambda_arn    = module.lambda.lambda_arn
  environment   = var.environment
  cron_schedule = var.cron_schedule
}
