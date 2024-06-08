provider "google" {
  project = "capstone-425808"
  region  = "asia-southeast2"
}

resource "google_cloud_run_v2_service" "production" {
  name     = "service-a"

  location = "asia-southeast1"

  template {
    containers {
      image = "asia-southeast2-docker.pkg.dev/capstone-425808/capstone/app:latest"
      volume_mounts {
        name       = "bucket"
        mount_path = "/models"
      }
    }

    volumes {
      name = "bucket"
      gcs {
        bucket    = google_storage_bucket.bucket.name
        read_only = false
      }
    }
  }
}

resource "google_storage_bucket" "bucket" {
    name     = "cloudrun-service"
    location = "asia-southeast2"
}