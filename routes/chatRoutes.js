const { createChat, deleteChat } = require('../controllers/chatController');
const authMiddleware = require('../middlewares/authMiddleware');

const router = require('express').Router();

router.post("/", authMiddleware, createChat);
router.route("/:chatId").delete(authMiddleware, deleteChat);

module.exports = router;