# ------------------------------------------------------------------------------
# Policy Documents
# ------------------------------------------------------------------------------
data "aws_iam_policy_document" "payment_reconciler_trust" {
  statement {
    effect = "Allow"

    actions = [
      "sts:AssumeRole",
    ]

    principals {
      type = "Service"

      identifiers = [
        "lambda.amazonaws.com",
      ]
    }
  }
}

data "aws_iam_policy_document" "payment_reconciler_execution" {
  statement {
    effect = "Allow"

    actions = [
      "logs:DescribeQueries",
      "logs:GetLogRecord",
      "s3:PutAccountPublicAccessBlock",
      "logs:PutDestinationPolicy",
      "logs:StopQuery",
      "logs:TestMetricFilter",
      "logs:DeleteDestination",
      "logs:CreateLogGroup",
      "logs:GetLogDelivery",
      "logs:ListLogDeliveries",
      "logs:CreateLogDelivery",
      "logs:DeleteResourcePolicy",
      "logs:PutResourcePolicy",
      "logs:DescribeExportTasks",
      "s3:GetAccountPublicAccessBlock",
      "logs:GetQueryResults",
      "s3:ListAllMyBuckets",
      "logs:UpdateLogDelivery",
      "logs:CancelExportTask",
      "logs:DeleteLogDelivery",
      "s3:HeadBucket",
      "s3:GetObject",
      "logs:PutDestination",
      "logs:DescribeResourcePolicies",
      "logs:DescribeDestinations"
    ]

    resources = [
      "*"
    ]
  }

  statement {
    effect = "Allow"

    actions = [
      "s3:*",
      "logs:*"
    ]

    resources = [
      "arn:aws:logs:::log-group:/aws/lambda/${var.project_name}",
      "arn:aws:logs:*:*:log-group:*:*:*",
      "arn:aws:s3:::${var.payment_reconciler_bucket}/*",
      "arn:aws:s3:::${var.config_bucket_name}/${var.env}/${var.project_name}/${var.app_env_directory}",
    ]
  }
}

# ------------------------------------------------------------------------------
# Roles
# ------------------------------------------------------------------------------
resource "aws_iam_role" "payment_reconciler_execution" {
  name               = "payment-reconciler-execution"
  assume_role_policy = "${data.aws_iam_policy_document.payment_reconciler_trust.json}"
}

# ------------------------------------------------------------------------------
# Role Policies
# ------------------------------------------------------------------------------
resource "aws_iam_role_policy" "payment_reconciler_execution" {
  name   = "payment_reconciler_execution"
  role   = "${aws_iam_role.payment_reconciler_execution.id}"
  policy = "${data.aws_iam_policy_document.payment_reconciler_execution.json}"
}

output "execution_role" {
  value = "${aws_iam_role.payment_reconciler_execution.arn}"
}