#!/bin/bash

# reset terraform
rm -rf .terraform
rm -f .terraform.lock.hcl
rm -f .terraform.tfstate

terraform init
terraform plan
terraform apply
