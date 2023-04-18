const { signup, login, searchUsers, getUser, updateProfile } = require("../controllers/userController");
const authMiddleware = require("../middlewares/authMiddleware");
const router = require("express").Router();

router.post("/signup", signup);
router.post("/login", login);
router.get("/search",authMiddleware, searchUsers);
router.get("/:userId", getUser);
router.put("/", authMiddleware, updateProfile);
module.exports = router;