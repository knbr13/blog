const { createChat } = require('../controllers/chatController');
const authMiddleware = require('../middlewares/authMiddleware');

const router = require('express').Router();

router.post("/", authMiddleware, createChat);

module.exports = router;