const { createMessage, getMessages } = require("../controllers/messageController");

const router = require("express").Router();

router.route("/").post(createMessage).get(getMessages);

module.exports = router;
