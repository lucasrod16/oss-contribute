terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
  backend "gcs" {
    bucket = "lucasrod16-tfstate"
    prefix = "ossprojects"
  }
}

provider "google" {
  project = "groovy-momentum-434802-g9"
  region  = "us-central1"
}

variable "image_digest" {
  description = "The digest of the container image"
  type        = string
}

resource "google_cloud_run_v2_service" "oss_projects" {
  name                = "oss-projects"
  location            = "us-central1"
  deletion_protection = false

  template {
    containers {
      image = "lucasrod96/oss-projects@${var.image_digest}"
    }
  }
}

resource "google_cloud_run_v2_service_iam_member" "noauth" {
  name     = google_cloud_run_v2_service.oss_projects.name
  location = google_cloud_run_v2_service.oss_projects.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

# Network Endpoint Group for Cloud Run
resource "google_compute_region_network_endpoint_group" "oss_projects_neg" {
  name                  = "oss-projects-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  cloud_run {
    service = google_cloud_run_v2_service.oss_projects.name
  }
}

# Backend Service
resource "google_compute_backend_service" "oss_projects_backend" {
  name                            = "oss-projects-backend"
  load_balancing_scheme           = "EXTERNAL_MANAGED"
  connection_draining_timeout_sec = 0
  backend {
    group = google_compute_region_network_endpoint_group.oss_projects_neg.id
  }
}

# SSL Certificate
resource "google_compute_managed_ssl_certificate" "osscontribute_cert" {
  name = "custom-domains-osscontribute-com-49a4-cert"
  managed {
    domains = ["osscontribute.com"]
  }
}

# Static IP Address
resource "google_compute_global_address" "custom_domains_ip" {
  name = "custom-domains-49a4-ip"
}

# Load Balancer (HTTPS and HTTP redirect)
resource "google_compute_global_forwarding_rule" "https" {
  name                  = "custom-domains-49a4-fe"
  target                = google_compute_target_https_proxy.default.id
  port_range            = "443"
  ip_address            = google_compute_global_address.custom_domains_ip.id
  load_balancing_scheme = "EXTERNAL_MANAGED"
}

resource "google_compute_global_forwarding_rule" "http" {
  name                  = "custom-domains-49a4-fe-http"
  target                = google_compute_target_http_proxy.default.id
  port_range            = "80"
  ip_address            = google_compute_global_address.custom_domains_ip.id
  load_balancing_scheme = "EXTERNAL_MANAGED"
}

resource "google_compute_target_https_proxy" "default" {
  name             = "custom-domains-49a4-proxy"
  url_map          = google_compute_url_map.default.id
  ssl_certificates = [google_compute_managed_ssl_certificate.osscontribute_cert.id]
}

resource "google_compute_target_http_proxy" "default" {
  name    = "custom-domains-49a4-proxy-http"
  url_map = google_compute_url_map.http_redirect.id
}

resource "google_compute_url_map" "default" {
  name = "custom-domains-49a4"

  default_url_redirect {
    strip_query = true
  }

  host_rule {
    hosts        = ["osscontribute.com"]
    path_matcher = "osscontribute-com"
  }

  path_matcher {
    name            = "osscontribute-com"
    default_service = google_compute_backend_service.oss_projects_backend.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.oss_projects_backend.id
    }
  }
}

resource "google_compute_url_map" "http_redirect" {
  name = "custom-domains-49a4-http"
  default_url_redirect {
    https_redirect         = true
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
    strip_query            = false
  }
}

output "load_balancer_ip" {
  value = google_compute_global_address.custom_domains_ip.address
}
