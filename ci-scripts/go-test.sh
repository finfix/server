#!/usr/bin/env bash

# Делаем, чтобы любая ошибка в скрипте приводила к ошибке в pipeline
set -o pipefail -e
# Устанавливаем curl и tar
apt-get update && apt-get install -y curl tar 1> /dev/null
# Устанавливаем go
curl -L -o go.tar.gz "https://dl.google.com/go/go1.20.4.linux-amd64.tar.gz"
# Создаем папку для установки go
mkdir -p /usr/local
# Распаковываем архив с go в папку /usr/local
tar -C /usr/local -xzf go.tar.gz
# Добавляем путь к go в переменную PATH
export PATH=$PATH:/usr/local/go/bin
# Устанавливаем зависимости проекта
go mod download
# Запускаем тесты
go test -v ./...
