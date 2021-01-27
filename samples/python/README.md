# Lead Integration - Python

[Portuguese version](README-ptbr.md)

A basic sample using python to show how to receive leads from zap+.

Please, check out our [docs](https://developers.grupozap.com/) for more information.


## Technology and Resources

- [Python 3.8.0](https://www.python.org/downloads/release/python-380/)
- [Pipenv](https://github.com/pypa/pipenv)
- [FastAPI](https://github.com/tiangolo/fastapi)
- [Uvicorn](https://github.com/encode/uvicorn)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)


## How to Install

#### Cloning

```
git clone <github_repo>
cd sample/python
```

#### Building

- Rename env.template to .env
- Define the [SECRET_KEY](https://developers.grupozap.com/leads/integration/#validacao-de-seguranca-secret-key) value in this file

```
make docker/build
```

## How to Run

#### Using docker

```
make docker/run
```

#### Locally

```
make local/install
make local/shell
make local/run
```

**Helpful commands**

*Please, check all available commands in the [Makefile](Makefile) for more information*.
