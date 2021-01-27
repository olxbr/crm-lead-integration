APP_NAME="lead-integration-python"
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    DOCKER_USER=$(shell id -u $(USER)):$(shell id -g $(USER))
endif
ifeq ($(UNAME_S),Darwin)
    DOCKER_USER=
endif

local/lint:
	flake8 app/

local/check-packages:
	pipenv check --system -e PIPENV_PYUP_API_KEY=""

local/bandit:
	bandit -r app *.py

local/install:
	pipenv install --dev --skip-lock

local/shell:
	pipenv shell

local/run:
	python run.py

local/test:
	python -m pytest -c tests/pytest.ini \
	--pyargs ./tests -v --junitxml=results.xml \
	--cov-fail-under 100 --cov-report xml \
	--cov-report term \
	--cov-report html --cov ./app

docker/build:
	CURRENT_UID=${DOCKER_USER} docker-compose build ${APP_NAME}

docker/up:
	CURRENT_UID=${DOCKER_USER} docker-compose up -d

docker/down:
	CURRENT_UID=${DOCKER_USER} docker-compose down --remove-orphans

docker/lint:
	CURRENT_UID=${DOCKER_USER} docker-compose run ${APP_NAME} flake8 app/

docker/check-packages:
	CURRENT_UID=${DOCKER_USER} docker-compose run -e PIPENV_PYUP_API_KEY="" ${APP_NAME} pipenv check --system

docker/bandit:
	CURRENT_UID=${DOCKER_USER} docker-compose run ${APP_NAME} bandit -r app *.py

docker/verify:
	make docker/lint
	make docker/bandit
	make docker/check-packages

docker/run:
	CURRENT_UID=${DOCKER_USER} docker-compose run --service-port ${APP_NAME} python run.py

docker/test:
	CURRENT_UID=${DOCKER_USER} docker-compose run ${APP_NAME} \
	python -m pytest -s -c tests/pytest.ini \
	--pyargs ./tests -v  \
	--cov-fail-under 100 --cov-report xml \
	--cov-report term \
	--cov-report html --cov ./app
