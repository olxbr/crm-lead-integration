# Lead Integration - Python

[English version](README.md)

Um simples projeto para demonstrar como receber leads do zap+ usando python.

Por favor, confira nossa [documentação](https://developers.grupozap.com/) para mais informações.


## Recursos utilizados

- [Python 3.8.0](https://www.python.org/downloads/release/python-380/)
- [Pipenv](https://github.com/pypa/pipenv)
- [FastAPI](https://github.com/tiangolo/fastapi)
- [Uvicorn](https://github.com/encode/uvicorn)
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
make local/shell
make local/run
```

**Mais comandos disponíveis**

*Confira todos os comandos disponíveis no arquivo [Makefile](Makefile)*.
