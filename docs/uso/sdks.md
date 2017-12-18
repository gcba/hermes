# SDKs

## JavaScript

### Enviar una calificación

```javascript
// Módulo ES2015+
import Ratings from 'ratings-sdk';
```

```javascript
const sdk = new Ratings({
    api: <ENDPOINT>,
    token: <TOKEN>,
    app: <APP>,
    version: '3.0',
    range: <RANGE>,
    platform: <PLATFORM>
});

// Opcional, omitir para calificaciones anónimas
sdk.user = {
    name: 'Juan Pérez',
    email: 'juan@example.com', // Opcional si está el mibaId
    mibaId: 'dc62591b-1cd3-4c6c-a943-f682e8860e08' // Opcional si está el email
}

sdk.create({
    rating: 5,
    description: 'Excelente', // Opcional
    comment: 'Me encantó' // Opcional
})
.then((json) => {
    // ...
})
.catch((errors) => {
    // ...
});
```

## Swift

### Enviar una calificación

```swift
import RatingsSDK
```

```swift
let sdk = Ratings(api: <ENDPOINT>, token: <TOKEN>, app: <APP>, platform: <PLATFORM>, range: <RANGE>)

sdk.timeout = 5 // Opcional, por defecto son 3 segundos

// Opcional, omitir para calificaciones anónimas
sdk.user = RatingsUser(
    name: "Juan Pérez",
    email: "juan@example.com", // Opcional si está el mibaId
    mibaId: "dc62591b-1cd3-4c6c-a943-f682e8860e08" // Opcional si está el email
)

// Description y comment son opcionales
sdk.create(rating: 5, description: "Excelente", comment: "Me encantó") { response, error in
    guard error == nil else {
        // ...
    }

    // ...
}
```

## Java

### Enviar una calificación

```java
import gcba.ratings.sdk.Ratings;
import gcba.ratings.sdk.RatingsResult;
```
```java
Ratings sdk = new Ratings(<ENDPOINT>, <TOKEN>, <APP>, <PLATFORM>, <RANGE>);
RatingsResult result;

sdk.setTimeout(5); // Opcional, por defecto son 3 segundos

// Opcional, omitir para calificaciones anónimas
// Parámetros: nombre (String), email (String), mibaID (String)
// El email es opcional si está el mibaId, y viceversa
sdk.setUser("Juan Pérez", "juan@example.com", "dc62591b-1cd3-4c6c-a943-f682e8860e08");

// Parámetros: calificación (int), descripción (String), mensaje (String)
// Descripción y mensaje son opcionales
result = sdk.create(5, "Excelente", "Me encantó");

if (result.error != null) {
    // ..
}
```
