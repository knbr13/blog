const {
  createChat,
  deleteChat,
  getChats,
  updateGroup,
  getChat,
  clearChat,
} = require("../controllers/chatController");

const router = require("express").Router();

router.route("/").get(getChats).post(createChat);
router.route("/:chatId").get(getChat).delete(deleteChat).put(updateGroup);
router.delete("/clear/:chatId", clearChat);
module.exports = router;
