# Lead Integration - Go

[English version](README.md)

Um simples projeto usando Go para demonstrar como receber leads do zap+.

Por favor, confira nossa [documentação](https://developers.grupozap.com/) para mais informações.

## Recursos Utilizados

- [Go 1.16.4](https://golang.org/)
- [Gorilla Mux 1.8.0](https://github.com/gorilla/mux)
- [Joho GoDotEnv 1.3.0](https://github.com/joho/godotenv)
- [Docker](https://www.docker.com/get-started)

## Como instalar

### Clonando

``` shell
git clone https://github.com/olxbr/crm-lead-integration.git
cd samples/go
```

### Configurando SECRET_KEY

- Modifique o valor de `SECRET_KEY` no arquivo [.env.example](.env.example)
- Renomeie o arquivo `.env.example` para `.env`

### Build

- Para executar uma build local:
``` shell
make local/build
```
- Para executar uma build do Docker:
``` shell
make docker/build
```

### Executando 

- Para executar localmente:
``` shell
make local/run
```
- Para executar utilizando Docker:
``` shell
make docker/run
```

**Helpful commands**

*Please, check all available commands in the [Makefile](Makefile) for more information*.