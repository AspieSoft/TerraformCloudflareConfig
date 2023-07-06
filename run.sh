#!/bin/bash

# reset terraform
rm -rf .terraform
rm -f .terraform.lock.hcl
rm -f .terraform.tfstate
rm -f terraform.tfstate
rm -f terraform.tfstate.backup

terraform init
terraform plan
terraform apply

# reset terraform
rm -rf .terraform
rm -f .terraform.lock.hcl
rm -f .terraform.tfstate
rm -f terraform.tfstate
rm -f terraform.tfstate.backup
