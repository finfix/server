#!/usr/bin/env bash

set -o pipefail -e

if [ -n "$1" ]; then
  PROJECT_NAME="$1"
  SHARE="../$1"
  SCRIPTS="../ci-scripts"
else
  PROJECT_NAME="$CI_PROJECT_NAME"
  SHARE="$SHARE"
  SCRIPTS="$SCRIPTS"
fi

if [ -n "$2" ]; then
  SERVICE_PORT="$2"
else
  SERVICE_PORT="$SERVICE_PORT"
fi

sed \
-e "s|@{project_name}|$PROJECT_NAME|" \
-e "s|@{service_port}|$SERVICE_PORT|" \
< "$SCRIPTS/templates/Dockerfile" > "$SHARE/Dockerfile"
