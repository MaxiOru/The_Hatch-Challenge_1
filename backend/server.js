require('dotenv').config();
const express = require('express');
const mongoose = require('mongoose');
const cookieParser = require('cookie-parser');
const methodOverride = require('method-override');
const path = require('path');

// Controladores
const authController = require('./controllers/authController');
const productController = require('./controllers/productController');
const orderController = require('./controllers/orderController');

const app = express();
const PORT = process.env.PORT || 3000;

// Conexión a MongoDB
mongoose.connect(process.env.MONGODB_URI)
    .then(() => console.log('MongoDB Conectado (Backend)'))
    .catch(err => console.error(err));

// Configuración
app.set('view engine', 'pug');
app.set('views', path.join(__dirname, 'views'));
app.use(express.static(path.join(__dirname, 'public')));
app.use('/uploads', express.static(path.join(__dirname, 'uploads')));
app.use(express.urlencoded({ extended: true }));
app.use(cookieParser());
app.use(methodOverride('_method'));

// Rutas
app.use('/auth', authController());
app.use('/products', productController());
app.use('/orders', orderController());

// Redirección raíz
app.get('/', (req, res) => res.redirect('/products'));

app.listen(PORT, () => {
    console.log(`Panel de Admin corriendo en http://localhost:${PORT}`);
});
