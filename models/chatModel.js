const mongoose = require("mongoose");

const chatSchema = mongoose.Schema(
  {
    members: [
      { type: mongoose.Schema.Types.ObjectId, ref: "User", required: true },
    ],
    groupAdmin: { type: mongoose.Schema.Types.ObjectId, ref: "User" },
    isGroup: { type: Boolean },
    lastMessage: { type: Date, default: Date.now() },
    name: String,
    groupPicture: String,
    messagesDeletedAt: {
      type: [
        {
          userId: {
            type: mongoose.Schema.Types.ObjectId,
            required: true,
            ref: "User",
          },
          date: {
            type: Date,
            default: Date.now(),
          },
        },
      ],
      required: true,
    },
  },
  { timestamps: true }
);

module.exports = mongoose.model("Chat", chatSchema);
