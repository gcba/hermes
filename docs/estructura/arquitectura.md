# Arquitectura

## Lógica

### Admin

El Admin seguirá el patrón MVC (Model-View-Controller). Las vistas serán renderizadas por el servidor, y no se utilizarán frameworks frontend para gestionar la UI.

### APIs

Se utilizará una adaptación de MVC compuesta por el Modelo, un Controlador, un Parser de Requests y un Generador de Responses. Este último será el encargado de generar las responses JSON. El Parser transformará las responses a un formato intermedio, útil para el Controlador. Tanto el Parser como el Generador admitirán el uso de middlewares.

Controlador, Parser y Generador serán únicos para toda la aplicación, con posibilidades de adoptar reglas en el Controlador para casos particulares. Sólo en el caso del Modelo habrá uno para cada entidad.

## Física

### Admin y APIs

Tanto el Admin como las APIs compartirán una máquina virtual. El Admin no estará expuesto en forma directa, sino detrás de Nginx. En el caso de las APIs un API Gateway tomará el lugar de Nginx.

### Base de datos

Se utlizarán dos servidores de bases de datos replicados en estructura master-slave. El servidor master estará encargado únicamente de las escrituras, y el slave será de sólo lectura.