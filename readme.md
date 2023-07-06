# Terraform Cloudflare Config

[![donation link](https://img.shields.io/badge/buy%20me%20a%20coffee-square-blue)](https://buymeacoffee.aspiesoft.com)

A default config for new cloudflare domains.

I recommend forking this repo and changing the config to your liking.

Note: you will also need to install the [terraform](https://developer.hashicorp.com/terraform/downloads) cli for linux.

> Notice: run.sh assumes you are using linux.

## Installation

```shell script
git clone https://github.com/AspieSoft/TerraformCloudflareConfig.git
```

## Setup

Add your `api_token` and `zone_id` to cloudflare.tf, then run `TerraformCloudflareConfig/run.sh` to update the config for that domain.
