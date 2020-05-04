resource "aws_cloudwatch_event_rule" "payment_reconciler" {
  name        = "${var.service}-${var.environment}"
  description = "Call payment Reconciler lambda"
  schedule_expression ="${var.cron_schedule}"
}

resource "aws_cloudwatch_event_target" "call_payment_reconciler_lambda" {
    rule = "${aws_cloudwatch_event_rule.payment_reconciler.name}"
    target_id = "${var.service}-${var.environment}"
    arn = "${var.lambda_arn}"
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_payment_reconciler" {
    statement_id = "AllowExecutionFromCloudWatch"
    action = "lambda:InvokeFunction"
    function_name = "${var.service}-${var.environment}"
    principal = "events.amazonaws.com"
    source_arn = "${aws_cloudwatch_event_rule.payment_reconciler.arn}"
    depends_on = ["${function_name}"]
}