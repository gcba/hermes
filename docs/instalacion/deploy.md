# Deploy

Hacer un deploy de Hermes en **RHEL 7** implica:

1. Instalar y configurar PHP, Nginx, Go, Redis y Postgres
2. Instalar y configurar el Admin y las APIs
3. Configurar permisos
4. Configurar paths y scripts de inicio
5. Configurar systemd para los servicios y las APIs


## PHP

### Instalar

```
# yum install rh-php71 rh-php71-php-fpm rh-php71-php-mbstring rh-php71-php-opcache rh-php71-php-pdo rh-php71-php-pgsql rh-php71-php-soap rh-php71-php-xml
# scl enable rh-php71 bash
```

### Configurar

En `/etc/opt/rh/rh-php71/php.ini`:

#### Cambiar

- `short_open_tag Off` a `short_open_tag On`.
- `cgi.fix_pathinfo 1` a `cgi.fix_pathinfo 0`.

#### Agregar

```ini
[opcache]
opcache.enable=1
opcache.memory_consumption=512
opcache.interned_strings_buffer=64
opcache.max_accelerated_files=32531
opcache.validate_timestamps=0
opcache.save_comments=1
opcache.fast_shutdown=0
```

En `/etc/opt/rh/rh-php71/php-fpm.d/www.conf`:

#### Cambiar

- `user = apache` a `user = nginx`
- `group = apache` a `group = nginx`
- `listen = 127.0.0.1:9000` a `listen = /var/run/php-fpm.sock`
- `;listen.owner = nobody` a `listen.owner = nginx`
- `;listen.group = nobody` a `listen.group = nginx`
- `;listen.mode = 0660` a `listen.mode = 0660`

## Nginx

### Instalar

```
# yum install nginx14
# scl enable nginx14 bash
```

### Configurar

En `/opt/rh/nginx14/root/etc/nginx/nginx.conf`:

#### Cambiar

- `gzip off` a `gzip on`.

#### Reemplazar

```conf
 server {
    listen 80;
    listen [::]:80;
    server_name localhost;
    root <REPO>/admin/public;

    add_header X-Frame-Options "SAMEORIGIN";
    add_header X-XSS-Protection "1; mode=block";
    add_header X-Content-Type-Options "nosniff";

    index index.html index.htm index.php;

    charset utf-8;

    location /ratings {
        proxy_pass http://127.0.0.1:<PUERTO_API_CALIFICACIONES>;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /stats {
        proxy_pass http://127.0.0.1:<PUERTO_API_ESTADISTICAS>;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /admin {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    error_page 404 /index.php;

    location ~ \.php$ {
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass unix:/run/php-fpm.sock;
        fastcgi_index index.php;
        include fastcgi.conf;
    }

    location ~ /\.(?!well-known).* {
        deny all;
    }
}
```


## Go

### Instalar

```
# wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
# tar xzvf go1.9.2.linux-amd64.tar.gz
# mv go /usr/local/
```


## Redis

### Instalar

```
# yum install rh-redis32-redis
# scl enable rh-redis32 bash
```

### Configurar

En `/etc/opt/rh/rh-redis32/redis.conf`:

#### Cambiar

- `supervised no` a `supervised systemd`.
- `# requirepass foobared` a `requirepass <REDIS_PASSWORD>`.


## Postgres

### Instalar

Si no está instalado por defecto:

```
# yum install postgresql
```

### Configurar

En `/var/lib/pgsql/data/pg_hba.conf`:

#### Remplazar

```conf
# "local" is for Unix domain socket connections only
local   all             all                                     peer
# IPv4 local connections:
host    all             all             127.0.0.1/32            trust
# IPv6 local connections:
host    all             all             ::1/128                 trust
```


## Admin y APIs

Seguir las instrucciones de instalación:

- [Admin](admin.md)
- [APIs](apis.md)

## Permisos

### Admin

Asegurarse que el usuario bajo el que corren Nginx y PHP-FPM (`nginx`) pueda acceder a los archivos del Admin:

```
# usermod -a -G nginx,<USUARIO_OWNER_DEL_ADMIN> nginx
# sudo chown -R nginx:<USUARIO_OWNER_DEL_ADMIN> <REPO>/admin
# sudo chmod -R 755 <REPO>/admin/storage
```

### APIs

Crear un nuevo usuario bajo el cual correrán las APIs:

```
# useradd apis -s /sbin/nologin -M
```

Y asegurarse de que pueda acceder a los archivos de las APIs:

