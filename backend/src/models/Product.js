const mongoose = require('mongoose');

const productSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,
        trim: true
    },
    description: {
        type: String,
        required: true
    },
    price: {
        type: Number,
        required: true,
        min: 0
    },
    category: {
        type: String,
        required: true,
        enum: ['pescado', 'accesorios', 'ropa', 'decoracion', 'otros']
    },
    stock: {
        type: Number,
        required: true,
        min: 0,
        default: 0
    },
    image: {
        type: String,
        default: 'default-product.png'
    }
}, {
    timestamps: true
});

module.exports = mongoose.model('Product', productSchema);
