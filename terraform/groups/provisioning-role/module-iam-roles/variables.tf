# AWS
variable "aws_account_id" {
  type        = string
  description = "The AWS account ID."
}

# IAM
variable "iam_roles" {
  type        = map
  description = "A map of IAM roles to be created."
}
variable "iam_role_policies" {
  type        = map
  description = "A map of IAM policies and the role to attach them to."
}
variable "iam_policies_arns" {
  type        = map
  description = "A map of IAM policies and their ARNs."
}
variable "iam_role_aws_policies" {
  type        = map
  description = "A map of AWS Managed policies and the role to attach them to."
}
