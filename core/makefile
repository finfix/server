proto:
	./scripts/proto.sh

update_go_dependencies:
	./scripts/update_go_dependencies.sh

mock:
	./scripts/mock.sh

dockerfile_build:
	../ci-scripts/dockerfile-build.sh core 8068

local_docker_build:
	docker build -f Dockerfile -t core .

compose_build:
	../ci-scripts/compose-build.sh core

local_docker_start:
	docker-compose --project-directory ../. up coin-core -d
