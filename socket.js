const { Server } = require("socket.io");

const initSocket = (server) => {
  const io = new Server(server, {
    cors: {
      origin: "*",
    },
  });

  io.on("connection", (socket) => {
    socket.on("joinRoom", ({chatId}) => {
      socket.join(chatId);
    });

    socket.on("publishMessage", ({chatId}) => {
      io.to(chatId).emit("receiveMessage", {chatId});
    });
  });
};

module.exports = { initSocket };
