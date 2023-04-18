const { createChat, deleteChat, getChats } = require('../controllers/chatController');

const router = require('express').Router();

router.route("/").get(getChats).post(createChat);
router.route("/:chatId").delete(deleteChat);

module.exports = router;