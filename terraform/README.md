# Payment Reconciler - terraform 

## Introduction
* This provisions the `payment-reconciler` Lambda function, together with a cloud watch event to trigger it.
* Use [the Concourse pipeline](https://ci.platform.aws.chdev.org/teams/development/pipelines/payment-reconciler) to
provision it in an automated manner. That pipeline uses the terraform configurations found in this directory.  

## Configuration

The following configuration properties are required to deploy the reconciler:

- environment: `Required` The environment to be deployed to 
- aws_profile: `Required` The AWS profile
- cron_schedule: `Required` A cron schedule for the triggering of the lambda
- aws_region: `Required` The AWS region
- release_bucket_name `Required` The name of the release bucket containing the distribution 

Depending on the target deployment environment, it may be necessary to override the default values for these properties 
too:

- remote_state_bucket: `Optional` The S3 bucket where the terraform state for the reconciler is stored
- remote_state_key: `Optional` The location in the S3 bucket where terraform state shared by all deployed artefacts is 
stored

In addition to these, for a manual terraform deployment, use 
[platform-tools](https://github.com/companieshouse/platform-tools). The following arguments must be 
provided on the command line as in the example below:

```
terraform-runner -g lambda -c plan -p development-eu-west-2 -e cidev
```

The following inputs will be required for terraform too:

- var.release_version: `Required` The Github release version to be deployed
- var.vault_password: `Required` The password required for access to the terraform vault for other configuration parameters
- var.vault_username: `Required` The user required for access to the terraform vault
- provider.vault.address: `Required` The address of the terraform vault

## Further background

* Further background regarding the deployment of this Lambda function can be found 
[here](https://companieshouse.atlassian.net/wiki/spaces/CHS/pages/1650000125/GCI-1042+Deployment+of+payment-reconciler).
