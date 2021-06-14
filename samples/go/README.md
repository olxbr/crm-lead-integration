# Lead Integration - Go

[Portuguese version](README-ptbr.md)

A basic sample using Go to show how to receive leads from zap+.

Please, check out our [docs](https://developers.grupozap.com/) for more information.

## Technology and Resources

- [Go 1.16](https://golang.org/)
- [Gorilla Mux 1.8](https://github.com/gorilla/mux)
- [Joho GoDotEnv 1.3](https://github.com/joho/godotenv)
- [Docker](https://www.docker.com/get-started)

## How to install

### Cloning

``` shell
git clone https://github.com/olxbr/crm-lead-integration.git
cd samples/go
```

### Setting SECRET_KEY

- Copy `.env.example` file and rename it to `.env`
- Change `SECRET_KEY` value in `.env` file

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

**Helpful commands**

*Please, check all available commands in the [Makefile](Makefile) for more information*.