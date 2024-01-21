# DUDAS random

* Cuando actualizo un campo de fila de x tabla, es realmente necesario devolver toda la data? 
Un usuario quiere editar su username, para que le voy a dar toda la data si quiere editar algo en específico
Quizar realmente convenga cuando se trate de un reporte...
Pero a su vez el usuario puede tener mucha data la cual editar:
direccion
edad
telefono
correo electronico
pero pero... si, es informacion editable. Pero queremos que pueda editar todo de una? por ahi el camino va mas por ahi.
si... que pueda editar todo lo que quiere, pero por secciones

Por ahi lo mas ideal seria un enpoint que edite los campos que quiera pero por seccion
EJ: 
Por un lado, en base al id se modifica la informacion relacionada al lugar donde vive el sujeto
direccion
codigo postal (Esto en realidad iria en otra tabla)
pais(Esto en realidad iria en otra tabla)

Por otro, campos relacionados a lo personal
edad
genero

Un usario puede tener N relaciones 
  
Esto igual se extiende a más posibilidades con diferentes tablas, interesante analizar caso por casos

--- 
Ejemplo random

Tengo una tabla(front) que maneja muchos campos
El front manda a editar solamente uno de esos campos
El front generalmente te manda toda la data sin importar ya que maneja un estado 


---
Otro
front: onFocus => manda una peticion para guardar 

--- 
Tambien es importante saber las especificaciones de que se edita y como se maneja el estado en el front, en base a esto pueden surgir N enpoint especifos.