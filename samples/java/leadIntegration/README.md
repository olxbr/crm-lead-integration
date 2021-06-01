# Lead Integration - Java

A basic sample using Java to show how to receive leads from zap+.

Please, check out our [docs](https://developers.grupozap.com/) for more information.

## Technology and Resources

- [Java 11](https://openjdk.java.net/projects/jdk/11/)
- [Maven 3.8.1](https://maven.apache.org/)
- [Spring Boot 2.5.0](https://spring.io/projects/spring-boot)
- [Docker](https://www.docker.com/get-started)

## How to install

### Cloning

``` shell
git clone https://github.com/olxbr/crm-lead-integration.git
cd samples/java
```

### Setting SECRET_KEY

Change `SECRET_KEY` valeu in [application.properties](src/main/resources/application.properties) file

### Build

- Local build execute:
``` shell
make local/build
```
- Docker build execute:
``` shell
make docker/build
```

### Running 

- Local run execute:
``` shell
make local/run
```
- Docker run execute:
``` shell
make docker/run
```

### Deploy

- Execute the following command to deploy:
``` shell
make deploy
```

**Helpful commands**

*Please, check all available commands in the [Makefile](Makefile) for more information*.