const mongoose = require("mongoose");

const chatSchema = mongoose.Schema({
    members: [
      { type: mongoose.Schema.Types.ObjectId, ref: "User", required: true },
    ],
    groupAdmin: { type: mongoose.Schema.Types.ObjectId, ref: "User" },
    isGroup: { type: Boolean }
  },
  { timestamps: true }
);

module.exports = mongoose.model("Chat", chatSchema);
