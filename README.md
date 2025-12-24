# Proyecto: Tienda Online de Larry el Pingüino 🐧

Este proyecto consta de dos partes principales: un Panel de Administración (Node.js) y una Tienda Online (Go). Todo funciona sin JavaScript en el cliente, como les gusta a los pingüinos.

## Requisitos Previos

1.  **MongoDB**: Debe estar instalado y corriendo en el puerto `27017`.
2.  **Node.js**: Para el panel de administración.
3.  **Go**: Para la tienda online.

---

## Configuración de MongoDB

Antes de ejecutar los servidores, **MongoDB debe estar corriendo**. Aquí te explicamos cómo:

### Verificar si MongoDB está instalado

```bash
mongod --version
```

Si ves la versión de MongoDB, está instalado. Si no, descárgalo desde [mongodb.com/try/download/community](https://www.mongodb.com/try/download/community).

### Iniciar MongoDB

#### Opción 1: Como Servicio de Windows (Recomendado)

Si instalaste MongoDB como servicio, inícialo con:

```powershell
net start MongoDB
```

Para verificar que está corriendo:

```powershell
Get-Service -Name MongoDB
```

Deberías ver `Status: Running`.

#### Opción 2: Manualmente

Si prefieres iniciarlo manualmente, abre una terminal **independiente** y ejecuta:

```bash
mongod
```

O con una ruta de datos específica:

```bash
mongod --dbpath "C:\data\db"
```

⚠️ **IMPORTANTE**: Deja esta terminal abierta mientras trabajas con el proyecto. MongoDB se detendrá si cierras la terminal.

### Verificar la Conexión

MongoDB debería estar escuchando en `http://localhost:27017`. Puedes verificarlo abriendo esa URL en tu navegador (verás un mensaje como "It looks like you are trying to access MongoDB over HTTP").

---

## Configuración Inicial

### Variables de Entorno

Antes de ejecutar los servidores, debes configurar las variables de entorno:

**Backend:**
1. Copia el archivo de ejemplo:
   ```bash
   cd backend
   cp .env.example .env
   ```
2. Edita `.env` si necesitas cambiar la configuración (puerto, MongoDB, JWT secret)

**Frontend:**
1. Copia el archivo de ejemplo:
   ```bash
   cd frontend
   cp .env.example .env
   ```
2. Edita `.env` si necesitas cambiar la configuración (puerto, MongoDB)

⚠️ **IMPORTANTE**: Nunca subas los archivos `.env` a GitHub. Ya están incluidos en `.gitignore`

---

## 1. Panel de Administración (Backend)

Aquí Paula gestiona los productos y ve los pedidos.

### Instalación y Ejecución

1.  Abre una terminal y ve a la carpeta `backend`:
    ```bash
    cd backend
    ```
2.  Instala las dependencias:
    ```bash
    npm install
    ```
3.  Asegúrate de haber configurado el archivo `.env` (ver sección de Configuración Inicial)
4.  Inicia el servidor:
    ```bash
    npm start
    ```
5.  Abre tu navegador en: [http://localhost:3000](http://localhost:3000)

### Credenciales de Acceso
El sistema creará automáticamente el usuario la primera vez que intentes entrar con estas credenciales:
*   **Usuario:** `paula`
*   **Contraseña:** `admin123`

**Nota de Seguridad:** Las contraseñas se almacenan hasheadas con bcrypt, cumpliendo con las mejores prácticas de seguridad.

---

## 2. Tienda Online (Frontend)

Aquí los pingüinos compran pescado.

### Instalación y Ejecución

1.  Abre una **nueva** terminal y ve a la carpeta `frontend`:
    ```bash
    cd frontend
    ```
2.  Asegúrate de haber configurado el archivo `.env` (ver sección de Configuración Inicial)
3.  Descarga las dependencias de Go:
    ```bash
    go mod tidy
    ```
4.  Ejecuta el servidor:
    ```bash
    go run main.go
    ```
5.  Abre tu navegador en: [http://localhost:8080](http://localhost:8080)

---

## Características del Sistema

### Imágenes de Productos
*   Los productos pueden tener imágenes asociadas
*   Las imágenes deben colocarse en `frontend/static/images/` con formato PNG
*   Al crear/editar un producto, ingresa el nombre de la imagen (sin extensión)
*   Ejemplo: para `producto_n1.png`, ingresa `producto_n1` en el formulario

### Diseño Visual
*   **CSS Puro:** Sistema completamente estilizado con CSS personalizado, sin dependencias externas
*   **Frontend:** Estilos en `frontend/static/css/styles.css`
*   **Backend:** Estilos en `backend/public/css/admin-styles.css`
*   **Responsive:** Diseño adaptable a dispositivos móviles y tablets
*   **Flexbox:** Layout moderno con productos mostrando imagen a la derecha y descripción a la izquierda

## Notas Técnicas

*   **Arquitectura MVC:** Se ha implementado el patrón Modelo-Vista-Controlador en ambos proyectos para una mejor organización.
    *   **Backend:** Controladores en `/controllers`, Rutas en `/routes`, Modelos en `/models`, Vistas en `/views`.
    *   **Frontend:** Controladores en `/controllers`, Modelos en `/models`, Configuración en `/config`, Templates en `/templates`.
*   **Base de Datos:** Ambos sistemas se conectan a la misma base de datos MongoDB llamada `larry_shop`.
*   **Autenticación:** Se usa JWT. Dado que no se permite JavaScript en el cliente, el token se almacena en una cookie `HttpOnly` segura.
*   **Seguridad:**
    *   Las contraseñas se hashean con **bcrypt** (10 salts rounds)
    *   Tokens JWT con clave secreta configurable en `.env`
    *   Variables sensibles protegidas en archivos `.env` (no se suben a GitHub)
*   **Archivos Estáticos:**
    *   **Frontend (Go):** Servidos desde `frontend/static/` (imágenes y CSS)
    *   **Backend (Node.js):** Servidos desde `backend/public/` (CSS del panel admin)
*   **Cero JS en Cliente:** No hay ni una línea de JavaScript en el frontend ni en el panel de admin (lado cliente). Todo es renderizado en el servidor (SSR).
