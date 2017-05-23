# Diseño

## Modelo de datos

### Entidades

## Admin

Se seguirán las convenciones y estructuras por defecto de Laravel, pero se buscará mantener livianos los controladores. Toda funcionalidad que resulte de utilidad común será extractada a un servicio aparte.

La lógica de negocio también estará encapsulada en un servicio aparte.

## APIs

Se estructurarán en capas:

```
           Base de datos
 _________________________________
|                                 |
|             Modelos             |
|_________________________________|
|                                 |
|           Controlador           |
|_________________________________|
|                |                |
|     Parser     |     Builder	  |
|________________|________________|
     Requests         Responses
```

Las funcionalidades comunes se extraerán en packages para ser utilizadas en ambas APIs.

Controlador, Parser y Builder, al ser únicos trabajarán con una abstracción de las operaciones que se terminará concretando con las llamadas a los Modelos. Para ello se utilizará una implementación del patrón Visitor.

En el caso de la API de estadísticas, el Controlador implementará una estructura similar a Pipes and Filters, a fin de obtener los datos necesarios y efectuar todas las transformaciones en el orden correcto. Dada su naturaleza, esta API no será REST sino Graphql.

Por el contrario, la API de calificaciones será netamente REST.