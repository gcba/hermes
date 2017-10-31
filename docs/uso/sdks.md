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

Se requiere como mínimo iOS 9.0.

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

Se requiere como mínimo el SDK Android 15.0.

### Instalación

1. Importar en Android Studio el archivo `ratings-sdk-release.aar` que se encuentra en `<REPO>/sdks/java/ratings-sdk/build/outputs/aar`:
```
File > New > New Module... > Import .JAR/.AAR Package
```

2. Agregar las dependencias en `build.gradle`:
```groovy
compile project(':ratings-sdk-release')
compile 'com.google.code.gson:gson:2.8.1'
compile('com.goebl:david-webb:1.3.0') {
    exclude group: 'org.json', module: 'json'
}
```

3. Asegurarse que el permiso para acceder a internet esté en `AndroidManifest.xml`:
```xml
<uses-permission android:name="android.permission.INTERNET"/>
```

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
// Parámetros: nombre (String), email (String), mibaID (String). El email es opcional si está el mibaId, y viceversa.
sdk.setUser("Juan Pérez", "juan@example.com", "dc62591b-1cd3-4c6c-a943-f682e8860e08");

// Parámetros: calificación (int), descripción (String), mensaje (String). Descripción y mensaje son opcionales.
result = sdk.create(5, "Excelente", "Me encantó");

if (result.error != null) {
    // ..
}
```
