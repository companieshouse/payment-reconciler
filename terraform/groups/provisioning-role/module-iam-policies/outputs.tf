# Outputs
output "iam_policies_arns" {
  value = {
    "payment-reconciler-terraform-plan"    = aws_iam_policy.iam_policy_plan.arn
    "payment-reconciler-terraform-apply"   = aws_iam_policy.iam_policy_apply.arn
    "payment-reconciler-terraform-destroy" = aws_iam_policy.iam_policy_destroy.arn
  }
}
