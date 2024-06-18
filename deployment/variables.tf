# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# You must define the following environment variables.
# ---------------------------------------------------------------------------------------------------------------------

# GOOGLE_CREDENTIALS
# or
# GOOGLE_APPLICATION_CREDENTIALS

variable "gcp_project_id" {
  description = "The ID of the GCP project in which these resources will be created."
  default     = "capstone-425808"
}

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "cloudrun_location" {
  description = "The Zone to launch the Cloud Instance into."
  type        = string
  default     = "asia-southeast2"
}

variable "cloudrun_name" {
  description = "The Name of the Cloud Run service to create."
  type        = string
  default     = "capstone"

}

variable "bucket_name" {
  description = "The Name of the example Bucket to create."
  type        = string
  default     = "bucket-capstone-425808"
}

variable "bucket_location" {
  description = "The location to store the Bucket. This value can be regional or multi-regional."
  type        = string
  default     = "asia-southeast2"
}

variable "image" {
  description = "The Docker image to deploy to Cloud Run."
  type        = string
  default     = "asia-southeast2-docker.pkg.dev/capstone-425808/capstone/capstone"
}
