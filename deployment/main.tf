resource "google_cloud_run_v2_service" "default" {
  name     = "cloudrun-service"

  location     = "asia-southeast2"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
      volume_mounts {
        name       = "bucket"
        mount_path = "/var/www"
      }
    }

    volumes {
      name = "bucket"
      gcs {
        bucket    = google_storage_bucket.default.name
        read_only = false
      }
    }
  }
}

resource "google_storage_bucket" "default" {
    name     = "cloudrun-service"
    location = "US"
}