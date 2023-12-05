swagger:
	swag init -g ./app/cmd/main.go -o ./app/docs --parseDependency --parseInternal

update_go_dependencies:
	./scripts/update_go_dependencies.sh

dockerfile_build:
	../ci-scripts/dockerfile-build.sh json 8069

local_docker_build:
	docker build -f Dockerfile -t json .

compose_build:
	../ci-scripts/compose-build.sh json

local_docker_start:
	docker-compose --project-directory ../. up coin-json -d
