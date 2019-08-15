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
  vpc_config {
    subnet_ids         = ["${var.subnet_ids}"]
    security_group_ids = ["${list(var.security_group_ids)}"]
  }
}

output "arn" {
  value = "${aws_lambda_function.payment_reconciler.arn}"
}