resource "aws_cloudwatch_event_rule" "payment_reconciler_daily" {
  name        = "${var.service}-${var.environment}"
  description = "Call payment Reconciler lambda"
  schedule_expression ="cron(0 23 * * ? *)"
}

resource "aws_cloudwatch_event_target" "call_payment_reconciler_lambda_everyday" {
    rule = "${aws_cloudwatch_event_rule.payment_reconciler_daily.name}"
    target_id = "${var.service}-${var.environment}"
    arn = "${var.arn}"
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_payment_reconciler" {
    statement_id = "AllowExecutionFromCloudWatch"
    action = "lambda:InvokeFunction"
    function_name = "${var.service}-${var.environment}"
    principal = "events.amazonaws.com"
    source_arn = "${aws_cloudwatch_event_rule.payment_reconciler_daily.arn}"
}