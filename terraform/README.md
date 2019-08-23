# Payment Reconciler - terraform 

## Introduction
This provisions the payment-reconciler Lambda function which is triggered by the cloud watch event.

## Config
The following config is required to deploy the project:
- release_bucket_name: `Optional` Release bucket for payment-reconciler, defaults to `release.ch.gov.uk`
- release_version: `Required` Github release version
- stage: `Required` deployment stage
- environment: `Required` to read the environment variable from.

## Cheat Sheet
tested on terraform version `0.11.7`, intended use with concourse. There should be no need for a manual deployment, instead add a specific environment to the concourse pipeline.

```sh
terraform get

source ./environments/development/dev/terraform.cfg

#initialize terraform
terraform init 

# Plan provisioning
TF_VAR_release_version=<release number> terraform plan -var-file=environments/development/<environment>/vars

# Apply
TF_VAR_release_version=<release number> terraform apply -var-file=environments/development/<environment>/vars
```

So far we only have development environments configuration. 