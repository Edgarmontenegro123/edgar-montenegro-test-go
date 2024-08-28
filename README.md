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

   Análisis del problema:
   Dado un arreglo de enteros, necesitamos eliminar los elementos consecutivos cuya suma sea = 0, por
   último devolvemos lo que quedó del array original.

   Solución planteada:
   1) Verificamos si el arreglo entregado es nulo o vacío y devolvemos un array vacío.
   2) Iteramos sobre el array dando seguimiento sobre la suma de los subarray consecutivos.
   3) Sumamos los elementos y vamos almacenando sumas parciales con el fin de comparar el resultado.
   4) Eliminamos los elementos de los subarray que sumen 0.
   5) Repetimos el proceso con el fin de eliminar todas las opciones posibles.
   6) Retornamos el resultado del array final (puede ser un array vacío).


4. Aprendizaje – 1 puntos
   Se te ha asignado poder entender esta arquitectura para poder aplicarlo a un
   nuevo proyecto, se necesita que expliques a detalle cual es la función de cada
   término mostrado en la imagen, en caso de desconocer de uno de los términos,
   explicar lo que creas que significaría según el flujo que sigue cada proceso así
   mismo si existiera cosas que mejorarías agrégalo como un detalle junto a su ¿Por
   qué?

<img alt="" src = '/Images/arquitectura.png' height = '282' width="571">

   Respuesta:
   Según la arquitectura que muestra la imagen podemos ver:
   1) Client Apps: Se refiere a 2 interfaces de una aplicación, una para la versión desktop (web) y
   la otra para aplicaciones móviles (mobile). Serían las interfaces que utilizan nuestros usuarios, también
   conocidos como puntos de entrada para el sistema.
   2) API Gateway: Sería la API que actúa como intermediario y gestiona las solicitudes de parte del usuario.
   En este punto podrían manejarse funciones como autenticación, autorizaciones y enrutamiento a los
   diferentes microservicios.
   3) Microservicios: En este punto podemos observar 4 microservicios diferentes, Catalog, Shopping Cart,
   Discount y Ordering, por sus nombres podemos deducir que el aplicativo maneja 4 microservicios diferentes
   que como su nombre indica tienen funciones diferentes, uno se encargaría de lo que es el catálogo de productos,
   el segundo se encarga de lo que sería el carrito de compras, el siguiente se encarga de aplicar descuentos y el
   último se encargaría de procesar la orden, la clave en este punto es la separación de responsabilidades.
   4) Bases de Datos: En esta parte de la imagen podemos entender que cada microservicio tiene su propia DB,
   lo que permite que las operaciones sean independientes y seguramente facilitan su escalabilidad, además de esta
   manera podemos asegurarnos de que si un microservicio falla, no afectará a los otros microservicios.
   5) Por último Message Broker: Como su nombre lo indica, esto sería una especie de intermediario de mensajes,
   el cual facilita la comunicación entre los diferentes servicios (normalmente de manera asíncrona para permitir
   la mejora en la capacidad de respuesta), además puede gestionar la distribución de mensajes, garantizando que
   cada microservicio reciba los datos que necesita procesar.

5. Demostrando tus hallazgos – 1 puntos
   Se te ha asignado la tarea de presentar una nueva tecnología que mejorará el
   rendimiento del sistema de órdenes a un grupo de stakeholders que no son
   técnicos. La tecnología implica el uso de microservicios para mejorar la
   escalabilidad y reducir el tiempo de respuesta.
   ¿Cómo explicarías esta nueva tecnología y sus beneficios en términos no técnicos
   para asegurar que todos comprendan su importancia y el impacto que tendrá en el
   negocio?

   Respuesta:
   Además de mi experiencia como desarrollador web, tuve la oportunidad de trabajar por casi 10 años en la industria
   gastronómica, y al ser un servicio que en su gran mayoría todos hemos tenido la oportunidad de experimentar,
   creo que puede funcionar como una analogía fácil de entender. Imaginemos nuestro sistema de órdenes actual
   como un restaurante de gran tamaño, el cual posee una sola cocina que se encarga de manejar todas las
   órdenes de todos los clientes. Mientras sea un día tranquilo cómo suelen ser los martes y miércoles,
   todo funciona perfectamente. Pero cuando hablamos de un fin de semana, donde la cantidad de clientes es mayor
   y los pedidos aumentan significativamente, generalmente los pedidos se atrasan, generan demora y los clientes
   tienden a molestarse por la espera. A partir de esta situación, lo que se propone es dividir esa única cocina,
   en "cocinas más pequeñas", donde cada una se especializa en un tipo de plato!. Esas "cocinas más pequeñas" serían
   lo que en la industria IT conocemos como microservicios. Con esa nueva estructura, cuando llegan esos momentos de
   recibir muchísimas órdenes, cada "cocina" puede trabajar de manera independiente e incluso de manera más rápida.
   Eso significa que podríamos recibir o procesar una mayor cantidad de pedidos y a su vez reducir los tiempos
   de espera de nuestros queridos clientes!. Además en caso de que alguna de las "pequeñas cocinas" tenga algún
   problema, al dividirnos de esta manera, aseguramos que las otras "pequeñas cocinas" puedan seguir funcionando
   con normalidad. Al adoptar esta nueva "tecnología", aseguramos que nuestro negocio esté cada vez mejor preparado
   para manejar el crecimiento y que nuestros clientes reciban un mejor servicio.
   En resumen, esta tecnología nos va a ayudar a mejorar la eficiencia operativa y la satisfacción de nuestros usuarios,
   mientras que nuestro negocio podría crecer de manera rápida y segura!.


# Instrucciones para levantar el proyecto:

   * Clonar el proyecto
   * Desde el directorio raíz instalar dependencias: go mod tidy
   * Ejecutar el proyecto: go run main.go