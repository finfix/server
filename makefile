swagger:
	swag init -g ./app/cmd/main.go -o ./app/docs --parseDependency --parseInternal

local_docker_build:
	docker build -f Dockerfile -t coin-server .

local_docker_start:
	docker-compose --project-directory ../. up coin-server -d

lint:
	golangci-lint run
