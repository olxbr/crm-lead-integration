# Lead Integration - NodeJS

[English version](README.md)

Um simples projeto para demonstrar como receber leads do zap+ usando nodejs.

Por favor, confira nossa [documentação](https://developers.grupozap.com/) para mais informações.

## Recursos utilizados

- [NodeJS v14.15](https://nodejs.org/ko/blog/release/v14.15.2/)
- [Express](https://github.com/expressjs/express)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

## Como instalar

#### Clonando o projeto

```
git clone <github_repo>
cd sample/python
```

#### Criando a imagem docker

- Renomeie o arquivo `env.template` para `.env`
- Defina a [SECRET_KEY](https://developers.grupozap.com/leads/integration/#validacao-de-seguranca-secret-key) nesse arquivo.

```
make docker/build
```

## Como rodar

#### Usando docker

```
make docker/run
```

#### Localmente

```
make local/install
make local/run
```

**Mais comandos disponíveis**

*Confira todos os comandos disponíveis no arquivo [Makefile](Makefile)*.
