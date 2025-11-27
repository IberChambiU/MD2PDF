# MD2PDF

Servicio que convierte documentos Markdown a PDF y los almacena con un nombre UUID único.

## ¿Cómo funciona la conversión?

El servicio convierte archivos Markdown a PDF utilizando Chrome/Chromium headless mediante la librería `chromedp`. El proceso es el siguiente:

1. Recibe el contenido Markdown a través de la API
2. Convierte el Markdown a HTML
3. Usa Chrome/Chromium headless para renderizar el HTML y generar el PDF
4. Guarda el PDF con un UUID único
5. Retorna el UUID para acceder al archivo generado

## Variables de Entorno

### Aplicación Principal

- **`PORT`**: Puerto en el que se ejecutará el servidor. Por defecto: `3030`
- **`URL`**: Dirección URL base del servidor. Por defecto: `http://localhost`
- **`SECURE`**: Habilita la encriptación mediante TLS (`true`/`false`). Por defecto: `false`
- **`URL_SERVIDOR`**: URL del servidor de archivos. Por defecto: `http://localhost:3030`
- **`CHROMEDP_PATH`**: Ruta al ejecutable de Chromium/Chrome. Por defecto: se busca en las rutas comunes del sistema
- **`SIMETRIC_KEY`**: Clave simétrica para encriptación. Por defecto: `ComprobemosSiAlguienPuedeLeerlos`

### Limpieza de Archivos

- **`PDF_CLEANUP_INTERVAL_MINUTES`**: Intervalo en minutos para ejecutar la limpieza de archivos PDF antiguos. Por defecto: `24`
- **`PDF_CLEANUP_FILE_MAX_AGE_MINUTES`**: Tiempo máximo en minutos que un archivo PDF puede existir antes de ser eliminado. Por defecto: `60`

### Autenticación API

- **`API_USER`**: Usuario para autenticación básica de las APIs protegidas. Por defecto: `default_user`
- **`API_PASSWORD`**: Contraseña para autenticación básica de las APIs protegidas. Por defecto: `default_password`

## Licencia

Este proyecto está bajo la **Licencia MIT**, una de las licencias de código abierto más permisivas.

Es de **uso completamente libre** para fines personales y comerciales. Puedes:
- Usar el código libremente
- Modificarlo según tus necesidades
- Distribuirlo
- Incluirlo en proyectos comerciales

## Contribuciones

Este proyecto es de **código abierto** y acepta contribuciones libres. Se agradecen y valoran:
- Comentarios y sugerencias
- Reportes de bugs
- Mejoras y optimizaciones
- Nueva funcionalidad
- Mejoras en la documentación

Si deseas colaborar:

1. Haz un fork del repositorio
2. Crea una rama para tu feature (`git checkout -b feature/nueva-funcionalidad`)
3. Realiza tus cambios y commit (`git commit -m 'Agrega nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

Todas las contribuciones son bienvenidas.
