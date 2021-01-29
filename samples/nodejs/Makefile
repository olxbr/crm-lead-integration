APP_NAME="lead-integration-nodejs"
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    DOCKER_USER=$(shell id -u $(USER)):$(shell id -g $(USER))
endif
ifeq ($(UNAME_S),Darwin)
    DOCKER_USER=
endif

local/install:
	npm install

local/lint:
	npm run lint

local/run:
	npm run start

local/test:
	NODE_ENV=test npm run test

local/coverage:
	NODE_ENV=test npm run cov

local/verify:
	make local/lint
	make local/test
	make local/coverage

docker/build:
	CURRENT_UID=${DOCKER_USER} docker-compose build

docker/lint:
	CURRENT_UID=${DOCKER_USER} docker-compose run ${APP_NAME} npm run lint

docker/run:
	CURRENT_UID=${DOCKER_USER} NODE_ENV=test docker-compose run --service-port ${APP_NAME} npm run start

docker/test:
	CURRENT_UID=${DOCKER_USER} NODE_ENV=test docker-compose run ${APP_NAME} npm run test

docker/coverage:
	CURRENT_UID=${DOCKER_USER} NODE_ENV=test docker-compose run ${APP_NAME} npm run cov

docker/verify:
	make docker/lint
	make docker/test
	make docker/coverage
