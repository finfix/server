#!/usr/bin/env bash

# Делаем, чтобы любая ошибка в скрипте приводила к ошибке в pipeline
set -o pipefail -e
# Проверяем, есть ли ssh-agent на сервере, если нет, то устанавливаем
command -v ssh-agent || ( apt-get update && apt-get install -y openssh-client )
# Запускаем ssh-agent
eval $(ssh-agent -s)
# Создаем папку для ssh ключа
mkdir -p ~/.ssh
# Раздаем права папке
chmod -R 700 ~/.ssh
# Добавляем ssh ключ в ssh-agent
echo "$SSH_KEY" | tr -d '\r' > ~/.ssh/id_rsa
# Раздаем права ssh ключу
chmod 600 ~/.ssh/id_rsa
# Добавляем хост в список известных, если
echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > ~/.ssh/config
# Добавляем ssh ключ в ssh-agent
ssh-add ~/.ssh/id_rsa
# Получаем публичный ключ сервера
ssh-keyscan -p $PORT $HOST >> ~/.ssh/known_hosts
# Раздаем права публичному ключу
chmod 644 ~/.ssh/known_hosts
