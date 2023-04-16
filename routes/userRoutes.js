const { signup } = require("../controllers/userController");

const router = require("express").Router();

router.post("/signup", signup);

module.exports = router;