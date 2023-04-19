const express = require("express");
const cors = require("cors");
require("dotenv").config();
const userRoutes = require("./routes/userRoutes");
const chatRoutes = require("./routes/chatRoutes");
const messageRoutes = require("./routes/messageRoutes");
const authMiddleware = require("./middlewares/authMiddleware");

const app = express();

app.use(cors());
app.use(express.json());

app.use("/user", userRoutes);
app.use("/chat", authMiddleware, chatRoutes);
app.use("/message", authMiddleware, messageRoutes);

require("./db")().then(() => {
  app.listen(process.env.PORT, () => console.log("the server is up"));
});
