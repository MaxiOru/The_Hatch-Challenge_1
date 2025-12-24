const mongoose = require('mongoose');

const productSchema = new mongoose.Schema({
    name: String,
    description: String,
    price: Number,
    image: String // Opcional según enunciado, pero útil
});

module.exports = mongoose.model('Product', productSchema);
