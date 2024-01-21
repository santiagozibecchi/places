Algunas ideas para ampliar el alcance del proyecto


## Homework

* Cantidad de visualizaciones por lugar
* Cambiar el tipo VARCHAR y USAR lo que corresponde (DATE) para las fechas!!!
* Reveer el modelado de la base de datos y agregar más relaciones
* Buscador de usarios
* Buscador de lugares
* Documentar todos los endpoints en POSTMAN
* Separar las rutas del main.go
* Agregar validacion para los email
* Agregar validacion para los generos, solo permitir 3 => female, male, other if true especificar tu genero, ademas siempre guardar en minuscula

## Interrogantes 
* No convendrá guardar toda la data en la db siempre en minuscula 🤔

## Extras

* Autenticación
* Cookies
* Dockerizar app
* Implementar repository
* Estudiar mongoDB
* Cantidad de lugares vistos por el usuario, cuales y cuales son los mas vistos
* Lugares favoritos del usuario
* Guardar preferencias
* Incluir API del clima

* Independizar la app de la base de datos, debe funcionar tanto para SQL como mongoDB (bson)


## Hacer uso de middlewares

* w.Header().Set("Content-Type", "application/json") // sería util meter esto en un middleware... en peticiones GET.
* [mux] http://localhost:8080/api/v1/places/ => Agregar una barra de mas da error 404, habría que evitarlo.

