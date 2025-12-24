const mongoose = require('mongoose');

const orderSchema = new mongoose.Schema({
    customerName: String,
    address: String,
    productName: String, // Simplificado para el ejercicio
    status: { type: String, default: 'Pendiente' },
    date: { type: Date, default: Date.now }
});

module.exports = mongoose.model('Order', orderSchema);
