terraform {
  required_version = ">= 1.3.0, < 1.4.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.72.0, < 6.0"
    }
    vault = {
      source  = "hashicorp/vault"
      version = ">= 3.18.0, < 5.0"
    }
  }

  backend "s3" {
    encrypt = true
  }
}

provider "aws" {
  region = var.aws_region
}
