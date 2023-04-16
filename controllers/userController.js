const User = require("../models/userModel");
const validator = require("validator");

const signup = async (req, res) => {
    const {firstName, lastName, email, password} = req.body;
    if(!firstName || !lastName || !email || !password) return res.status(400).json({error: "Missing some required data"});
    if(!validator.isEmail(email)) return res.status(400).json({error: "Invalid email"});
    if(!validator.isStrongPassword(password)) return res.status(400).json({error: "Not a strong password"});

    try {
        const userExists = await User.findOne({email});
        if(userExists) return res.status(400).json({error: "Email already exists"});
        const user = await User.create({firstName, lastName, email, password});
        const token = user.generateJWTToken();
        res.status(201).json({user, token});
    } catch (error) {
        res.status(400).json({error: error.message});
    }
}

module.exports = {signup};