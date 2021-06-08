deploy:
	@make local/test
	@make local/clean
	@make local/build
	@make docker/build
	@make docker/run

docker/build:
	@make local/build
	docker build --build-arg JAR_FILE=target/*.jar -t crm-lead-integration/java-sample --no-cache .

docker/run:
	docker run -d -p 8080:8080 --name lead-integration-java crm-lead-integration/java-sample

docker/stop:
	docker container stop lead-integration-java
	docker container rm lead-integration-java

local/build:
	./mvnw package

local/run:
	java -jar target/leadIntegration-0.0.1.jar

local/setup:
	./mvnw install

local/clean:
	./mvnw clean

local/test:
	./mvnw test