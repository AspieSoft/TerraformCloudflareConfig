#!/bin/bash

cd $(dirname "$0")

# reset terraform
rm -rf .terraform
rm -f .terraform.lock.hcl
rm -f .terraform.tfstate
rm -f terraform.tfstate
rm -f terraform.tfstate.backup

if ! [ "$(grep '"<Insert Zone ID>"' 'cloudflare.tf')" = "" ]; then
  read -p "Enter Zond ID: " zoneID
  sed -r -i "s/\"<Insert Zone ID>\"/\"$zoneID\"/" 'cloudflare.tf'
fi

terraform init
terraform plan
terraform apply

if ! [ "$zoneID" = "" ]; then
  sed -r -i "s/\"$zoneID\"/\"<Insert Zone ID>\"/" 'cloudflare.tf'

  # reset terraform
  rm -rf .terraform
  rm -f .terraform.lock.hcl
  rm -f .terraform.tfstate
  rm -f terraform.tfstate
  rm -f terraform.tfstate.backup
fi
