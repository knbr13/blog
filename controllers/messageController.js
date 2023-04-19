const Message = require("../models/messageModel");

const createMessage = async (req, res) => {
  const { senderId, chatId, messageText } = req.body;
  if (!senderId || !chatId || !messageText)
    return res.status(400).json({ error: "Missing some required data" });
  try {
    const message = await Message.create({ senderId, chatId, messageText });
    res.status(201).json(message);
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

module.exports = { createMessage };
