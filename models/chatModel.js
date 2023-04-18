const mongoose = require("mongoose");

const chatSchema = mongoose.Schema({
    members: [
      { type: mongoose.Schema.Types.ObjectId, ref: "User", required: true },
    ],
    senderId: {
      type: mongoose.Schema.Types.ObjectId,
      ref: "User",
      required: true,
    },
    message: { type: String, required: true },
    groupAdmin: { type: mongoose.Schema.Types.ObjectId, ref: "User" },
  },
  { timestamps: true }
);

module.exports = mongoose.model("Chat", chatSchema);
