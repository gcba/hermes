# SDKs

## JavaScript

### Instalación

#### Módulo ES6

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

sdk.create({
    rating: 5,
    description: 'Excelente',
    comment: 'Me encantó'
})
.then((json) => {
    // ...
})
.catch((error) => {
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

// Rating only; no user
sdk.create(rating: 5) { response, error in
    guard error == nil else {
        // ...
    }

    // ...
}

// Rating and description only; no user
sdk.create(rating: 4, description: "Bueno") { response, error in
    guard error == nil else {
        // ...
    }

    // ...
}

// Rating, description and comment; no user
sdk.create(rating: 3, description: "Regular", comment: "Lorem ipsum dolor...") { response, error in
    guard error == nil else {
        // ...
    }

    // ...
}

// Rating, description and comment; user name and mibaId only
sdk.user = RatingsUser(name: "Juan Pérez", mibaId: "04860d65-7e93-49e8-a983-a4007d23ffa5")

sdk.create(rating: 2, description: "Malo", comment: "Lorem ipsum dolor...") { response, error in
    guard error == nil else {
        // ...
    }

    // ...
}

// Rating, description and comment; user name and email only
sdk.user = RatingsUser(name: "Juan Pérez", email: "juan@example.com")

sdk.create(rating: 1, description: "Muy Malo", comment: "Lorem ipsum dolor...") { response, error in
    guard error == nil else {
        // ...
    }

    // ...
}

// Rating, description and comment; user name, email and mibaId
sdk.user = RatingsUser(name: "Juan Pérez", email: "juan@example.com", mibaId: "08108a49-4c68-47da-8510-93922b6b2d76")

sdk.create(rating: 5, description: "Muy Bueno", comment: "Lorem ipsum dolor...") { response, error in
    guard error == nil else {
        // ...
    }

    // ...
}
```