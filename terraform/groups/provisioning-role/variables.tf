# AWS
variable "region" {
  type        = string
  description = "The AWS region."
}

# IAM
variable "iam_roles" {
  type        = map(string)
  description = "A map of IAM roles to be created."
}
variable "iam_role_policies" {
  type        = map(string)
  description = "A map of IAM policies to be created and the role to attach them to."
}
variable "iam_role_aws_policies" {
  type        = map(string)
  description = "A map of AWS Managed policies and the role to attach them to."
}
variable "iam_policy_region" {
  type        = string
  description = "The AWS region to configure in IAM policies."
}
