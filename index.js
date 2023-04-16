const express = require("express");
const cors = require("cors");
require("dotenv").config();
const userRoutes = require("./routes/userRoutes");

const app = express();

app.use(cors());
app.use(express.json());

app.use("/user", userRoutes);

require("./db")().then(() => {
    app.listen(process.env.PORT, () => console.log("the server is up"));
});