variable handler {
  type        = string
}
variable memory_megabytes {
  type        = string
}
variable release_bucket_name {
  type        = string
}
variable runtime {
  type        = string
}
variable timeout_seconds {
  type        = string
}
variable service {
  type        = string
}
variable release_version {
  type        = string
}
variable execution_role {
  type        = string
  description = "IAM role from lambda-roles module used for executing the Lambda."
}
variable subnet_ids {
  type        = list(string)
  description = "A list of subnet IDs associated with the Lambda in its VPC configuration."
}
variable security_group_ids {
  type        = list(string)
  description = "A list of security group IDs associated with the Lambda in its VPC configuration."
}
variable environment {
  type        = string
}
variable aws_profile {
  type        = string
}
