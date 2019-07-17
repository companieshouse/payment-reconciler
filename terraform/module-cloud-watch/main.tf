resource "aws_cloudwatch_event_rule" "payment_reconciler_daily" {
  name        = "payment_reconciler_daily"
  description = "Call payment Reconciler lambda"
  schedule_expression ="cron(0 12 * * ? *)"
}

resource "aws_cloudwatch_event_target" "call_payment_reconciler_lambda_everyday" {
    rule = "${aws_cloudwatch_event_rule.payment_reconciler_daily.name}"
    target_id = "payment_reconciler"
    arn = "${var.arn}"
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_payment_reconciler" {
    statement_id = "AllowExecutionFromCloudWatch"
    action = "lambda:InvokeFunction"
    function_name = "${var.project_name}"
    principal = "events.amazonaws.com"
    source_arn = "${aws_cloudwatch_event_rule.payment_reconciler_daily.arn}"
}