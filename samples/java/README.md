# Lead Integration - Java

[Portuguese version](README-ptbr.md)

A basic sample using Java to show how to receive leads from zap+.

Please, check out our [docs](https://developers.grupozap.com/) for more information.

## Technology and Resources

- [Java 11](https://openjdk.java.net/projects/jdk/11/)
- [Maven 3.8.1](https://maven.apache.org/)
- [Gradle 7.0.2](https://gradle.org/)
- [Spring Boot 2.5.0](https://spring.io/projects/spring-boot)
- [Lombok 1.18.20](https://projectlombok.org/features/all)  
- [Docker](https://www.docker.com/get-started)

## How to install

### Cloning

``` shell
git clone https://github.com/olxbr/crm-lead-integration.git
cd samples/java
```

### Setting SECRET_KEY

Change `SECRET_KEY` value in [application.properties](src/main/resources/application.properties) file

### Build

- Maven local build execute:
``` shell
make maven/build
```
- Gradle local build execute:
``` shell
make gradle/build
```
- Docker build execute:
``` shell
make docker/build
```

### Running 

- Maven local run execute:
``` shell
make maven/run
```
- Gradle local run execute:
``` shell
make gradle/run
```
- Docker run execute:
``` shell
make docker/run
```

**Helpful commands**

*Please, check all available commands in the [Makefile](Makefile) for more information*.
