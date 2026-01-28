locals {

#  additional_iam_policies_json = [data.aws_iam_policy_document.penalty_payment_error_reporter_execution.json]
  stack_name                   = "company-requests"
  lambda_env_vars              = merge(local.service_secrets, var.open_lambda_environment_variables)
  lambda_vpc_access_subnet_ids = data.aws_subnets.application.ids
  application_subnet_pattern   = local.stack_secrets["application_subnet_pattern"]
  vpc_name                     = local.stack_secrets["vpc_name"]

  stack_secrets   = data.vault_generic_secret.stack_secrets.data
  service_secrets = data.vault_generic_secret.service_secrets.data
}
