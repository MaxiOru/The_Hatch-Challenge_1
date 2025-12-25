const mongoose = require('mongoose');

const orderSchema = new mongoose.Schema({
    customerName: { type: String, required: true },
    address: String,
    productName: String, // Simplificado para el ejercicio
    status: { 
        type: String, 
        enum: ['Pendiente', 'Entregado', 'Cancelado'],
        default: 'Pendiente' 
    },
    date: { type: Date, default: Date.now }
});

module.exports = mongoose.model('Order', orderSchema);
