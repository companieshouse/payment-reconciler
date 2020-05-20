data "aws_iam_policy_document" "destroy" {
  statement {
    sid       = "DestroyPolicyListedResourcesWithConditions"
    effect    = "Allow"
    resources = [
      "arn:aws:ec2:eu-west-2:${var.aws_account_id}:security-group/*"
    ]
    actions   = [
      "ec2:DeleteSecurityGroup"
    ]
    condition {
      test     = "ForAllValues:StringLike"
      variable = "ec2:ResourceTag/Name"
      values   = [
        "*-payment-reconciler-sg"
      ]
    }
  }

  statement {
    sid       = "DestroyPolicyAllResources"
    effect    = "Allow"
    resources = ["*"]
    actions   = [
      "ec2:DescribeNetworkInterfaces",
      "ec2:DeleteNetworkInterface",
    ]
  }

  statement {
    sid       = "DestroyPolicyListedResources"
    effect    = "Allow"
    resources = [
      "arn:aws:events:eu-west-2:${var.aws_account_id}:rule/payment-reconciler-*",
      "arn:aws:iam::${var.aws_account_id}:role/payment-reconciler-execution-*",
      "arn:aws:lambda:eu-west-2:${var.aws_account_id}:function:payment-reconciler-*",
    ]
    actions = [
      "events:DeleteRule",
      "events:RemoveTargets",
      "iam:ListInstanceProfilesForRole",
      "iam:DeleteRolePolicy",
      "iam:DeleteRole",
      "lambda:DeleteFunction",
    ]
  }
}
