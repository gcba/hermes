# Deploy

Hacer un deploy de Hermes en RHEL 7 implica:

1. Instalar y configurar PHP, Nginx, Go, Redis y Postgres
2. Instalar y configurar el Admin y las APIs
3. Configurar permisos
4. Configurar paths y scripts de inicio
5. Configurar systemd para los servicios y las APIs


## PHP

### Instalar

```bash
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


## Nginx

### Instalar

```bash
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
    root /srv/asi-236-hermes/admin/public;

    add_header X-Frame-Options "SAMEORIGIN";
    add_header X-XSS-Protection "1; mode=block";
    add_header X-Content-Type-Options "nosniff";

    index index.html index.htm index.php;

    charset utf-8;

    location /ratings {
        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /stats {
        proxy_pass http://127.0.0.1:7000;
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

```bash
# wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
# tar xzvf go1.9.2.linux-amd64.tar.gz
# mv go /usr/local/
```


## Redis

### Instalar

```bash
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

```bash
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

Asegurarse que el usuario bajo el que corren Nginx y PHP-FM pueden acceder a los archivos del Admin. Este usuario es generalmente `nginx`:

```bash
# usermod -a -G nginx,<USUARIO_OWNER_DEL_ADMIN> nginx
# sudo chown -R nginx:<USUARIO_OWNER_DEL_ADMIN> <REPO>/admin
```

### APIs

Crear un nuevo usuario bajo el cual correrán las APIs:

```bash
# useradd apis -s /sbin/nologin -M
```

Y asegurarse de que pueda acceder a los archivos de las APIs:

```bash
# usermod -a -G <USUARIO_OWNER_DE_LAS_APIS>,apis <USUARIO_OWNER_DE_LAS_APIS>
# chown -R <USUARIO_OWNER_DE_LAS_APIS>:apis <REPO>/apis
```