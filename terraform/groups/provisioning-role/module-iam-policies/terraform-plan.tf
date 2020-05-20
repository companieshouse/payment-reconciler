data "aws_iam_policy_document" "plan" {
  statement {
    sid       = "PlanPolicyAllResources"
    effect    = "Allow"
    resources = ["*"]
    actions   = [
      "ec2:DescribeAccountAttributes",
      "sts:GetCallerIdentity",
    ]
  }

  statement {
    sid       = "PlanPolicyListedResources"
    effect    = "Allow"
    resources = [
      "arn:aws:events:eu-west-2:${var.aws_account_id}:rule/payment-reconciler-cidev",
      "arn:aws:iam::${var.aws_account_id}:role/payment-reconciler-execution-cidev",
    ]
    actions = [
      "events:DescribeRule",
      "events:ListTagsForResource",
      "iam:GetRole",
      "iam:GetRolePolicy",
    ]
  }
}
