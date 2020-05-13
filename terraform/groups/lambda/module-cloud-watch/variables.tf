variable service {
  type        = string
}

variable lambda_arn {
  type        = string
  description = "The Lambda ARN to configure as the target of the CloudWatch event"
}

variable environment {
  type        = string
}

variable cron_schedule {
  type        = string
}
