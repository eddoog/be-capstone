provider "google" {
  project = var.gcp_project_id
  region  = var.cloudrun_location
}


resource "google_cloud_run_service_iam_policy" "noauth" {
  location = var.cloudrun_location
  project  = var.gcp_project_id
  service  = var.cloudrun_name

  policy_data = data.google_iam_policy.noauth.policy_data
}

resource "google_storage_bucket_iam_policy" "editor" {
  bucket = google_storage_bucket.bucket.name
  policy_data = data.google_iam_policy.viewer.policy_data
}


resource "google_cloud_run_v2_service" "production" {
  name     = var.cloudrun_name

  location = var.cloudrun_location

  template {
    containers {
      image = var.image
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
    project = var.gcp_project_id
    name     = var.bucket_name
    location = var.bucket_location
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
