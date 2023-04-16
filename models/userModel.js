const mongoose = require("mongoose");

const userSchema = mongoose.Schema({
    firstName: {
        type: String,
        required: true
    },
    lastName: {
        type: String,
        required: true
    },
    email: {
        type: String,
        required: true,
        unique: true
    },
    profilePicture: String,
    about: {
        type: String,
        default: "Available"
    }
}, {timestamps: true});

module.exports = mongoose.model("User", userSchema);