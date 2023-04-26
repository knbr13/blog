const {
  createChat,
  deleteChat,
  getChats,
  updateGroup,
  getChat
} = require("../controllers/chatController");

const router = require("express").Router();

router.route("/").get(getChats).post(createChat);
router.route("/:chatId").get(getChat).delete(deleteChat).put(updateGroup);

module.exports = router;
