#!/usr/bin/env bash

set -e

echo "Generating Go models..."

oapi-codegen \
  -generate types \
  -package api_models \
  -o internal/api_models/types.gen.go \
  openapi/api.yaml

echo "Done."



