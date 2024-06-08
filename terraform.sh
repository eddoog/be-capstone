#! /bin/sh

cd deployment

terraform init

terraform validate

terraform plan