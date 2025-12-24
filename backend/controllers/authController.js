const express = require('express');
const User = require('../models/User');
const jwt = require('jsonwebtoken');
const bcrypt = require('bcrypt');

function authController() {
    const router = express.Router();

    router.get('/login', (req, res) => {
        res.render('login');
    });

    router.post('/login', async (req, res) => {
        const { username, password } = req.body;
        
        // Lógica de seed automático simplificada con bcrypt
        let user = await User.findOne({ username });
        if (!user && username === 'paula' && password === 'admin123') {
            const hashedPassword = await bcrypt.hash('admin123', 10);
            user = await User.create({ username: 'paula', password: hashedPassword });
        }

        // Verificar contraseña con bcrypt
        if (user && await bcrypt.compare(password, user.password)) {
            const token = jwt.sign({ _id: user._id }, process.env.JWT_SECRET);
            res.cookie('token', token, { httpOnly: true });
            res.redirect('/products');
        } else {
            res.render('login', { error: 'Credenciales incorrectas' });
        }
    });

    router.get('/logout', (req, res) => {
        res.clearCookie('token');
        res.redirect('/auth/login');
    });

    return router;
}

module.exports = authController;
