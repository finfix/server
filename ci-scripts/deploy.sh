#!/usr/bin/env bash

# Делаем, чтобы любая ошибка в скрипте приводила к ошибке в pipeline
set -o pipefail -e
# Проверяем существование нужной нам директории
if ssh $USER@$HOST -p $PORT [ ! -d /opt/coin/server/auth ]; then
    ssh $USER@$HOST -p $PORT "mkdir -p /opt/coin/server/auth"
fi
# Копируем docker-compose в папку конкретного проекта на сервере
scp -P $PORT -p $SHARE/docker-compose.yml  $USER@$HOST:/opt/coin/server/$CI_PROJECT_NAME/docker-compose.yml
# из сервера логинимся в реестре образов
ssh $USER@$HOST -p $PORT "docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY"
# на сервере останавливаем контейнер, удаляем его, скачиваем новый образ и запускаем снова
ssh $USER@$HOST -p $PORT "cd /opt/coin && docker compose pull coin-$CI_PROJECT_NAME"

ssh $USER@$HOST -p $PORT "cd /opt/coin && if docker inspect coin-$CI_PROJECT_NAME >/dev/null 2>&1; then docker compose stop coin-$CI_PROJECT_NAME && docker rm coin-$CI_PROJECT_NAME; fi && docker compose up coin-$CI_PROJECT_NAME -d"
