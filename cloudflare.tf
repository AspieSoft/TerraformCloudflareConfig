terraform {
  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "~> 4.0"
    }
  }
}

provider "cloudflare" {
  # api_token = "<Insert Cloudflare API Token>"
}

variable "zone_id" {
  default = "<Insert Zone ID>"
}

resource "cloudflare_zone_settings_override" "settings" {
  zone_id = var.zone_id

  settings {
    tls_1_3 = "on"
    automatic_https_rewrites = "on"
    ssl = "strict"
    # waf = "on"

    always_online = "off"
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
