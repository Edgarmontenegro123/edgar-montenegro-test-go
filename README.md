# think-us-challenge

## Piezas del pastel (10 puntos)
1. Problemas de diseño – 1 puntos
   Has sido encargado de plantear una solución al siguiente problema.
   Imagina que tenemos un sistema escrito en cualquier lenguaje de programación compilado,
   este sistema se conecta a una base de datos SQL, el sistema funciona muy bien cuando hay
   baja demanda de transacciones, pero cuando la cantidad de transacciones aumenta el sistema
   deja de responder solicitudes, se encontró que la base de datos es el cuello de botella y no
   acepta solicitudes en paralelo para completar las transacciones. ¿Qué solución plantearías
   para receptar más cantidad de transacciones adaptándonos al cuello de botella de la base de
   datos?

   
2. Encripta tu mensaje – 3.5 puntos
   Has sido encargado de desarrollar una nueva forma de encriptar comunicaciones.
   Básicamente, cada vocal del mensaje de entrada deberá ser precedida por otra cadena,
   llamada la clave. La función recibirá dos parámetros de cadena: el primero será la clave y el
   segundo, el mensaje. La función devolverá una cadena.
   Ejemplo:
   Clave: dcj
   Mensaje: I love prOgrAmming!
   Esto debería devolver: dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!
   No se tomarán en cuenta las vocales con acentos.
   Cuando se reciba una cadena vacía, debería devolver una cadena vacía. Si el mensaje es
   nulo o vacío, devolver una cadena vacía. Si la clave es nula o vacía, entonces utiliza
   DCJ como valor predeterminado.


3. Suma a cero – 3.5 puntos
   Dado un arreglo de enteros, elimina todos los nodos consecutivos cuya suma sea cero y
   devuelve los nodos restantes. Un arreglo vacío también puede ser un resultado válido. Si se
   recibe un valor nulo, devuelve un arreglo vacío.
   Ejemplo:
   Entrada: [3, 4, -7, 5, -6, 2, 5, -1, 8]
   Salida: [5, 8]


4. Aprendizaje – 1 puntos
   Se te ha asignado poder entender esta arquitectura para poder aplicarlo a un
   nuevo proyecto, se necesita que expliques a detalle cual es la función de cada
   término mostrado en la imagen, en caso de desconocer de uno de los términos,
   explicar lo que creas que significaría según el flujo que sigue cada proceso así
   mismo si existiera cosas que mejorarías agrégalo como un detalle junto a su ¿Por
   qué?

   <img alt="" src = '/Images/arquitectura.png' height = '282' width="571">

5. Demostrando tus hallazgos – 1 puntos
   Se te ha asignado la tarea de presentar una nueva tecnología que mejorará el
   rendimiento del sistema de órdenes a un grupo de stakeholders que no son
   técnicos. La tecnología implica el uso de microservicios para mejorar la
   escalabilidad y reducir el tiempo de respuesta.
   ¿Cómo explicarías esta nueva tecnología y sus beneficios en términos no técnicos
   para asegurar que todos comprendan su importancia y el impacto que tendrá en el
   negocio?
