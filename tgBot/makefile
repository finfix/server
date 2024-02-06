proto:
	./scripts/proto.sh

mock:
	./scripts/mock.sh

dockerfile_build:
	../ci-scripts/dockerfile-build.sh auth 8071

local_docker_build:
	docker build -f Dockerfile -t auth .

compose_build:
	../ci-scripts/compose-build.sh auth

local_docker_start:
	docker-compose --project-directory ../. up coin-auth -d
