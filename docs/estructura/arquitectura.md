# Arquitectura

## Lógica

### Admin

El Admin seguirá el patrón MVC (Model-View-Controller). Las vistas serán renderizadas por el servidor, y no se utilizarán frameworks frontend para gestionar la UI.

### APIs

Se utilizará una adaptación de MVC compuesta por el Modelo, un Controlador, un Parser de Requests y un Generador de Responses. Este último será el encargado de generar las responses JSON. El Parser transformará las responses a un formato intermedio, útil para el Controlador. Tanto el Parser como el Generador admitirán el uso de middlewares.

Controlador, Parser y Generador serán únicos para toda la aplicación, con posibilidades de adoptar reglas en el Controlador para casos particulares. Sólo en el caso del Modelo habrá uno para cada entidad.

## Física

![Architecture](https://github.com/gcba/hermes/raw/master/docs/images/architecture.png)

### Admin y APIs

Tanto el Admin como las APIs compartirán una máquina virtual. Los servicios no estarán expuestos de forma directa, sino detrás de Nginx.

### Base de datos

Se utlizarán dos servidores de bases de datos replicados en estructura master-slave. El servidor master estará encargado únicamente de las escrituras, y el slave será de sólo lectura.