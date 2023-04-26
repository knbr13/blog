const { Server } = require("socket.io");

const initSocket = (server) => {
  const io = new Server(server, {
    cors: {
      origin: "*",
    },
  });

  io.on("connection", (socket) => {
    socket.on("joinRoom", ({ chatId }) => {
      socket.join(chatId);
    });

    socket.on("publishMessage", ({ chatId, name }) => {
      console.log(name);
      io.to(chatId).emit("receiveMessage", { chatId });
      socket.to(chatId).emit("notification", name);
    });
  });
};

module.exports = { initSocket };
