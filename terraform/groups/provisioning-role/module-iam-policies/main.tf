# Create policies
resource "aws_iam_policy" "iam_policy_plan" {
  name        = "payment-reconciler-terraform-plan"
  description = var.policy_description
  policy      = data.aws_iam_policy_document.plan.json
}

resource "aws_iam_policy" "iam_policy_apply" {
  name        = "payment-reconciler-terraform-apply"
  description = var.policy_description
  policy      = data.aws_iam_policy_document.apply.json
}

resource "aws_iam_policy" "iam_policy_destroy" {
  name        = "payment-reconciler-terraform-destroy"
  description = var.policy_description
  policy      = data.aws_iam_policy_document.destroy.json
}
