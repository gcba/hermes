# APIs

## Instalación

### Prerequisitos

- Go 1.9.
- Postgres 10.

### Procedimiento

Implica clonar el repo y configurar las variables de entorno.

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

# Keys
export HERMES_RATINGS_PUBLICKEY=<RATINGS_PUBLICKEY> # Ruta a la clave pública para la autenticación JWT
export HERMES_STATS_PUBLICKEY=<STATS_PUBLICKEY> # Ruta a la clave pública para la autenticación JWT
```

Luego cargarlo con:

```bash
$ source .env
```

Ahora se pueden correr las APIs.

```bash
$ go run main.go start ratings
```

```bash
$ go run main.go start stats
```

## API de calificaciones

Esta API no debe usarse directamente, sino por medio de los [SDKs](sdks.md).

## API de estadísticas

### Endpoints

Hay un único endpoint: `POST /stats`.

### Headers

|Key                 |Value                                    |
|--------------------|-----------------------------------------|
|**Content-Type**    |`application/json; charset=utf-8`        |
|**Accept**          |`application/json`                       |
|**Authorization**   |`Bearer <TOKEN>`                         |

### Body

#### Count

##### Tabla

Devuelve la cantidad de registros en una tabla.

```json
{
    "query": "query Example($field: Field!) { count(field: $field) }",
    "variables": {
        "field": {
            "name": "messages"
        }
    }
}
```

##### Columna

Devuelve la cantidad de registros en una columna, ignorando nulls.

```json
{
    "query": "query Example($field: Field!) { count(field: $field) }",
    "variables": {
        "field": {
            "name": "messages.transport_id"
        }
    }
}
```

##### Columna con condición

Devuelve la cantidad de registros en una columna que cumplen con una condición, ignorando nulls.

```json
{
    "query": "query Example($field: Field!) { count(field: $field) }",
    "variables": {
        "field": {
            "name": "messages.status",
            "eq": 0
        }
    }
}
```

##### AND

Permite agregar condiciones adicionales que deben verificarse conjuntamente.

```json
{
    "query": "query Example($field: Field!, $and: [Field!]) { count(field: $field, and: $and) }",
    "variables": {
        "field": {
            "name": "ratings.rating",
            "gt": 2
        },
        "and": {
            "name": "ratings.rating",
            "lt": 5
        }
    }
}
```

```json
{
    "query": "query Example($field: Field!, $and: [Field!]) { count(field: $field, and: $and) }",
    "variables": {
        "field": {
            "name": "ratings.rating",
            "gt": 2
        },
        "and": [
            {
                "name": "ratings.rating",
                "lt": 5
            },
            {
                "name": "ratings.has_message",
                "eq": true
            }
        ]
    }
}
```

##### OR

Permite agregar condiciones adicionales que no necesariamente deban verificarse a la vez.

```json
{
    "query": "query Example($field: Field!, $or: [Field!]) { count(field: $field, or: $or) }",
    "variables": {
        "field": {
            "name": "ratings.rating",
            "eq": 2
        },
        "or": {
            "name": "ratings.rating",
            "eq": 5
        }
    }
}
```

```json
{
    "query": "query Example($field: Field!, $or: [Field!]) { count(field: $field, or: $or) }",
    "variables": {
        "field": {
            "name": "ratings.rating",
            "eq": 2
        },
        "or": [
            {
                "name": "ratings.rating",
                "eq": 5
            },
            {
                "name": "ratings.rating",
                "eq": 3
            }
        ]
    }
}
```

#### Average

##### Columna

Devuelve el promedio de los valores en una columna numérica, ignorando nulls.

```json
{
    "query": "query Example($field: Field!) { average(field: $field) }",
    "variables": {
        "field": {
            "name": "ratings.rating"
        }
    }
}
```

##### Columna con condición

Devuelve el promedio de los valores en una columna numérica que cumplen con una condición, ignorando nulls.

```json
{
    "query": "query Example($field: Field!) { average(field: $field) }",
    "variables": {
        "field": {
            "name": "ratings.rating",
            "gt": 2
        }
    }
}
```

##### AND

Permite agregar condiciones adicionales que deben verificarse conjuntamente. Ídem **count**.

##### OR

Permite agregar condiciones adicionales que no necesariamente deban verificarse a la vez. Ídem **count**.

#### Operadores

Para construir condiciones se pueden usar los siguientes operadores:

|Key                 |Value                                    |
|--------------------|-----------------------------------------|
|**eq**              |Igual a *(equal)*                        |
|**ne**              |No igual a *(equal)*                     |
|**gt**              |Mayor a *(greater than)*                 |
|**lt**              |Menor a *(lower than)*                   |
|**gte**             |Mayor o igual a *(greater than or equal)*|
|**lte**             |Menor o igual a *(lower than or equal)*  |

#### Registros disponibles

En todos los casos los registros que han sido borrados lógicamente no son tenidos en cuenta para los cálculos.

### Responses

#### Count

```json
{
    "meta": {
        "code": ...,
        "message": ...
    },
        "data": {
        "count": ...
    }
}
```

#### Average

```json
{
  "meta": {
        "code": ...,
        "message": ...
    },
    "data": {
        "average": ...
    }
}
```

#### Errores

```json
{
    "meta": {
        "code": ...,
        "message": ...
    },
    "errors": [
        {
            "message": ...,
            "path": [ ... ]
        }
    ]
}
```