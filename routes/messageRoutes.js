const { createMessage } = require("../controllers/messageController");

const router = require("express").Router();

router.route("/").post(createMessage);

module.exports = router;
