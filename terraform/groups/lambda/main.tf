provider "aws" {
  region = var.aws_region
}

terraform {
  backend "s3" {
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
  aws_profile          = var.aws_profile
  subnet_ids           = var.subnet_ids
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
  vpc_id      = var.vpc_id
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
