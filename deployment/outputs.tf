output "service_name" {
  value = google_cloud_run_v2_service.production.name
}

output "url" {
  value = google_cloud_run_v2_service.production.uri
}

output "bucket_name" {
  value = google_storage_bucket.bucket.name
}

output "bucket_url" {
  value = google_storage_bucket.bucket.url
}
