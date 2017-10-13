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
    api: 'http://127.0.0.1:5000',
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
.then((response) => {
    response.json().then((json) => {
        // ...
    });
})
.catch((error) => {
    // ..
});
```