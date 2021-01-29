# Lead Integration - NodeJS

[Portuguese version](README-ptbr.md)

A basic sample using nodejs to show how to receive leads from zap+.

Please, check out our [docs](https://developers.grupozap.com/) for more information.

## Technology and Resources

- [NodeJS v14.15](https://nodejs.org/ko/blog/release/v14.15.2/)
- [Express](https://github.com/expressjs/express)
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
make local/run
```

**Helpful commands**

*Please, check all available commands in the [Makefile](Makefile) for more information*.
