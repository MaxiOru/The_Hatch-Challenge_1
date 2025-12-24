const express = require('express');
const Order = require('../models/Order');
const auth = require('../middleware/auth');

function orderController() {
    const router = express.Router();

    router.get('/', auth, async (req, res) => {
        const orders = await Order.find().sort({ date: -1 });
        res.render('orders', { orders });
    });

    return router;
}

module.exports = orderController;
