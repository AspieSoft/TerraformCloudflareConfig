terraform {
  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "~> 4.0"
    }
  }
}

provider "cloudflare" {
  # api_token = "export CLOUDFLARE_API_TOKEN=\"<Insert API Token>\" >> .zshrc"
}

variable "zone_id" {
  default = "<Insert Zone ID>"
}

variable "domain" {
  default = "aspiemail.com"
}

resource "cloudflare_zone_settings_override" "settings" {
  # name = var.domain
  zone_id = var.zone_id

  settings {
    tls_1_3 = "on"
    automatic_https_rewrites = "on"
    ssl = "full"
    # waf = "on"

    always_online = "on"
    always_use_https = "on"
    brotli = "on"
    browser_cache_ttl = 43200
    browser_check = "on"
    cache_level = "aggressive"
    challenge_ttl = 1800
    early_hints = "on"
    email_obfuscation = "on"
    # http2 = "on"
    http3 = "on"
    hotlink_protection = "on"
    ipv6 = "on"

    opportunistic_encryption = "on"
    opportunistic_onion = "on"
    tls_client_auth = "on"
    universal_ssl = "on"
    privacy_pass = "on"
    rocket_loader = "on"
    # visitor_ip = "on"
    # webp = "on"
    websockets = "on"

    minify {
      css = "on"
      html = "on"
      js = "on"
    }

    security_header {
      enabled = true
      include_subdomains = true
      max_age = 15552000
      preload = true
      nosniff = false
    }
  }
}
