# Places

## Endpoints habilitados

### /api/v1/users
Obtener todos los usuarios:               GET  /api/v1/users => Query Params (sort & kind)
Actualizar usuario:                       PUT  /api/v1/users/{id}
Eliminar usuario:                         DELETE /api/v1/users/{id}
Crear usuario: *                          POST  /api/v1/users
* require fields {
    "name": "",
    "lastname": "",
    "email": "",
    "username": "",
    "gender": ""
}

### Places
Obtener todos los lugares:                GET /api/v1/places
Obtener un lugar:                         GET /api/v1/places/{placeId}
Eliminar un lugar:                        DELETE /api/v1/places/{placeId}
Actualizar un lugar:                      PUT /api/v1/places/{placeId}
Crear un lugar: *                         POST /api/v1/places
* require fields {
    "kind": "",
    "name": "",
    "country": "",
    "location": "",
    "address": "",
    "start_time": "",
    "end_time": "",
    "description": ""
}

### Comments
Crear un comentario en el lugar:         POST /api/v1/place/{placeId}/user/{userId}

### Search
Buscar lugar:                            GET /api/v1/places/placeName/{name}