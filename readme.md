# Terraform Cloudflare Config

[![donation link](https://img.shields.io/badge/buy%20me%20a%20coffee-square-blue)](https://buymeacoffee.aspiesoft.com)

A default config for new cloudflare domains.

I recommend forking this repo and changing the config to your liking.

Note: you will also need to install the [terraform](https://developer.hashicorp.com/terraform/downloads) cli.

## Installation

```shell script
git clone https://github.com/AspieSoft/TerraformCloudflareConfig.git
```

## Setup and Usage

1. Generate a cloudflare api token with access to your cloudflare config.
2. Set the key to an environment variable with the name `CLOUDFLARE_API_TOKEN` (in linux `echo 'export CLOUDFLARE_API_TOKEN="<Insert Cloudflare API Token>"' >> .zshrc`) or you can add it to a local file next to your cloudflare.tf file named `cloudflare_api_token.key`
3. run `TerraformCloudflareConfig/run`, you may be propted to enter the Zone ID for the cloudflare domain you want to change
