const { default: mongoose } = require("mongoose");
const Chat = require("../models/chatModel");

const createChat = async (req, res) => {
  const { members, isGroup } = req.body;
  const uniqueMembers = new Set(members);
  if (
    !uniqueMembers ||
    !uniqueMembers.size ||
    (uniqueMembers.has(req.user._id) && uniqueMembers.size === 1)
  )
    return res
      .status(400)
      .json({ error: "The group should contains at least two members" });
  try {
    const validSet = new Set(
      [...uniqueMembers].filter((element) =>
        mongoose.Types.ObjectId.isValid(element)
      )
    );
    if (validSet.size !== uniqueMembers.size)
      return res
        .status(400)
        .json({ error: "The group should contains at least two members" });
    if (!uniqueMembers.has(req.user._id)) uniqueMembers.add(req.user._id);
    let chat;
    if (isGroup) {
      chat = await Chat.create({
        members: [...uniqueMembers],
        isGroup,
        groupAdmin: req.user._id,
      });
      return res.status(201).json(chat);
    }
    chat = await Chat.create({ members: [...uniqueMembers] });
    res.status(201).json(chat);
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

const deleteChat = async (req, res) => {
  const { chatId } = req.params;
  try {
    const deletedChat = await Chat.findByIdAndDelete(chatId);
    if (!deletedChat) return res.status(400).json({ error: "no such chat" });
    res.status(200).json(deletedChat);
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

const getChats = async (req, res) => {
  try {
    const chats = await Chat.find({
      members: { $in: [req.user._id] },
    }).sort({ updatedAt: -1 });
    res.status(200).json(chats);
  } catch (error) {
    return res.status(500).json({ message: "Internal Server Error" });
  }
};

module.exports = { createChat, deleteChat, getChats };
