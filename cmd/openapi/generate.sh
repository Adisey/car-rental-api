#!/usr/bin/env bash

set -e

$(go env GOPATH)/bin/oapi-codegen \
  -generate types \
  -package api_models \
  -o internal/api_models/types.gen.go \
  openapi/api.yaml



