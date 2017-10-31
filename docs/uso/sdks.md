# SDKs

## JavaScript

### Instalación

#### Módulo ES2015+

Se encuentra en `<REPO>/sdks/js`.

##### NPM

```bash
$ npm install '<REPO>/sdks/js'
```

##### Yarn

```bash
$ yarn add file:<REPO>/sdks/js
```

##### Importación

```javascript
import Ratings from 'hermes-ratings-sdk';
```

#### Módulo UMD

Los archivos minificados y sin minificar están en `<REPO>/sdks/js/dist`.

```html
<script src="js/ratings.min.js"></script>
```

### Enviar una calificación

```javascript
const sdk = new Ratings({
    api: <ENDPOINT>,
    token: <TOKEN>,
    app: <APP>,
    version: '3.0',
    range: <RANGE>,
    platform: <PLATFORM>
});

// Opcional; omitir para calificaciones anónimas
sdk.user = {
    name: "Juan Pérez",
    email: "juan@example.com", // Opcional si está el mibaId
    mibaId: "04860d65-7e93-49e8-a983-a4007d23ffa5" // Opcional si está el email
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

### Instalación

#### Cocoapods

En el `Podfile` del proyecto:

```ruby
pod 'RatingsSDK', :path => '<REPO>/sdks/swift'
```

#### Manual

Agregar al proyecto los archivos `RatingsSDK.swift`, `RatingsUser.swift` y `RatingsError.swift` que están en `<REPO>/sdks/swift/RatingsSDK`.

### Enviar una calificación

```swift
import RatingsSDK
```

```swift
let sdk = Ratings(api: <ENDPOINT>, token: <TOKEN>, app: <APP>, platform: <PLATFORM>, range: <RANGE>)

sdk.timeout = 5 // Opcional; por defecto son 3 segundos

// Opcional; omitir para calificaciones anónimas
sdk.user = RatingsUser(
    name: "Juan Pérez",
    email: "juan@example.com", // Opcional si está el mibaId
    mibaId: "08108a49-4c68-47da-8510-93922b6b2d76" // Opcional si está el email
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
Ratings rating = new Ratings(<ENDPOINT>, <TOKEN>, <APP>, <PLATFORM>, <RANGE>);
RatingsResult result;

// Opcional; parámetros: Nombre, email, mibaID. El email es opcional si está el mibaId, y viceversa.
rating.setUser("Juan Pérez", "juan@example.com", "dc62591b-1cd3-4c6c-a943-f682e8860e08");

// Parámetos: calificación, descripción, mensaje. Descripcion y mensaje son opcionales.
result = rating.create(5, "Excelente", "Me encantó");

if (result.error != null) {
    // ..
}
```
