variable aws_region {
  default = "eu-west-1"
}

variable handler {
  default = "payment-reconciler"
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

variable environment {}

variable app_env_directory {
  default = "env"
}

variable release_version {}


variable aws_bucket {}

variable service {
  default = "payment-reconciler"
}

variable vpc_id {
  
}
variable application_ids {
  
}
