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
    const chatExists = await Chat.findOne({ members: [...uniqueMembers] });
    if (chatExists)
      return res.status(400).json({ error: "This chat is already created" });
    chat = await Chat.create({ members: [...uniqueMembers] });
    res.status(201).json(chat);
  } catch (error) {
    return res.status(500).json({ error: "Internal Server Error" });
  }
};

const deleteChat = async (req, res) => {
  const { chatId } = req.params;
  try {
    const deletedChat = await Chat.findByIdAndDelete(chatId);
    if (!deletedChat) return res.status(400).json({ error: "no such chat" });
    res.status(200).json(deletedChat);
  } catch (error) {
    return res.status(500).json({ error: "Internal Server Error" });
  }
};

const getChats = async (req, res) => {
  try {
    const chats = await Chat.find({
      members: { $in: [req.user._id] },
    })
      .sort({ updatedAt: -1 })
      .populate({
        path: "members",
        select: "firstName lastName profilePicture _id",
        match: { _id: { $ne: req.user._id } },
      });
      
    res.status(200).json(chats);
  } catch (error) {
    return res.status(500).json({ error: "Internal Server Error" });
  }
};

const updateGroup = async (req, res) => {
  const { chatId } = req.params;
  const { members } = req.body;
  try {
    const chat = await Chat.findById(chatId);
    if (!chat.isGroup)
      return res.status(400).json({
        error: "this chat is not a group chat, you can delete the whole chat",
      });
    if (chat.groupAdmin !== req.user._id)
      return res
        .status(401)
        .json({ error: "you don't have the access the add or delete members" });
    const newChat = await Chat.findByIdAndUpdate(
      chatId,
      { members },
      { new: true }
    );
    res.status(200).json(newChat);
  } catch (error) {
    return res.status(500).json({ error: "Internal Server Error" });
  }
};

module.exports = { createChat, deleteChat, getChats, updateGroup };
