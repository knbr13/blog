const { Server } = require("socket.io");

const initSocket = (server) => {
  const io = new Server(server, {
    cors: {
      origin: "*",
    },
  });

  io.on("connection", (socket) => {
    socket.on("online", ({ userId }) => {
      socket.join(userId);
    });

    socket.on("newRoom", ({ members }) => {
      members.forEach((user) => io.to(user).emit("fetchNewRooms"));
    });

    socket.on("joinRoom", ({ chatId }) => {
      socket.join(chatId);
    });

    socket.on("deleteGroup", ({ chatId, adminName, chatName }) => {
      socket.to(chatId).emit("groupDeleted", { adminName, chatName, chatId });
    });

    socket.on("updateGroup", ({ members, groupId }) => {
      members.forEach((member) =>
        socket.to(member).emit("groupUpdated", { groupId })
      );
    });

    socket.on("publishMessage", ({ chatId, name }) => {
      io.to(chatId).emit("receiveMessage", { chatId });
      socket.to(chatId).emit("notification", { name, chatId });
    });
  });
};

module.exports = { initSocket };
