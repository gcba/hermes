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

`POST /ratings`

Crear una nueva calificación (que puede incluir un comentario o no)

`OPTIONS /ratings`

 Obtener la información necesaria (campos, tipos de datos, etc) necesaria para hacer un POST exitoso al endpoint.

### Request

`POST /ratings`: sólo JSON body. Por lo tanto se requerirá el header `Content-Type: application/json`.

```json
{
    "rating": int,
    "comment": string, // optional
    "app": string,
    "range": string,
    "user": { // optional
        "name": string,
        "email": string
    },
    "platform": {
        "key": string,
        "version": string
    },
    "device": { // optional
        "name": string,
        "brand": string,
        "screen": {
            "width": int,
            "height": int,
            "PPI": float
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
       ...
   }
}
```

#### OPTIONS

```json
{
   "meta": {
       ...
   },
   "methods": [
       {
           "verb": ...,
           "url": ...,
           "headers": {
               ...
           },
           "parameters": {
               ...
           }
       },
   ]
}
```

## API de estadísticas

Esta API será de sólo lectura, con un enfoque de obtención y transformación de  datos cuantitativos. Tiene por fin poder responder preguntas relacionadas con los datos de forma confiable, flexible y expresiva.

### Endpoints

Dado que se trata de una API GraphQL, sólo es necesario un endpoint. En él se admiten los métodos GET y POST.

`GET /stats`

`POST /stats`