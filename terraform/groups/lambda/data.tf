data "vault_generic_secret" "stack_secrets" {
  path = "applications/${var.aws_profile}/${var.environment}/${local.stack_name}-stack"
}

data "vault_generic_secret" "service_secrets" {
  path = "applications/${var.aws_profile}/${var.environment}/${local.stack_name}-stack/${var.service}"
}

data "aws_vpc" "vpc" {
  filter {
    name   = "tag:Name"
    values = [local.vpc_name]
  }
}

data "aws_subnets" "application" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.vpc.id]
  }

  filter {
    name   = "tag:Name"
    values = [local.application_subnet_pattern]
  }
}

# todo - i don't think this policy is required any longer. commenting out with
#        a view to reintroducing permissions or deleting after testing in cidev
# data "aws_iam_policy_document" "payment-reconciler" {
#   statement {
#     effect = "Allow"
#
#     actions = [
#       "s3:PutAccountPublicAccessBlock",
#       "s3:GetAccountPublicAccessBlock",
#       "s3:ListAllMyBuckets",
#       "s3:HeadBucket",
#       "s3:GetObject",
#       "ec2:CreateNetworkInterface",
#       "ec2:DescribeNetworkInterfaces",
#       "ec2:DeleteNetworkInterface",
#       "logs:DescribeQueries",
#       "logs:GetLogRecord",
#       "logs:PutDestinationPolicy",
#       "logs:StopQuery",
#       "logs:TestMetricFilter",
#       "logs:DeleteDestination",
#       "logs:CreateLogGroup",
#       "logs:GetLogDelivery",
#       "logs:ListLogDeliveries",
#       "logs:CreateLogDelivery",
#       "logs:DeleteResourcePolicy",
#       "logs:PutResourcePolicy",
#       "logs:DescribeExportTasks",
#       "logs:GetQueryResults",
#       "logs:UpdateLogDelivery",
#       "logs:CancelExportTask",
#       "logs:DeleteLogDelivery",
#       "logs:PutDestination",
#       "logs:DescribeResourcePolicies",
#       "logs:DescribeDestinations"
#     ]
#
#     resources = [
#       "*"
#     ]
#   }
#
#   statement {
#     effect = "Allow"
#
#     actions = [
#       "s3:*",
#       "logs:*"
#     ]
#
#     resources = [
#       "arn:aws:logs:::log-group:/aws/lambda/${var.service}",
#       "arn:aws:logs:*:*:log-group:*:*:*",
#     ]
#   }
# }
