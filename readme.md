# Terraform Cloudflare Config

[![donation link](https://img.shields.io/badge/buy%20me%20a%20coffee-paypal-blue)](https://paypal.me/shaynejrtaylor?country.x=US&locale.x=en_US)

A default config for new cloudflare domains.

I recommend forking this repo and changing the config to your liking.

Note: you will also need to install the [terraform](https://developer.hashicorp.com/terraform/downloads) cli.
This does not apply to windows users. The terraform windows exe is included to simplify the install.

## Installation

```shell script
git clone https://github.com/AspieSoft/TerraformCloudflareConfig.git
```

## Setup and Usage

1. Generate a cloudflare api token with access to your cloudflare config. You will need to include the following zone edit permissions:
   - Config Rules
   - Cache Rules
   - Origin Rules
   - Zone Settings
   - SSL and Certificates
2. Set the key to an environment variable with the name `CLOUDFLARE_API_TOKEN` (in linux `echo 'export CLOUDFLARE_API_TOKEN="<Insert Cloudflare API Token>"' >> .zshrc`) or you can add it to a local file next to your cloudflare.tf file named `cloudflare_api_token.key`.
3. run `TerraformCloudflareConfig/run`, you may be prompted to enter the Zone ID for the cloudflare domain you want to change. *(Note: add the flag `-y` to autoyes confirmation prompts when possible)*

## Optional

Add a file named `zone.list` and add multiple zone id's *(one per line)* to automatically configure multiple domains with the same settings. *(Note: this will enable the autoyes feature)*
