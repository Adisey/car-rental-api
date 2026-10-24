#!/usr/bin/env bash

# ---Before add test --start--
# set -a
# source .env.local
# set +a
# atlas "$@"
# ---Before add test --end--

ENV_FILE="${ENV_FILE:-.env.local}"
echo "Using env file: $ENV_FILE"
set -a
source "$ENV_FILE"
set +a

atlas "$@"