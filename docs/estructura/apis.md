# APIs

## Disposiciones comunes

### Autenticación

Por token JWT firmado con clave privada.

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

