local/clean:
	rm -rf bin

local/build:
	go build -o bin/lead-integration

local/test:
	go test

local/run:
	go run main.go

local/setup:
	go mod vendor && go mod tidy

bin/run:
	bin/lead-integration

docker/build:
	docker build -t crm-lead-integration/go-sample --no-cache .

docker/run:
	docker run -d -p 8000:8000 --name lead-integration-go crm-lead-integration/go-sample

docker/stop:
	docker container stop lead-integration-go
	docker container rm lead-integration-go
