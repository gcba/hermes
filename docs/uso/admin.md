# Admin

## Crear tokens para las APIs

### Por consola

#### Api de calificaciones

```bash
$ php artisan token:ratings
```

#### Api de estadísticas

```bash
$ php artisan token:stats
```

### Por el Admin

1. Iniciar sesión con un usuario administrador.
2. Ir a `Administración > Compass`.
3. Abrir la pestaña **Commands**.
4. Hacer click en el comando `php artisan token:ratings` o en `php artisan token:stats` (según corresponda) para que aparezca el botón **Run Command**.
5. Hacer click en **Run Command**. No hace falta escribir nada en **Additional Args?**.

El token generado permanecerá visible en Compass hasta que se haga click en **clear output**.
