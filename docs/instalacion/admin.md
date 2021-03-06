# Admin

## Instalación

Implica clonar el repo, instalar las dependencias, configurar las variables de entorno, correr las migraciones y crear un usuario administrador.

### Prerequisitos

- PHP 7.1 con las extensiones:
    - mbstring
    - openssl
    - opcache
    - pdo
    - pgsql
    - soap
    - tokenizer
    - xml
- Postgres 10 configurado (o no) con réplica master/slave.
- Redis 4.
- Una cuenta de Mailgun.

Para instalar las dependencias es necesario tener [Composer](https://getcomposer.org/).

### Procedimiento

Luego de clonar el repo, dentro de `<REPO>/admin` crear los siguientes directorios:

```
/storage
|___/app
|___/framework
|   |___/cache
|   |___/sessions
|   |___/views
|___/logs
```

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

HERMES_RATINGS_PRIVATEKEY=<RATINGS_PRIVATEKEY> # Ruta a la clave privada (RSA), debe estar en formato PEM
HERMES_STATS_PRIVATEKEY=<STATS_PRIVATEKEY> # Ruta a la clave privada (RSA), debe estar en formato PEM
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

## Mensajes

Para poder enviar y recibir mensajes es necesario configurar una route y un webhook en Mailgun y correr la cola de tareas.

### Route

Una route `catch_all` permite redirigir a Hermes todos los mensajes entrantes. Puede configurarse por consola:

```bash
$ php artisan mailgun:routes <APP_URL>/webhooks/messages/receive
```

El comando `mailgun:routes` puede también listar todas las routes configuradas en la cuenta de Mailgun o eliminarlas.

```bash
$ php artisan mailgun:routes --help

Usage:
  mailgun:routes [options] [--] [<url>]

Arguments:
  url                   The url of the route to create

Options:
      --delete          Delete all routes
      --list            List all existing routes
  -h, --help            Display this help message
  -q, --quiet           Do not output any message
  -V, --version         Display this application version
      --ansi            Force ANSI output
      --no-ansi         Disable ANSI output
  -n, --no-interaction  Do not ask any interactive question
      --env[=ENV]       The environment the command should run under
  -v|vv|vvv, --verbose  Increase the verbosity of messages: 1 for normal output, 2 for more verbose output and 3 for debug

Help:
  Create a new Mailgun Route
```

### Webhook

Un webhook de tipo `deliver` es necesario para que Mailgun pueda notificar a Hermes que un mensaje fue recibido por su destinatario. Puede configurarse por consola:

```bash
$ php artisan mailgun:webhooks deliver <APP_URL>/webhooks/messages/notify
```

El comando `mailgun:webhooks` puede también listar todos los webhooks configurados en la cuenta de Mailgun o eliminarlos.

```bash
$ php artisan mailgun:webhooks --help

Usage:
  mailgun:webhooks [options] [--] [<type>] [<url>]

Arguments:
  type                  The type of webhook to create. Can be: open, click, unsubscribe, spam, bounce, drop, deliver
  url                   The url of the webhook

Options:
      --delete          Delete all webhooks
      --list            List all existing webhooks
  -h, --help            Display this help message
  -q, --quiet           Do not output any message
  -V, --version         Display this application version
      --ansi            Force ANSI output
      --no-ansi         Disable ANSI output
  -n, --no-interaction  Do not ask any interactive question
      --env[=ENV]       The environment the command should run under
  -v|vv|vvv, --verbose  Increase the verbosity of messages: 1 for normal output, 2 for more verbose output and 3 for debug

Help:
  Create a new Mailgun Webhook
```

### Cola de tareas

Envía/recibe mensajes y realiza operaciones sobre los mismos. Corre como un proceso aparte:

```bash
$ php artisan queue:work redis
```

## Contenido por defecto

Si el entorno (`APP_ENV`) es `local`, al correr las migraciones se crearán registros de ejemplo en la base de datos: aplicaciones, calificaciones, mensajes, usuarios, etc.