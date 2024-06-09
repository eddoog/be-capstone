#! /bin/sh

cd deployment

terraform init

terraform fmt

terraform validate

terraform plan

# terraform apply -auto-approve