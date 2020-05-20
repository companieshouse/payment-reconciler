# AWS
variable "aws_account_id" {
  type        = string
  description = "The AWS account ID."
}

# IAM
variable "iam_role_policies" {
  type        = map
  description = "A map of IAM policies and the role to attach them to."
}
variable "policy_description" {
  type        = string
  description = "A description of TF source to add to each policy created."
  default     = "Provisioned with payment-reconciler/groups/provisioning-role"
}