provider "google" {
  project = var.gcp_project_id
  region  = var.cloudrun_location
}

terraform {
  backend "gcs" {
    bucket  = var.bucket_name
    prefix  = "terraform/state"
  }
  
}


resource "google_project_service" "service_usage" {
  project = var.gcp_project_id
  service = "serviceusage.googleapis.com"
}

resource "google_project_service" "cloud_run_api" {
  service = "run.googleapis.com"
}

resource "google_storage_bucket_iam_policy" "editor" {
  bucket      = google_storage_bucket.bucket.name
  policy_data = data.google_iam_policy.viewer.policy_data
}


resource "google_cloud_run_v2_service" "production" {
  name         = var.cloudrun_name
  location     = var.cloudrun_location
  launch_stage = "BETA"
  ingress      = "INGRESS_TRAFFIC_ALL"


  template {
    containers {
      image = var.image
      ports {
        container_port = 8080
      }
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

  depends_on = [
    google_project_service.cloud_run_api
  ]
}

resource "google_storage_bucket" "bucket" {
  project       = var.gcp_project_id
  name          = var.bucket_name
  location      = var.bucket_location
  storage_class = "STANDARD"

}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

data "google_iam_policy" "viewer" {
  binding {
    role = "roles/storage.objectViewer"
    members = [
      "allUsers",
    ]
  }
}
