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

   Para resolver este problema me gustaría proponer varias soluciones posibles dependiendo del caso,
   incluso podría ser una combinación de diferentes estrategias!

   * La primera solución que pensaría sería refactorizar el código que maneja las solicitudes a la DB para
   que funcione de una manera asíncrona si es que no está implementada dicha solución. Esto mejora la
   capacidad del sistema para seguir aceptando solicitudes mientras se espera la respuesta de la DB.
   * Otra solución es implementar un sistema de memoria cache, en el cual se pueden almacenar los resultados
   de las consultas más frecuentes o de los datos que cambian con poca frecuencia. Esto va a reducir la cantidad
   de solicitudes directas a la DB por un lado y por el otro va a mejorar la velocidad de respuesta de esos
   datos que serían de lectura.
   * También podríamos optimizar las consultas a la DB y utilizar índices, por ejemplo en vez de realizar
   un SELECT * FROM... Podríamos analizar si es necesario traer todos los datos o solo una parte, esto hará
   que el tiempo de ejecución sea menor.
   * Analizar si es posible agregar una escalabilidad bien sea vertical u horizontal, la primera dando más
   recursos como CPU o memoria para el procesamiento de datos y la segunda que permitiría repartir la carga entre
   múltiples bases de datos por ejemplo.
   * Teniendo en cuenta el punto anterior, de ser posible, podríamos separar los datos en 2 ó más bases de datos,
   manejando por ejemplo los datos más críticos en una para que el resto de datos no compitan por recursos.
   * Configurar el sistema para que escale automáticamente en función de la cantidad de peticiones. Esto
   garantiza que nuestro sistema pueda manejar picos de transacciones sin necesidad de intervención manual.

   Seguramente existen muchas soluciones más o menos costosas, pero para empezar con lo descrito anteriormente,
   podemos dar mantenimiento y seguridad de que nuestro sistema no falle debido a un cuello de botella en nuestra DB.
   
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

   Análisis del problema:
   Desarrollar una función que encripte un mensaje específico agregando una "clave" antes de cada vocal en el
   mensaje original.
   Parámetros de entrada: clave y mensaje.
   Consideraciones:
   * Se especifica que se tomarán en cuenta las vocales (minúsculas y mayúsculas), pero no
   se tomarán en cuenta las vocales con acento.
   * Si la clave es nula o vacía, se debe utilizar "DCJ" como valor predeterminado.
   * Si el mensaje es nulo o vacío, se retornará una cadena vacía.
   * La función debe respetar mayúsculas y minúsculas según el mensaje original.

   Solución planteada:
   1) Validar la entrada de la clave y el mensaje.
   2) Iterar sobre cada caracter del mensaje.
   3) Identificar las vocales y anteponer la clave.
   4) Construir la cadena resultante.
   5) Retornar la cadena encriptada.


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


# Instrucciones para levantar el proyecto:

   * Clonar el proyecto
   * Desde el directorio raíz instalar dependencias: go mod tidy
   * Ejecutar el proyecto: go run main.go