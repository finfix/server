update_dependencies:
	./scripts/update_go_depedencies.sh

proto:
	./scripts/proto.sh

dockerfile_build:
	../ci-scripts/dockerfile-build.sh logger 8070

local_docker_build:
	docker build -f Dockerfile -t logger .

compose_build:
	../ci-scripts/compose-build.sh logger

local_docker_start:
	docker-compose --project-directory ../. up coin-logger -d
