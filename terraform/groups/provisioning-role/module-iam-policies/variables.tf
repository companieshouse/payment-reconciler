# AWS
variable "aws_account_id" {
  type        = string
  description = "The AWS account ID."
}

# IAM
variable "iam_role_policies" {
  type        = map(string)
  description = "A map of IAM policies and the role to attach them to."
}
variable "policy_description" {
  type        = string
  description = "A description of TF source to add to each policy created."
  default     = "Provisioned with payment-reconciler/groups/provisioning-role"
}
variable "iam_policy_region" {
  type        = string
  description = "The AWS region to configure in IAM policies."
}
