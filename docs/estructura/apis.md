# APIs

## Disposiciones comunes

### Autenticación

Delegada en la API Gateway.

## API de calificaciones

La API de calificaciones es primariamente de escritura. Tiene por fin poder agregar nuevas calificaciones y comentarios.

### Principios

#### 1.  Mostly RESTful

En donde resulte más apropiado (en vistas de las prioridades y necesidades del caso) se adoptará otro enfoque.

#### 2. Simplicidad

- Se seguirá un mismo esquema para todos los recursos.
- Se emplearán nombres concretos y autodescriptivos.
- Sólo se proveerán responses JSON.

#### 3. Consistencia

Las normas aquí descriptas deben aplicarse de manera uniforme y exhaustiva.

### Endpoints

`POST /ratings`

Crear una nueva calificación (que puede incluir un comentario o no)

`OPTIONS /ratings`

 Obtener la información necesaria (campos, tipos de datos, etc) necesaria para hacer un POST exitoso al endpoint.

## API de estadísticas

Esta API será de sólo lectura, con un enfoque de obtención y transformación de  datos cuantitativos. Tiene por fin poder responder preguntas relacionadas con los datos de forma confiable, flexible y expresiva.

### Endpoints

Dado que se trata de una API GraphQL, sólo es necesario un endpoint. En él se admiten los métodos GET y POST.

`GET /stats`

`POST /stats`