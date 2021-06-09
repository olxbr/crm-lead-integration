local/clean:
	rm -rf bin

local/build:
	go build -o bin/leadIntegration;

local/test:
	go test

local/run:
	go run main.go

local/setup:
	go get

bin/run:
	bin/leadIntegration

docker/build:
	@make local/build
	docker build -t crm-lead-integration/go-sample --no-cache .

docker/run:
	docker run -d -p 8000:8000 --name lead-integration-go crm-lead-integration/go-sample
