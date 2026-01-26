data "vault_generic_secret" "lambda_environment_variables" {
  path = "applications/${var.aws_profile}/${var.environment}/${var.service}/lambda_environment_variables"
}

# ------------------------------------------------------------------------------
# Lambdas
# ------------------------------------------------------------------------------
resource "aws_lambda_function" "payment_reconciler" {
  s3_bucket     = var.release_bucket_name
  s3_key        = "${var.service}/${var.service}-${var.release_version}.zip"
  function_name = "${var.service}-${var.environment}"
  role          = var.execution_role
  handler       = var.handler
  memory_size   = var.memory_megabytes
  timeout       = var.timeout_seconds
  runtime       = var.runtime

  vpc_config {
    subnet_ids         = var.subnet_ids
    security_group_ids = var.security_group_ids
  }
  environment {
    variables = merge (
      data.vault_generic_secret.lambda_environment_variables.data,
      var.open_lambda_environment_variables,
      { "SFTP_PORT"=58937 }
    )
  }
}

output "lambda_arn" {
  value = aws_lambda_function.payment_reconciler.arn
}
