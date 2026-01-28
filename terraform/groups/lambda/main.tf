module "lambda" {
  source = "git@github.com:companieshouse/terraform-modules.git//aws/lambda?ref=1.0.342"

  environment           = var.environment
  function_name         = var.service
  lambda_runtime        = var.lambda_runtime
  lambda_handler        = var.handler

  lambda_code_s3_bucket = var.release_bucket_name
  lambda_code_s3_key    = var.release_artifact_key

  lambda_memory_size         = var.memory_megabytes
  lambda_timeout_seconds     = var.timeout_seconds
  lambda_logs_retention_days = var.lambda_logs_retention_days

  lambda_env_vars = local.lambda_env_vars

  lambda_cloudwatch_event_rules = [
    {
      name                = "payment-reconciler-${var.environment}"
      description         = "Call payment reconciler lambda"
      schedule_expression = var.cron_schedule
    }
  ]

  #  additional_policies = local.additional_iam_policies_json

  lambda_sg_egress_rule = {
    from_port   = -1
    to_port     = -1
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  lambda_vpc_access_subnet_ids = local.lambda_vpc_access_subnet_ids
  lambda_vpc_id                = data.aws_vpc.vpc.id
}



#provider "aws" {
#  region = var.aws_region
#}

#terraform {
#  backend "s3" {
#    encrypt = true
#  }
#}

#provider "vault" {
#  auth_login {
#    path = "auth/userpass/login/${var.vault_username}"
#    parameters = {
#      password = var.vault_password
#    }
#  }
#}

#data "terraform_remote_state" "network_remote_state" {
#  backend = "s3"
#  config = {
#    bucket = var.remote_state_bucket
#    key    = var.remote_state_key
#    region = var.aws_region
#  }
#}

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
