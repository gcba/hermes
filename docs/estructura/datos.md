# Modelo de datos

![Schema](https://github.com/gcba/hermes/raw/master/docs/images/schema.png)

Consta de 15 tablas, sin contar las necesarias para roles y permisos (no aparecen en la imagen) dado que éstas son creadas y manejadas automáticamente por un componente de autorización.

## Tablas

### User

El usuario del backend.

#### Relaciones

- **One-to-many** con User, a través de `updated_by`
- **Many-to-many** con App, a través de la tabla `App_User`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(70)           |       |       |       |       |
|email           |varchar(100)          |       |       |X      |       |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|updated_by      |int                   |X      |X      |       |       |

### AppUser

#### Relaciones

- **Many-to-many** con App, a través de la tabla `AppUser_App`
- **Many-to-many** con Platform, a través de la tabla `AppUser_Platform`
- **Many-to-many** con Device, a través de la tabla `AppUser_Device`

El usuario de las apps que envía calificaciones y comentarios.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(70)           |       |       |       |       |
|email           |varchar(100)          |       |X      |       |X      |
|miba_id         |char(36)              |       |X      |X      |X      |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |

### App

Las apps sobre las que se envían calificaciones y comentarios.

#### Relaciones

- **One-to-many** con User, a través de `updated_by`
- **Many-to-many** con User, a través de la tabla `App_User`
- **Many-to-many** con Platform, a través de la tabla `App_Platform`
- **Many-to-many** con AppUser, a través de la tabla `AppUser_App`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(50)           |       |       |X      |X      |
|type            |char                  |       |       |       |       |
|key             |char(32)              |       |       |X      |X      |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|updated_by      |int                   |X      |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |

### Platform

La plataforma donde corren las apps y de donde provienen las calificaciones y comentarios.

#### Relaciones

- **Many-to-many** con App, a través de la tabla `App_Platform`
- **Many-to-many** con AppUser, a través de la tabla `AppUser_Platform`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(15)           |       |       |X      |       |
|key             |char(32)              |       |       |X      |X      |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |

### Device

En el caso de las aplicaciones móviles, el dispositivo desde donde se enviaron las calificaciones y comentarios.

#### Relaciones

- **One-to-many** con Brand, a través de `brand_id`
- **One-to-many** con Platform, a través de `platform_id`
- **Many-to-many** con AppUser, a través de la tabla `AppUser_Device`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(30)           |       |       |X      |X      |
|screen_width    |int                   |       |       |       |       |
|screen_height   |int                   |       |       |       |       |
|ppi             |int                   |       |X      |       |       |
|brand_id        |int                   |X      |X      |       |X      |
|platform_id     |int                   |X      |       |       |X      |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |

### App_User

Tabla intermedia para la relación many-to-many entre User y App.

|Campos          |Tipo                  |Default| FK?   | Null? |Unique?|Index? |
|----------------|----------------------|-------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |       |
|user_id         |int                   |       |X      |       |       |X      |
|app_id          |int                   |       |X      |       |       |X      |

### App_Platform

Tabla intermedia para la relación many-to-many entre App y Platform.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|app_id          |int                   |X      |       |       |X      |
|platform_id     |int                   |X      |       |       |X      |

### AppUser_App

Tabla intermedia para la relación many-to-many entre AppUser y App.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|appuser_id      |int                   |X      |       |       |X      |
|app_id          |int                   |X      |       |       |X      |

### AppUser_Platform

Tabla intermedia para la relación many-to-many entre AppUser y Platform.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|appuser_id      |int                   |X      |       |       |X      |
|platform_id     |int                   |X      |       |       |X      |

### AppUser_Device

Tabla intermedia para la relación many-to-many entre AppUser y Device.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|appuser_id      |int                   |X      |       |       |X      |
|device_id       |int                   |X      |       |       |X      |

### Browser

En el caso de las aplicaciones web, el browser desde donde se enviaron las calificaciones y comentarios.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(15)           |       |       |X      |       |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |

### Brand

En el caso de las aplicaciones móviles, la marca del dispositivo.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(30)           |       |       |X      |       |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |

### Rating

Las calificaciones de las apps.

#### Relaciones

- **One-to-many** con App, a través de `app_id`
- **One-to-many** con Range, a través de `range_id`
- **One-to-many** con AppUser, a través de `appuser_id`
- **One-to-many** con Platform, a través de `platform_id`
- **One-to-many** con Device, a través de `device_id`
- **One-to-many** con Browser, a través de `browser_id`

|Campos          |Tipo                  |Default| FK?   | Null? |Unique?|Index? |
|----------------|----------------------|-------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |       |
|rating          |smallint              |       |       |       |       |       |
|description     |varchar(30)           |       |       |X      |       |       |
|app_version     |varchar(15)           |       |       |X      |       |       |
|browser_version |varchar(15)           |       |       |X      |       |       |
|platform_version|varchar(15)           |       |       |       |       |       |
|has_message     |bool                  |false  |       |       |       |X      |
|app_id          |int                   |       |X      |       |       |X      |
|range_id        |int                   |       |X      |       |       |X      |
|platform_id     |int                   |       |X      |       |       |X      |
|device_id       |int                   |       |X      |       |       |X      |
|appuser_id      |int                   |       |X      |X      |       |X      |
|browser_id      |int                   |       |X      |X      |       |X      |
|created_at      |timestamp             |       |       |       |       |       |
|updated_at      |timestamp             |       |       |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |       |

### Message

Los mensajes de las conversaciones con los usuarios de las apps que enviaron calificaciones con comentarios.

#### Relaciones

- **One-to-many** con Rating, a través de `rating_id`
- **One-to-many** con User, a través de `created_by`

|Campos          |Tipo                  |Default| FK?   | Null? |Unique?|Index? |
|----------------|----------------------|-------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |       |
|message         |text                  |       |       |       |       |       |
|direction       |enum                  |       |       |       |       |X      |
|status          |smallint              |0      |       |       |       |X      |
|transport_id    |varchar(90)           |       |       |X      |X      |X      |
|rating_id       |int                   |       |X      |       |       |X      |
|created_at      |timestamp             |       |       |       |       |       |
|created_by      |int                   |X      |X      |       |       |       |
|updated_at      |timestamp             |       |       |X      |       |       |

### Range

Los rangos de calificaciones de cada app.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(11)           |       |       |       |       |
|from            |smallint              |       |       |       |       |
|to              |smallint              |       |       |       |       |
|key             |char(32)              |       |       |X      |X      |
|created_at      |timestamp             |       |       |       |       |
|updated_at      |timestamp             |       |X      |       |       |
|deleted_at      |timestamp             |       |X      |       |       |

## Timestamps

Los campos `created_at` y `updated_at ` son creados automáticamente por Laravel y no serán modificados.

## Borrado

Será lógico para `Platform`, `Device`, `Browser`, `Brand`, `Rating`, `Range`, `AppUser`, `App`

Será físico para `User` y `Message`.