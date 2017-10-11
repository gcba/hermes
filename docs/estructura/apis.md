# APIs

## Disposiciones comunes

### Autenticación

Delegada en la API Gateway.

## API de calificaciones

La API de calificaciones es primariamente de escritura. Tiene por fin poder agregar nuevas calificaciones y comentarios.

### Principios

#### 1. Mostly RESTful

En donde resulte más apropiado (en vistas de las prioridades y necesidades del caso) se adoptará otro enfoque.

#### 2. Simplicidad

- Se seguirá un mismo esquema para todos los recursos.
- Se emplearán nombres concretos y autodescriptivos.
- Sólo se proveerán responses JSON.

#### 3. Consistencia

Las normas aquí descriptas deben aplicarse de manera uniforme y exhaustiva.

### Endpoints

`OPTIONS /`

Obtener los endpoints de la API.

`OPTIONS /ratings`

 Obtener la información necesaria (campos, tipos de datos, etc) necesaria para hacer un POST exitoso al endpoint.

`POST /ratings`

Crear una nueva calificación (que puede incluir un comentario o no)

### Requests

`POST /ratings`: sólo JSON body. Por lo tanto se requerirá el header `Content-Type: application/json`.

```json
{
    "rating": int,
    "description": string,
    "comment": string, // optional
    "range": string,
    "app": {
        "key:" string,
        "version": string
    },
    "platform": {
        "key": string,
        "version": string
    },
    "user": { // optional
        "name": string,
        "email": string,
        "mibaId": string
    },
    "device": { // optional
        "name": string,
        "brand": string,
        "screen": {
            "width": int,
            "height": int,
            "ppi": float
        }
    },
    "browser": { // optional
        "name": string,
        "version": string
    }
}
```

### Responses

#### POST

```json
{
   "meta": {
       "code": ...,
       "message": ...
   }
}
```

#### OPTIONS

```json
{
   "meta": {
       "code": ...,
       "message": ...
   },
   "endpoints": [
       {
           "method": ...,
           "path": ...,
           "headers": {
               ...
           }
       },
   ]
}
```

#### Errores

```json
{
   "meta": {
       "code": ...,
       "message": ...
   },
   "errors": [... ]
}
```

## API de estadísticas

Esta API será de sólo lectura, con un enfoque de obtención y transformación de  datos cuantitativos. Tiene por fin poder responder preguntas relacionadas con los datos de forma confiable, flexible y expresiva.

### Endpoints

Dado que se trata de una API GraphQL, sólo es necesario un endpoint POST.

`POST /stats`

### Queries

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

DDevuelve el promedio de los valores en una columna numérica que cumplen con una condición, ignorando nulls.

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