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

1. Usuario
2. Soporte
3. Admin
4. Desarrollador
5. Público

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

### Admin

Todos los de Usuario y Soporte, más:

- Ver historial del usuario
- Ver log de actividades
- Filtrar log de actividades
- Agregar usuarios
- Asignar usuarios
- Borrar usuarios
- Borrar calificaciones