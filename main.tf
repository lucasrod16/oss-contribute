terraform {
  backend "gcs" {
    bucket = "lucasrod16-tfstate"
    prefix = "osscontribute"
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

resource "google_cloud_run_v2_service" "oss_contribute" {
  name                = "oss-contribute"
  location            = "us-central1"
  deletion_protection = false

  template {
    containers {
      image = "lucasrod96/oss-contribute:${var.image_digest}"
    }
  }
}

resource "google_cloud_run_v2_service_iam_member" "noauth" {
  name     = google_cloud_run_v2_service.oss_contribute.name
  location = google_cloud_run_v2_service.oss_contribute.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
