const User = require("../models/userModel");
const validator = require("validator");
const bcrypt = require("bcrypt");
const { default: mongoose } = require("mongoose");

const signup = async (req, res) => {
  const { firstName, lastName, email, password } = req.body;
  if (!firstName || !lastName || !email || !password)
    return res.status(400).json({ error: "Missing some required data" });
  if (!validator.isEmail(email))
    return res.status(400).json({ error: "Invalid email" });
  if (!validator.isStrongPassword(password))
    return res.status(400).json({ error: "Not a strong password" });

  try {
    const userExists = await User.findOne({ email });
    if (userExists)
      return res.status(400).json({ error: "Email already exists" });
    const user = await User.create({ firstName, lastName, email, password });
    const token = user.generateJWTToken();
    res.status(201).json({ user, token });
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

const login = async (req, res) => {
  const { email, password } = req.body;
  if (!email || !password)
    return res.status(400).json({ error: "Missing some required data" });
  if (!validator.isEmail(email))
    return res.status(400).json({ error: "Invalid email" });

  try {
    const user = await User.findOne({ email });
    if (!user) return res.status(400).json({ error: "No such user" });
    const validPassword = await bcrypt.compare(password, user.password);
    if (!validPassword)
      return res.status(400).json({ error: "Incorrect password" });
    const token = user.generateJWTToken();
    res.status(201).json({ user, token });
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

const searchUsers = async (req, res) => {
  const { name } = req.query;
  try {
    const users = await User.find({
      $and: [
        {
          $or: [
            { firstName: { $regex: `^${name}`, $options: "i" } },
            { lastName: { $regex: `^${name}`, $options: "i" } },
            {
              $expr: {
                $regexMatch: {
                  input: { $concat: ["$firstName", " ", "$lastName"] },
                  regex: new RegExp(`^${name}`, "i"),
                },
              },
            },
          ],
        },
        {
          _id: { $ne: req.user._id }
        }
      ],
    });
    return res.status(200).json({ users });
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

const getUser = async (req, res) => {
  const { userId } = req.params;

  if (!mongoose.Types.ObjectId.isValid(userId))
    return res.status(400).json({ error: "Invalid Id" });

  try {
    const user = await User.findOne({ _id: userId });
    const {password, ...otherFields} = user._doc;
    res.status(200).json({user: otherFields});
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

module.exports = { signup, login, searchUsers, getUser };