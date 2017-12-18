# APIs

## Instalación

### Prerequisitos

- Go 1.9.
- Postgres 10.

### Procedimiento

Implica clonar el repo, configurar las variables de entorno e instalar las dependencias.

En `<REPO>/apis/src/hermes` crear un archivo `.env` con los valores de configuración:

```bash
export GOPATH=<REPO>/apis

# Environment
export HERMES_RATINGS_ENV=DEV
export HERMES_STATS_ENV=DEV

# Ports
export HERMES_RATINGS_PORT=5000 # Se puede omitir, por defecto es '5000'
export HERMES_STATS_PORT=7000 # Se puede omitir, por defecto es '7000'

# Database
export HERMES_READDB_HOST=localhost # Se puede omitir, por defecto es 'localhost'
export HERMES_READDB_PORT=5432 # Se puede omitir, por defecto es '5432'
export HERMES_READDB_NAME=hermes
export HERMES_READDB_USER=hermes
export HERMES_READDB_PASSWORD=<READDB_PASSWORD>
export HERMES_READDB_SSLMODE=disable # Se puede omitir, por defecto es 'disable'.

export HERMES_WRITEDB_HOST=localhost # Se puede omitir, por defecto es 'localhost'
export HERMES_WRITEDB_PORT=5432 # Se puede omitir, por defecto es '5432'
export HERMES_WRITEDB_NAME=hermes
export HERMES_WRITEDB_USER=hermes
export HERMES_WRITEDB_PASSWORD=<WRITEDB_PASSWORD>
export HERMES_WRITEDB_SSLMODE=disable # Se puede omitir, por defecto es 'disable'.

# Keys -> deben estar en formato PEM
export HERMES_RATINGS_PUBLICKEY=<RATINGS_PUBLICKEY> # Ruta a la clave pública para la autenticación JWT
export HERMES_STATS_PUBLICKEY=<STATS_PUBLICKEY> # Ruta a la clave pública para la autenticación JWT
```

Luego cargarlo con:

```bash
$ source .env
```

En el mismo directorio instalar las dependencias con [dep](https://github.com/golang/dep):

```bash
$ dep ensure
```

Ahora se pueden correr las APIs.

```bash
$ go run main.go start ratings
```

```bash
$ go run main.go start stats
```

Para compilar el binario directamente:

```bash
$ go install hermes
```

El binario generado estará en `<REPO>/apis/bin`.