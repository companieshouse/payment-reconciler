# ------------------------------------------------------------------------------
# Lambdas
# ------------------------------------------------------------------------------
resource "aws_lambda_function" "payment_reconciler" {
  s3_bucket     = "${var.release_bucket_name}"
  s3_key        = "${var.workspace_key_prefix}/${var.service}-${var.version}.zip"
  function_name = "${var.service}-${var.environment}"
  role          = "${var.execution_role}"
  handler       = "${var.handler}"
  memory_size   = "${var.memory_megabytes}"
  timeout       = "${var.timeout_seconds}"
  runtime       = "${var.runtime}"
  
  vpc_config {
    subnet_ids         = ["${split(",", var.subnet_ids)}"]
    security_group_ids = ["${list(var.security_group_ids)}"]
  }
}

output "lambda_arn" {
  value = "${aws_lambda_function.payment_reconciler.arn}"
}