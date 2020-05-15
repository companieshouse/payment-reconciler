# Payment Reconciler - terraform 

## Introduction
* This provisions the `payment-reconciler` Lambda function, together with a cloud watch event to trigger it.
* Use the Concourse pipeline to provision it in an automated manner. That pipeline uses the terraform configurations 
found in this directory.  

## Configuration

To deploy the reconciler, we must specify the following configuration properties:

- environment: `Required` The environment to be deployed to 
- aws_profile: `Required` The AWS profile
- cron_schedule: `Required` A cron schedule for the triggering of the lambda
- aws_region: `Required` The AWS region
- release_bucket_name `Required` The name of the release bucket containing the distribution 
- remote_state_bucket: `Required` The S3 bucket where the terraform state for the network is stored
- remote_state_key: `Required` The location in the S3 bucket where the terraform state of the network components is read
 from

## Manual deployment

In addition to the above properties, for a manual terraform deployment, use 
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

### Environment variables to reduce repetition

To avoid having to enter the same values over and over again when testing the provisioning of the reconciler, use the 
following environment variables: 

| Variable name | Equivalent to | Example |
| ------------- | ---------------- | ------- |
| TF_VAR_vault_password | var.vault_password | export TF_VAR_vault_password=&lt;PASSWORD&gt; |
| TF_VAR_vault_username | var.vault_username | export TF_VAR_vault_username=&lt;USER NAME&gt; |
| VAULT_ADDR    | provider.vault.address | export VAULT_ADDR=&lt;URL&gt; |

## Further background

* Further background regarding the deployment of this Lambda function can be found 
[here](https://companieshouse.atlassian.net/wiki/spaces/CHS/pages/1650000125/GCI-1042+Deployment+of+payment-reconciler).
