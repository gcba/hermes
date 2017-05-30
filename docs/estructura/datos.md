# Modelo de datos

![Schema](https://github.com/gcba/hermes/raw/master/docs/images/schema.png)

Consta de 12 tablas, sin contar las necesarias para roles y permisos (no aparecen en la imagen) dado que éstas son creadas y manejadas por un componente de autorización.

## Tablas

### User

El usuario del backend.

#### Relaciones

- **One-to-many** con User, a través de `modified_by`
- **Many-to-many** con App, a través de la tabla `User_App`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(70)           |       |       |       |       |
|email           |varchar(100)          |       |       |X      |       |
|password        |binary(60)            |       |       |       |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |
|modified_by     |int                   |X      |X      |       |       |

### AppUser

El usuario de las apps que envía calificaciones y comentarios.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(70)           |       |       |       |       |
|email           |varchar(100)          |       |X      |       |X      |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### App

Las apps sobre las que se envían calificaciones y comentarios.

#### Relaciones

- **One-to-many** con User, a través de `modified_by`
- **Many-to-many** con User, a través de la tabla `User_App`
- **Many-to-many** con Platform, a través de la tabla `App_Platform`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(50)           |       |       |X      |X      |
|type            |char                  |       |       |       |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |
|modified_by     |int                   |X      |X      |       |       |

### User_App

Tabla intermedia para la relación many-to-many entre User y App.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|user_id         |int                   |X      |       |       |X      |
|app_id          |int                   |X      |       |       |X      |

### Platform

La plataforma donde corren las apps y de donde provienen las calificaciones y comentarios.

#### Relaciones

- **Many-to-many** con App, a través de la tabla `App_Platform`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(15)           |       |       |X      |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### App_Platform

Tabla intermedia para la relación many-to-many entre App y Platform.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|app_id          |int                   |X      |       |       |X      |
|platform_id     |int                   |X      |       |       |X      |

### Device

En el caso de las aplicaciones móviles, el dispositivo desde donde se enviaron las calificaciones y comentarios.

#### Relaciones

- **One-to-many** con Brand, a través de `brand_id`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(30)           |       |       |X      |X      |
|screen_width    |int                   |       |       |       |       |
|screen_heigth   |int                   |       |       |       |       |
|ppi             |int                   |       |X      |       |       |
|brand_id        |int                   |X      |X      |       |X      |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### Browser

En el caso de las aplicaciones web, el browser desde donde se enviaron las calificaciones y comentarios.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(15)           |       |       |X      |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### Brand

En el caso de las aplicaciones móviles, la marca del dispositivo.

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(30)           |       |       |X      |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### Rating

Las calificaciones de las apps.

#### Relaciones

- **One-to-many** con App, a través de `app_id`
- **One-to-many** con AppUser, a través de `appuser_id`
- **One-to-many** con Platform, a través de `platform_id`
- **One-to-many** con Device, a través de `device_id`
- **One-to-many** con Browser, a través de `browser_id`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|rating          |smallint              |       |       |       |       |
|description     |varchar(30)           |       |X      |       |       |
|app_version     |varchar(15)           |       |X      |       |       |
|platform_version|varchar(15)           |       |       |       |       |
|browser_version |varchar(15)           |       |X      |       |       |
|has_message     |bool NULL             |       |X      |       |X      |
|app_id          |int                   |X      |       |       |X      |
|appuser_id      |int                   |X      |       |       |X      |
|platform_id     |int                   |X      |X      |       |X      |
|device_id       |int                   |X      |X      |       |X      |
|browser_id      |int                   |X      |X      |       |X      |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### Messages

Los mensajes de las conversaciones con los usuarios de las apps que enviaron calificaciones con comentarios.

#### Relaciones

- **One-to-many** con Rating, a través de `rating_id`
- **One-to-many** con AppUser, a través de `appuser_id`
- **One-to-many** con User, a través de `user_id`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|message         |text                  |       |       |       |       |
|direction       |char                  |       |       |       |Sí     |
|rating_id       |int                   |X      |       |       |X      |
|appuser_id      |int                   |X      |       |       |X      |
|user_id         |int                   |X      |       |       |X      |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### Config

Los valores de configuración del backend.

#### Relaciones

- **One-to-many** con User, a través de `modified_by`

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|key             |varchar(30)           |       |       |X      |       |
|value           |varchar(254)          |       |X      |       |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |
|modified_by     |int                   |X      |X      |       |       |


## Timestamps

Los campos `created_at` y `modified_at` son creados automáticamente por Laravel y no serán modificados.