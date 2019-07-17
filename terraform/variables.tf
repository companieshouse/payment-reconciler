variable aws_region {
  default = "eu-west-1"
}

variable project_name {
  default = "payment-reconciler"
}

variable handler {
  default = "handler.handler"
}

variable memory_megabytes {
  default = "320"
}
variable release_bucket_name {
  default = "release.ch.gov.uk"
}
variable runtime {
  default = "go1.x"
}
variable timeout_seconds {
default = "6"
}

variable config_bucket_name {
  default = "config.ch.gov.uk"
}
variable env {}

variable app_env_directory {
  default = "env"
}
variable payment_reconciler_bucket {
  default = "payment_reconciler"
}

variable log_level {}

variable release_version {}