const { default: mongoose } = require('mongoose');
const Chat = require('../models/chatModel');

const createChat = async (req, res) => {
    const {members, isGroup} = req.body;
    if(!members || members.length < 1) return res.status(400).json({error: "Missing some required data"});
    try {
        if(members.some(element => mongoose.Types.ObjectId.isValid(element) === false));
        members.push(req.user._id);
        let chat;
        if(isGroup){
            chat = await Chat.create({members, isGroup, groupAdmin: req.user._id});
            return res.status(201).json(chat);
        }
        chat = await Chat.create({members});
        res.status(201).json(chat);
    } catch (error) {
        return res.status(500).json({ message: "Internal Server Error" });
    }
}


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