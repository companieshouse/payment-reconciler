# ------------------------------------------------------------------------------
# Lambdas
# ------------------------------------------------------------------------------
resource "aws_lambda_function" "payment_reconciler" {
  s3_bucket     = "${var.release_bucket_name}"
  s3_key        = "${var.service}/${var.service}-${var.release_version}.zip"
  function_name = "${var.service}-${var.environment}"
  role          = "${var.execution_role}"
  handler       = "${var.handler}"
  memory_size   = "${var.memory_megabytes}"
  timeout       = "${var.timeout_seconds}"
  runtime       = "${var.runtime}"
  
  vpc_config {
    subnet_ids         = ["${split(",", var.application_ids)}"]
    security_group_ids = ["${list(var.security_group_ids)}"]
  }
}

output "arn" {
  value = "${aws_lambda_function.payment_reconciler.arn}"
}