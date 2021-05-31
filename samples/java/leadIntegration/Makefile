docker/build:
	docker build --build-arg JAR_FILE=target/*.jar -t lead-integration-java .

docker/deploy:
	docker run -p 8080:8080 lead-integration-java

local/build:
	./mvnw package

local/run:
	java -jar target/leadIntegration-0.0.1-SNAPSHOT.jar

local/setup:
	./mvnw install

local/clean:
	./mvnw clean