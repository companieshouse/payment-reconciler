data "aws_iam_policy_document" "apply" {
  statement {
    sid       = "ApplyPolicyAllResources"
    effect    = "Allow"
    resources = ["*"]
    actions = [
      "ec2:CreateSecurityGroup",
      "ec2:CreateTags",
      "lambda:CreateFunction",
    ]
  }

  statement {
    sid    = "ApplyPolicyListedResources"
    effect = "Allow"
    resources = [
      "arn:aws:iam::${var.aws_account_id}:role/payment-reconciler-execution-*",
      "arn:aws:events:eu-west-2:${var.aws_account_id}:rule/payment-reconciler-*",
      "arn:aws:ec2:eu-west-2:${var.aws_account_id}:security-group/*",
      "arn:aws:lambda:eu-west-2:${var.aws_account_id}:function:payment-reconciler-devops1",
    ]
    actions = [
      "events:PutTargets",
      "events:EnableRule",
      "events:PutRule",
      "ec2:AuthorizeSecurityGroupEgress",
      "ec2:RevokeSecurityGroupEgress",
      "iam:PassRole",
      "iam:CreateRole",
      "iam:PutRolePolicy",
      "lambda:AddPermission",
    ]
  }
}
