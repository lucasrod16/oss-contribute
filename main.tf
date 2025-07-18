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
