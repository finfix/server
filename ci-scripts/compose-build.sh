#!/usr/bin/env bash

set -o pipefail -e

# Передается название сервиса, остальные значения захардкожены
if [ -n "$1" ]; then
  PROJECT_NAME="$1"
  SHARE="../$1"
  if [ -n "$2" ]; then
    IMAGE_NAME="registry.gitlab.com/mycoin/server/$1"
  else
    IMAGE_NAME="$1"
  fi
  SCRIPTS="../ci-scripts"
  ENV_FILE="../../env"
  NETWORK_MODE="network_mode: host"
  SERVICE_PORT="8000"
else
#  Если не передается название сервиса, то берем значения из переменных окружения гита
  PROJECT_NAME="$CI_PROJECT_NAME"
  SHARE="$SHARE"
  IMAGE_NAME="$CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME-$CI_PIPELINE_IID"
  SCRIPTS="$SCRIPTS"
  ENV_FILE=".env"
  NETWORK_MODE=""
  SERVICE_PORT="$SERVICE_PORT"
fi

sed \
-e "s|@{project_name}|$PROJECT_NAME|" \
-e "s|@{project_name}|$PROJECT_NAME|" \
-e "s|@{service_port}|$SERVICE_PORT|" \
-e "s|@{service_port}|$SERVICE_PORT|" \
-e "s|@{image_name}|$IMAGE_NAME|" \
-e "s|@{env_file}|$ENV_FILE|" \
-e "s|@{network_mode}|$NETWORK_MODE|" \
< "$SCRIPTS/templates/docker-compose.yml" > "$SHARE/docker-compose.yml"
