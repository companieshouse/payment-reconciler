provider "aws" {
  region  = var.region
  version = "~> 2.0"
}

terraform {
  backend "s3" {
  }
}

data "aws_caller_identity" "current" {}

module "iam-policies" {
   source                = "./module-iam-policies"
   aws_account_id        = data.aws_caller_identity.current.account_id
   iam_role_policies     = var.iam_role_policies
 }

 module "iam-roles" {
   source                = "./module-iam-roles"
   aws_account_id        = data.aws_caller_identity.current.account_id
   iam_policies_arns     = module.iam-policies.iam_policies_arns
   iam_roles             = var.iam_roles
   iam_role_policies     = var.iam_role_policies
   iam_role_aws_policies = var.iam_role_aws_policies
 }
