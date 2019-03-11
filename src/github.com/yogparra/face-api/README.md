# Repositorio
```
>git clone https://github.com/yogparra/addgolang.git
>cd addgolang/src/github.com/yogparra/face-api
```

# Detalle
```
    Pruebas de conección a la api de el cognitive-services de azure "face-api"    
```

# Azure
```
Face
https://centralus.api.cognitive.microsoft.com/face/v1.0
```

# Ejecución
```
ejecutar
>go run face.go

Respueta 
[
  {
    "faceAttributes": {
      "age": 28,
      "gender": "male"
    },
    "faceId": "658315a8-36d5-452d-83ed-e2880cef1339",
    "faceRectangle": {
      "height": 133,
      "left": 79,
      "top": 85,
      "width": 133
    }
  }
]
```

# Documentación
```
https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395236
```