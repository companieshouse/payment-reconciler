variable aws_region {
  default = "eu-west-1"
}

variable project_name {
  default = "payment-reconciler"
}

variable handler {
  default = "main"
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

variable env {}

variable app_env_directory {
  default = "env"
}

variable release_version {}

variable state_prefix {}
variable farm {}
variable aws_bucket {}