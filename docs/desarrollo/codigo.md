# Código

## Stack

### Admin

PHP 7.1 con Laravel 5.4

### APIs

Go 1.8 con Echo y Gorm

### Datos

Postgresql, Redis e InfluxDB

### API Gateway

Tyk

## Convenciones

### Go

El código debe formatearse con la herramienta `gofmt`.

### PHP

Se seguirá el estándar PSR-1.

https://github.com/php-fig/fig-standards/blob/master/accepted/PSR-1-basic-coding-standard.md

## Testeo

### APIs

Para las APIs el testeo será automatizado, con el objetivo de conseguir una cobertura del 100%. Se escribirán los siguientes tipos de tests:

- Unitarios
- Integración
- End-to-end

### Admin

En el caso del Admin, se testarán en forma automática las clases propias de la aplicación y en forma manual los flujos de la interfaz gráfica. Los tests automatizados serán sólo unitarios, y no habrá un objetivo concreto de cobertura.

## Repositorio

Todo el código y la documentación se alojarán en GitHub, en un solo repositorio.

https://github.com/gcba/hermes

### Ramas

Se trabajará con un esquema de 3 ramas: dev, qa y master, con una rama adicional para el deploy de la documentación.

#### dev

En dev se realizará el desarrollo. Al finalizar un sprint, se mergeará a qa para testear.

#### qa

En qa residirá un snapshot de código del sprint anterior para testar automática y manualmente. Al finalizar un sprint, si no quedan bugs sin solucionar (del sprint anterior) se mergeará a master. Si se fueran juntando bugs de varios sprints y no se puede mergear a master, se realizará un sprint dedicado a la corrección de bugs.

#### master

En master residirá el código de las releases más recientes. En las primera etapa del desarrollo (cuando no hay versiones usables aún) contendrá el código que más recientemente haya superado la etapa de qa.

Cuando ya se estén taggeando releases cada nuevo merge a master supondrá la creación de una nueva release.

#### gh_pages

Esta rama se utilizará para servir la documentación a través de Github Pages.

### Commits

Los mensajes de los commits deberán ser breves, concisos y estar escritos en inglés.

### Releases

Se seguirá el esquema SemVer (semantic versioning): major.minor.patch.

Es decir, cuando haya breaking changes se incrementará la versión major; cuando haya nuevas funcionalidades pero no breaking changes se incrementará la versión minor; y cuando sólo haya bugfixes pero no nuevas funcionalidades ni breaking changes se incrementará la versión patch.

En una primer etapa no habrá releases hasta que se complete una versión funcional del proyecto.

### Estructura de carpetas

```
/hermes
|___/admin
|   |____ ...
|___/apis
|	|___/src
|	    |___/ratings
|	    |	|___ ...
|    	|___/stats
|    	|	|___ ...
|		|___ ...
|___/sdks
|	|___/java
|	|	|___ ...
|	|___/swift
|	|	|___ ...
|	|___/js
|		|___ ...
|___/docs
	|___ ...
```

`hermes` es el directorio base de la aplicación.

`hermes/admin` contiene el proyecto Laravel del admin.

`hermes/apis` es el workspace de Go. Aquí debe apuntar la GOPATH.

`hermes/apis/src` contiene los proyectos Go (apis y librerías comunes propias).

`hermes/apis/src/ratings` contiene el proyecto Echo de la API de calificaciones.

`hermes/apis/src/stats` contiene el proyecto Echo de la API de estadísticas.

`hermes/sdks` contiene las librerías que facilitarán a las aplicaciones interactuar con la API de calificaciones.

`hermes/sdks/java` contiene el cliente para Java.

`hermes/sdks/swift` contiene el cliente para Swift.

`hermes/sdks/js` contiene el cliente para JavaScript.

`hermes/docs` contiene la documentación del proyecto, que se mantendrá sincronizada a la rama gh_pages mediante un build script.