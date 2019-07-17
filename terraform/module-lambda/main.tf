# ------------------------------------------------------------------------------
# Lambdas
# ------------------------------------------------------------------------------
resource "aws_lambda_function" "payment_reconciler" {
  s3_bucket     = "${var.release_bucket_name}"
  s3_key        = "${var.project_name}/${var.project_name}-${var.release_version}.zip"
  function_name = "${var.project_name}"
  role          = "${var.execution_role}"
  handler       = "${var.handler}"
  memory_size   = "${var.memory_megabytes}"
  timeout       = "${var.timeout_seconds}"
  runtime       = "${var.runtime}"

  environment {
    variables = {
      S3_BUCKET_NAME = "${var.payment_reconciler_bucket}"
    }
  }
}

output "arn" {
  value = "${aws_lambda_function.payment_reconciler.arn}"
}