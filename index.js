const express = require("express");
const cors = require("cors");
require("dotenv").config();
const userRoutes = require("./routes/userRoutes");
const chatRoutes = require("./routes/chatRoutes");

const app = express();

app.use(cors());
app.use(express.json());

app.use("/user", userRoutes);
app.use("/chat", chatRoutes);

require("./db")().then(() => {
    app.listen(process.env.PORT, () => console.log("the server is up"));
});