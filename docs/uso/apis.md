# API de calificaciones

Esta API no está pensada para ser usada directamente, sino a través de un SDK.

# API de estadísticas

### Endpoints

Hay un único endpoint POST.

`POST /stats`

### Headers

- **Content-Type:** `application/json; charset=utf-8`
- **Accept:** `application/json`

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

- **eq:** Igual a *(equal)*.
- **ne:** No igual a *(equal)*.
- **gt:** Mayor a *(greater than)*.
- **lt:** Menor a *(lower than)*.
- **gte:** Mayor o igual a *(greater than or equal)*.
- **lte:** Menor o igual a *(lower than or equal)*.

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