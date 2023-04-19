const {
  createChat,
  deleteChat,
  getChats,
  updateGroup,
} = require("../controllers/chatController");

const router = require("express").Router();

router.route("/").get(getChats).post(createChat);
router.route("/:chatId").delete(deleteChat).put(updateGroup);

module.exports = router;
