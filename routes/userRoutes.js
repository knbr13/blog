const { signup, login, searchUsers, getUser } = require("../controllers/userController");
const authMiddleware = require("../middlewares/authMiddleware");
const router = require("express").Router();

router.post("/signup", signup);
router.post("/login", login);
router.get("/search",authMiddleware, searchUsers);
router.get("/:userId", getUser);

module.exports = router;