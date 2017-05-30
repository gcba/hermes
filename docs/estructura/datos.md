# Modelo de datos

![Schema](https://github.com/gcba/hermes/raw/master/docs/images/schema.png)

Consta de 11 tablas, sin contar las necesarias para roles y permisos (no aparecen en la imagen) dado que éstas son creadas y manejadas por una librería de autorización.

## Entidades

### User

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(70)           |       |       |       |       |
|email           |varchar(100)          |       |       |X      |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |
|modified_by     |int                   |X      |X      |       |       |

### AppUser

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(70)           |       |       |       |       |
|email           |varchar(100)          |       |X      |       |X      |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### App

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(50)           |       |       |X      |X      |
|type            |char                  |       |       |       |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |
|modified_by     |int                   |X      |X      |       |       |

### User_App

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|user_id         |int                   |X      |       |       |X      |
|app_id          |int                   |X      |       |       |X      |

### Platform

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(15)           |       |       |X      |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### App_Platform

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|app_id          |int                   |X      |       |       |X      |
|platform_id     |int                   |X      |       |       |X      |

### Device

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

### Brand

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|name            |varchar(30)           |       |       |X      |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### Rating

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|rating          |smallint              |       |       |       |       |
|description     |varchar(30)           |       |X      |       |       |
|app_version     |varchar(15)           |       |X      |       |       |
|platform_version|varchar(15)           |       |       |       |       |
|has_message     |bool NULL             |       |X      |       |X      |
|app_id          |int                   |X      |       |       |X      |
|appuser_id      |int                   |X      |       |       |X      |
|platform_id     |int                   |X      |X      |       |X      |
|device_id       |int                   |X      |X      |       |X      |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |

### Messages

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

|Campos          |Tipo                  | FK?   | Null? |Unique?|Index? |
|----------------|----------------------|:-----:|:-----:|:-----:|:-----:|
|id              |int (PK)              |       |       |       |       |
|key             |varchar(30)           |       |       |X      |       |
|value           |varchar(254)          |       |X      |       |       |
|created_at      |datetime              |       |       |       |       |
|modified_at     |datetime              |       |X      |       |       |
|modified_by     |int                   |X      |X      |       |       |