const mongoose = require("mongoose");

const chatSchema = mongoose.Schema(
  {
    members: [
      { type: mongoose.Schema.Types.ObjectId, ref: "User", required: true },
    ],
    groupAdmin: { type: mongoose.Schema.Types.ObjectId, ref: "User" },
    isGroup: { type: Boolean },
    lastMessage: { type: Date, default: new Date() },
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
            default: new Date(2002, 5, 5, 0, 0, 0),
          },
        },
      ],
      required: true,
    },
  },
  { timestamps: true }
);

module.exports = mongoose.model("Chat", chatSchema);