```
# usermod -a -G <USUARIO_OWNER_DE_LAS_APIS>,apis <USUARIO_OWNER_DE_LAS_APIS>
# chown -R <USUARIO_OWNER_DE_LAS_APIS>:apis <REPO>/apis
```

## Paths y scripts de inicio

### Paths

En `/etc/profile.d` crear un archivo llamado `enable` con el siguiente contenido:

```bash
source /opt/rh/nginx14/enable
source /opt/rh/rh-php71/enable
source /opt/rh/rh-redis32/enable
```

Y otro llamado `export` que contenga:

```bash
PATH=$PATH:/usr/local/go/bin
PATH=$PATH:<REPO>/apis/bin
```

Luego cargar ambos archivos:

```bash
 $ source /etc/profile.d/enable
 $ source /etc/profile.d/export
```

### Systemd

#### PHP

Editar la configuración de Systemd:

```
# systemctl edit rh-php71-php-fpm
```

E ingresar lo siguiente:

```ini
.include /lib/systemd/system/rh-php71-php-fpm.service

[Service]
Restart=always
StartLimitInterval=0
```

#### Nginx

Editar la configuración de Systemd:

```
# systemctl edit nginx14-nginx
```

E ingresar lo siguiente:

```ini
.include /lib/systemd/system/nginx14-nginx.service

[Service]
Restart=always
StartLimitInterval=0
```

#### Redis

Editar la configuración de Systemd:

```
# systemctl edit rh-redis32-redis
```

E ingresar lo siguiente:

```ini
.include /lib/systemd/system/rh-redis32-redis.service

[Service]
Restart=always
StartLimitInterval=0
```

#### Postgres

Editar la configuración de Systemd:

```
# systemctl edit postgresql
```

E ingresar lo siguiente:

```ini
.include /lib/systemd/system/postgresql.service

[Service]
Restart=always
StartLimitInterval=0
```

#### Cola de tareas

En `/usr/lib/systemd/system/` crear un archivo llamado `hermes-queue.service` con el siguiente contenido:

```ini
[Unit]
Description=Laravel queue worker
After=rh-redis32-redis.service

[Service]
User=redis
Group=redis
Restart=always
ExecStart=/opt/rh/rh-php71/root/usr/bin/php <REPO>/admin/artisan queue:work redis --daemon --env=production
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
```

#### API de calificaciones

En `/usr/lib/systemd/system/` crear un archivo llamado `hermes-ratings.service` con el siguiente contenido:

```ini
[Unit]
Description=Ratings API
After=postgresql.service

[Service]
Type=simple
User=apis
Group=apis
Restart=always
ExecStart=/bin/bash -a -c 'source <REPO>/apis/src/hermes/.env && exec <REPO>/apis/bin/hermes start ratings'
StartLimitInterval=0
PermissionsStartOnly=true
ExecStartPre=/bin/mkdir -p /var/log/hermes-ratings
ExecStartPre=/bin/chown <USUARIO_OWNER_DE_LAS_APIS>:<USUARIO_OWNER_DE_LAS_APIS> /var/log/hermes-ratings
ExecStartPre=/bin/chmod 755 /var/log/hermes-ratings
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=hermes-ratings

[Install]
WantedBy=multi-user.target
```

#### API de estadísticas

En `/usr/lib/systemd/system/` crear un archivo llamado `hermes-stats.service` con el siguiente contenido:

```ini
[Unit]
Description=Stats API
After=postgresql.service

[Service]
Type=simple
User=apis
Group=apis
Restart=always
ExecStart=/bin/bash -a -c 'source <REPO>/apis/src/hermes/.env && exec <REPO>/apis/bin/hermes start stats'
StartLimitInterval=0
PermissionsStartOnly=true
ExecStartPre=/bin/mkdir -p /var/log/hermes-stats
ExecStartPre=/bin/chown <USUARIO_OWNER_DE_LAS_APIS>:<USUARIO_OWNER_DE_LAS_APIS> /var/log/hermes-stats
ExecStartPre=/bin/chmod 755 /var/log/hermes-stats
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=hermes-stats

[Install]
WantedBy=multi-user.target
```

#### Habilitar servicios

Finalmente, habilitar los servicios para que se ejecuten al bootear el sistema:

```
# systemctl daemon-reload
# chkconfig rh-php71-php-fpm on
# chkconfig nginx14-nginx on
# chkconfig rh-redis32-redis on
# chkconfig hermes-queue on
# chkconfig hermes-ratings on
# chkconfig hermes-stats on
```