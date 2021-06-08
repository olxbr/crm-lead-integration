# Lead Integration - Java

[English version](README.md)

Um simples projeto usando Java para demonstrar como receber leads do zap+.

Por favor, confira nossa [documentação](https://developers.grupozap.com/) para mais informações.

## Recursos Utilizados

- [Java 11](https://openjdk.java.net/projects/jdk/11/)
- [Maven 3.8.1](https://maven.apache.org/)
- [Gradle 7.0.2](https://gradle.org/)
- [Spring Boot 2.5.0](https://spring.io/projects/spring-boot)
- [Lombok 1.18.20](https://projectlombok.org/features/all)
- [Docker](https://www.docker.com/get-started)

## Como instalar

### Clonando

``` shell
git clone https://github.com/olxbr/crm-lead-integration.git
cd samples/java
```

### Configurando SECRET_KEY

Modifique o valor de `SECRET_KEY` no arquivo [application.properties](src/main/resources/application.properties)

### Build

- Para executar uma build local (Maven):
``` shell
make maven/build
```
- Para executar uma build local (Gradle):
``` shell
make gradle/build
```
- Para executar uma build do Docker:
``` shell
make docker/build
```

### Executando

- Para executar localmente (Maven):
``` shell
make maven/run
```
- Para executar localmente (Gradle):
``` shell
make gradle/run
```
- Para executar utilizando Docker:
``` shell
make docker/run
```

**Mais comandos disponíveis**

*Confira todos os comandos disponíveis no arquivo [Makefile](Makefile)*.
