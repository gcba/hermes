# Admin

## Instalación

Implica clonar el repo, instalar las dependencias, configurar las variables de entorno, correr las migraciones y crear un usuario administrador.

### Prerequisitos

- PHP 7.1 con las extensiones **pgsql** y **openssl**.
- Postgres 10.
- Redis 4.
- Una cuenta de Mailgun.

Para instalar las dependencias es necesario tener [Composer](https://getcomposer.org/).

### Procedimiento

En `<REPO>/admin` instalar las dependencias con:

```bash
$ composer install
```

En el mismo directorio, crear un archivo `.env` con los valores de configuración:

```bash
APP_NAME=Hermes
APP_ENV=local
APP_KEY= # Dejar vacío por el momento, se lo completará a continuación
APP_DEBUG=true
APP_LOG_LEVEL=debug
APP_URL=http://localhost:8000
APP_DATETIME_FORMAT='d/m/Y H:i:s'

LDAP_URL=<LDAP_URL> # URL del WSDL del servicio de autenticación
DB_CONNECTION=pgsql

HERMES_READDB_HOST=localhost # Opcional, por defecto es 'localhost'
HERMES_READDB_PORT=5432 # Opcional, por defecto es '5432'
HERMES_READDB_NAME=hermes # Opcional, por defecto es 'hermes'
HERMES_READDB_USER=hermes # Opcional, por defecto es 'hermes'
HERMES_READDB_PASSWORD=<READDB_PASSWORD>
HERMES_READDB_SSLMODE=disable

HERMES_WRITEDB_HOST=localhost # Opcional, por defecto es 'localhost'
HERMES_WRITEDB_PORT=5432 # Opcional, por defecto es '5432'
HERMES_WRITEDB_NAME=hermes # Opcional, por defecto es 'hermes'
HERMES_WRITEDB_USER=hermes # Opcional, por defecto es 'hermes'
HERMES_WRITEDB_PASSWORD=<WRITEDB_PASSWORD>
HERMES_WRITEDB_SSLMODE=disable # Opcional, por defecto es 'disable'

BROADCAST_DRIVER=redis # Opcional, por defecto es 'redis'
CACHE_DRIVER=redis # Opcional, por defecto es 'redis'
QUEUE_DRIVER=redis # Opcional, por defecto es 'redis'
SESSION_DRIVER=redis # Opcional, por defecto es 'redis'
SESSION_SECURE_COOKIE=false # Opcional, por defecto es 'false'

REDIS_HOST=localhost # Opcional, por defecto es 'localhost'
REDIS_PASSWORD=<REDIS_PASSWORD>
REDIS_PORT=6379 # Opcional, por defecto es '6379'

MAIL_DRIVER=mailgun # Opcional, por defecto es 'mailgun'
MAIL_HOST=smtp.mailgun.org # Opcional, por defecto es 'smtp.mailgun.org'
MAIL_PORT=587 # Opcional, por defecto es '587'
MAIL_USER=<MAILGUN_USERNAME>
MAIL_PASSWORD=<MAILGUN_PASSWORD>
MAIL_ENCRYPTION=tls # Opcional, por defecto es 'tls'
MAIL_FROM=<EMAIL_FROM>
MAIL_SENDER="Gobierno de la Ciudad de Buenos Aires"
MAIL_REPLY_TO=<EMAIL_REPLYTO>
MAIL_SUBJECT="Gracias por tu comentario"
MAIL_CATCH_ALL=<EMAIL_CATCHALL> # Envía todos los mensajes a esta dirección

MAILGUN_DOMAIN=<MAILGUN_DOMAIN>
MAILGUN_API_KEY=<MAILGUN_APIKEY>
MAILGUN_VALIDATION_KEY=<MAILGUN_VALIDATIONKEY>

HERMES_RATINGS_PRIVATEKEY=<RATINGS_PRIVATEKEY> # Ruta a la clave privada
HERMES_STATS_PRIVATEKEY=<STATS_PRIVATEKEY> # Ruta a la clave privada
```

Para generar y guardar automáticamente un valor en `APP_KEY`:

```bash
$ php artisan key:generate
```

Luego correr las migraciones:

```bash
$ php artisan migrate --seed
```

Finalmente crear un usuario administrador:

```bash
$ php artisan admin:create "Juan Perez" juanperez@buenosaires.gob.ar
```

Ahora ya se puede levantar la aplicación.

```bash
$ php artisan serve
```

## Cola de tareas

Es necesaria para enviar y recibir mensajes por Mailgun. Corre como un proceso aparte:

```bash
$ php artisan queue:run redis
```

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

## Contenido por defecto

Si el entorno (`APP_ENV`) es `local`, al correr las migraciones se crearán registros de ejemplo en la base de datos: aplicaciones, calificaciones, mensajes, usuarios, etc.