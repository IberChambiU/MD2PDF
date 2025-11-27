# Requerimientos Funcionales

- El sistema debe permitir la carga de archivos en formato .md.

- El sistema debe aceptar archivos mediante multipart/form-data.

- El sistema debe generar un identificador único (UUID) por cada conversión.

- El sistema debe convertir el archivo Markdown a HTML.

- El sistema debe generar un archivo PDF a partir del HTML.

- El sistema debe almacenar temporalmente el archivo PDF generado.

- El sistema debe proveer un endpoint POST /convert para subir y procesar archivos.

- El sistema debe proveer un endpoint GET /download/:id para descargar el PDF.

- El sistema debe retornar errores claros cuando falle la conversión.

- El sistema debe validar que el archivo subido tenga extensión .md.

- El sistema debe eliminar los archivos una vez pasada la ventana de retención configurada.

- El sistema debe proveer un endpoint de salud GET /healthz.

- El sistema debe permitir configurar el tiempo de retención mediante variables de entorno.

- El sistema debe manejar múltiples solicitudes concurrentes de conversión.

- El sistema debe notificar en la respuesta la URL de descarga.

- El sistema debe permitir convertir Markdown con tablas, listas y enlaces.

- El sistema debe soportar imágenes referenciadas dentro del Markdown.

- El sistema debe aceptar tamaños de archivo hasta un límite definido (ej. 10 MB).

- El sistema debe registrar logs de cada conversión realizada.

- El sistema debe devolver en JSON la información de cada job creado.

# Requerimientos No Funcionales

- El sistema debe ejecutarse en contenedores Docker.

- El sistema debe usar un contenedor base liviano (ej. Alpine).

- El binario de Go debe compilarse en una etapa separada (multi-stage build).

- El sistema debe ejecutarse como usuario no root en el contenedor.

- El sistema debe exponer el servicio en el puerto configurable (por defecto 3000).

- El sistema debe estar diseñado para ser escalable horizontalmente.

- El sistema debe iniciar en menos de 3 segundos.

- El sistema debe poder manejar al menos 50 conversiones concurrentes sin caídas.

- El tamaño máximo de cada contenedor no debe superar los 150 MB.

- El sistema debe permitir configuración vía variables de entorno.

- Los errores deben registrarse en logs legibles.

- El sistema debe funcionar en Linux, Windows y Mac (contenedorizado).

- El sistema debe usar memoria de forma eficiente (evitar leaks).

- El sistema debe ser portable entre diferentes orquestadores (Docker, Kubernetes).

- El sistema debe mantener una cobertura básica de pruebas unitarias.

- El sistema debe usar estándares REST en sus endpoints.

- El sistema debe contar con documentación básica de uso (README + ejemplos).

- El sistema debe garantizar que los archivos expirados se eliminen automáticamente.

- El sistema debe ser seguro ante archivos maliciosos (no ejecutar código embebido).

- El sistema debe cumplir con un tiempo de respuesta aceptable (<2s para archivos pequeños).