const mongoose = require('mongoose');

const userSchema = new mongoose.Schema({
    username: String,
    password: String // En un caso real, esto iría hasheado. Aquí simple.
});

module.exports = mongoose.model('User', userSchema);
