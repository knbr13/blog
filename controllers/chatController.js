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

// I think that I should sort the chats based-on the most recent messages.

// This is how the query will look like.

/*
Chat.find({ members: userId })
  .populate({
    path: 'messages',
    options: { sort: { createdAt: -1 }, limit: 1 }
  })
  .sort({ 'messages.createdAt': -1 })
*/

module.exports = { createChat };
