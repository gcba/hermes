# Análisis Funcional

## Requerimientos

### Funcionales

1. Registrar las calificaciones y comentarios enviados desde cada app
2. Generar estadísticas diarias, mensuales, anuales y pod período custom
3. Generar gráficos a partir de las estadísticas
4. Interactuar con el público calificador
5. Mostrar calificaciones y comentarios de cada app
6. Administrar el contenido
7. Administrar usuarios y permisos
8. Registrar los accesos y acciones realizados por cada usuario
9. Disponibilizar una API para las estadísticas

#### Estadísticas

TBD

### No funcionales

1. Alta disponiblidad
2. Fácil de extender
3. Fácil de mantener

## Actores

1. Usuario (Admin)
2. Soporte (Admin)
3. Supervisor (Admin)
4. Admin (Admin)
5. Desarrollador: Integra la aplicación con la API de calificaciones.
6. Público: Califica aplicaciones.

## Casos de uso

### Usuario

1. Seleccionar app
2. Ver calificaciones y comentarios
3. Ver mensajes
4. Filtrar mensajes
5. Ordenar mensajes
6. Buscar mensajes
7. Filtrar calificaciones
8. Ordenar calificaciones
9. Buscar comentarios
10. Ver estadísticas (gráficos)
11. Seleccionar métrica
12. Cambiar período de estadísticas
13. Ver usuario
14. Cambiar contraseña

### Soporte

Todos los de Usuario, más:

- Enviar mensaje al calificador

### Supervisor

Todos los de Soporte, más:

- Ver historial del usuario
- Ver log de actividades
- Filtrar log de actividades
- Agregar usuarios
- Asignar usuarios
- Borrar usuarios
- Crear apps
- Borrar apps
- Crear rangos
- Asignar rangos
- Crear plataformas

### Admin

Todos los de Usuario, Supervisor y Soporte, más:

- Borrar plataformas
- Borrar rangos
- Borrar calificaciones

## Roles y Permisos (Admin)

### Apps

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |X           |X           |               |
|Editar                |            |            |X           |X           |               |
|Borrar                |            |            |X           |X           |               |

### Users

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |            |            |X           |X           |               |
|Crear                 |            |            |X           |X           |               |
|Editar                |            |            |X           |X           |               |
|Borrar                |            |            |X           |X           |               |

### AppUsers

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |            |            |               |
|Editar                |            |            |            |            |               |
|Borrar                |            |            |            |X           |               |

### Ratings

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |            |            |               |
|Editar                |            |            |            |            |               |
|Borrar                |            |            |            |X           |               |

### Messages

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |X           |X           |X           |               |
|Editar                |            |            |            |            |               |
|Borrar                |            |            |            |X           |               |

### Ranges

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |X           |X           |               |
|Editar                |            |            |            |            |               |
|Borrar                |            |            |            |X           |               |

### Platforms

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |X           |X           |               |
|Editar                |            |            |X           |X           |               |
|Borrar                |            |            |            |X           |               |

### Devices

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |            |            |               |
|Editar                |            |            |            |            |               |
|Borrar                |            |            |            |X           |               |

### Browsers

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |            |            |               |
|Editar                |            |            |            |            |               |
|Borrar                |            |            |            |X           |               |

### Brands

|Permisos              | Usuario    | Soporte    | Supervisor | Admin      | Granularidad? |
|----------------------|:----------:|:----------:|:----------:|:----------:|---------------|
|Ver                   |X           |X           |X           |X           |Por app        |
|Crear                 |            |            |            |            |               |
|Editar                |            |            |            |            |               |
|Borrar                |            |            |            |X           |               |