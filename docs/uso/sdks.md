# SDKs

## JavaScript

### Instalación

### Con &lt;script&gt;

El build script genera un módulo **UMD**. Aechivos minificados y sin minificar están en `<REPO>/sdks/js/dist`.

```html
<script src="js/ratings.min.js"></script>
```

### Con

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