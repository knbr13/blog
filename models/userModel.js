const mongoose = require("mongoose");
const bcrypt = require("bcrypt");
const jwt = require("jsonwebtoken");

const userSchema = mongoose.Schema(
  {
    firstName: {
      type: String,
      required: true,
    },
    lastName: {
      type: String,
      required: true,
    },
    email: {
      type: String,
      required: true,
      unique: true,
      lowercase: true,
      trim: true,
    },
    password: {
      type: String,
      required: true,
    },
    profilePicture: String,
    about: {
      type: String,
      default: "Available",
    },
  },
  { timestamps: true }
);

userSchema.pre("save", async function (next) {
  if (!this.isModified("password")) return next();
  try {
    const salt = await bcrypt.genSalt(10);
    const hashedPassword = await bcrypt.hash(this.password, salt);
    this.password = hashedPassword;
    next();
  } catch (error) {}
});

userSchema.methods.generateJWTToken = function () {
  const user = this;
  const token = jwt.sign(
    { _id: user._id, email: user.email },
    process.env.JWT_SECRET,
    { expiresIn: "1d" }
  );
  return token;
};

module.exports = mongoose.model("User", userSchema);
