#! /bin/sh

cd deployment

terraform init -backend-config="bucket=bucket-capstone-425808"

terraform fmt

terraform validate

terraform plan -lock=false

# terraform apply -auto-approve -lock=false