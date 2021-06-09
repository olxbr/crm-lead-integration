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

maven/build:
	./mvnw package

maven/setup:
	./mvnw install

maven/clean:
	./mvnw clean

maven/test:
	./mvnw test

maven/run:
	java -jar target/lead-integration-0.0.1.jar

gradle/build:
	./gradlew build

gradle/clean:
	./gradlew clean

gradle/test:
	./gradlew test

gradle/run:
	java -jar build/libs/lead-integration-0.0.1.jar
