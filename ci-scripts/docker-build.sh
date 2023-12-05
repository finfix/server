#!/usr/bin/env sh

# Собираем образ
docker build -t $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME-$CI_PIPELINE_IID -f $SHARE/Dockerfile .
# Логинимся в реестре образов
docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
# Пушим образ в реестр
docker push -a $CI_REGISTRY_IMAGE
