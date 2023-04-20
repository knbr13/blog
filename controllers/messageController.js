const { default: mongoose } = require("mongoose");
const Message = require("../models/messageModel");
const Chat = require("../models/chatModel");

const createMessage = async (req, res) => {
  const { chatId, messageText } = req.body;
  if (!chatId || !messageText)
    return res.status(400).json({ error: "Missing some required data" });
  if (!mongoose.Types.ObjectId.isValid(chatId))
    return res.status(400).json({ error: "Invalid ID" });
  try {
    const chat = await Chat.findById(chatId);
    if(!chat) return res.status(404).json({error: "no such chat"});
    const message = await Message.create({
      senderId: req.user._id,
      chatId,
      messageText,
    });
    await Chat.findByIdAndUpdate(chatId, {lastMessage: Date.now()});
    res.status(201).json(message);
  } catch (error) {
    return res.status(500).json({ error: "Internal Server Error" });
  }
};

const getMessages = async (req, res) => {
  const { chatId } = req.params;
  if (!mongoose.Types.ObjectId.isValid(chatId))
    return res.status(400).json({ error: "Invalid Id" });
  try {
    const messages = await Message.find({ chatId });
    res.status(200).json(messages);
  } catch (error) {
    return res.status(500).json({ error: "Internal Server Error" });
  }
};

module.exports = { createMessage, getMessages };
