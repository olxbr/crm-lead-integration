# Lead Integration - Java

A basic sample using Java to show how to receive leads from zap+.

Please, check out our [docs](https://developers.grupozap.com/) for more information.

## Technology and Resources

- [Go 1.16.4](https://golang.org/)
- [Docker](https://www.docker.com/get-started)

## How to install

### Cloning

``` shell
git clone https://github.com/olxbr/crm-lead-integration.git
cd samples/go
```

### Setting SECRET_KEY

- Change `SECRET_KEY` valeu in [.env.example](.env.example) file
- Rename `.env.example` file to `.env`

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