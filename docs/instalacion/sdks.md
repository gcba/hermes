# SDKs

## JavaScript

### Módulo ES2015+

Se encuentra en `<REPO>/sdks/js`.

#### NPM

```bash
$ npm install '<REPO>/sdks/js'
```

#### Yarn

```bash
$ yarn add file:<REPO>/sdks/js
```

### Módulo UMD

Los archivos minificados y sin minificar están en `<REPO>/sdks/js/dist`.

```html
<script src="js/ratings.min.js"></script>
```

## Swift

Se requiere como mínimo iOS 9.0.

### Cocoapods

En el `Podfile` del proyecto:

```ruby
pod 'RatingsSDK', :git => 'https://github.com/gcba/hermes.git'
```

### Manual

Agregar al proyecto los archivos `RatingsSDK.swift`, `RatingsUser.swift` y `RatingsError.swift` que están en `<REPO>/sdks/swift/RatingsSDK`. Luego agregar las dependencias al `Podfile`:

```ruby
pod 'SwiftHTTP', '~> 2.0'
pod 'SwifterSwift/Foundation', '~> 3.1.1'
pod 'GBDeviceInfo', '~> 4.3'
```

## Java

Se requiere como mínimo el SDK Android 15.0.

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