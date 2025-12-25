const express = require('express');
const Product = require('../models/Product');
const auth = require('../middleware/auth');

function productController() {
    const router = express.Router();

    router.use(auth);

    router.get('/', async (req, res) => {
        const products = await Product.find();
        res.render('products', { products });
    });

    router.get('/new', (req, res) => {
        res.render('product-form', { product: {} });
    });

    router.post('/', async (req, res) => {
        // Agregar .png automáticamente si el usuario escribió algo en image
        if (req.body.image && req.body.image.trim() !== '') {
            req.body.image = req.body.image + '.png';
        }
        await Product.create(req.body);
        res.redirect('/products');
    });

    router.get('/edit/:id', async (req, res) => {
        const product = await Product.findById(req.params.id);
        res.render('product-form', { product });
    });

    router.put('/:id', async (req, res) => {
        // Agregar .png automáticamente si el usuario escribió algo en image
        if (req.body.image && req.body.image.trim() !== '') {
            req.body.image = req.body.image + '.png';
        }
        await Product.findByIdAndUpdate(req.params.id, req.body);
        res.redirect('/products');
    });

    router.delete('/:id', async (req, res) => {
        await Product.findByIdAndDelete(req.params.id);
        res.redirect('/products');
    });

    return router;
}

module.exports = productController;
