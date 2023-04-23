const { default: mongoose } = require("mongoose");

const moreThanTwoMembers = (req, res, members) => {
  const uniqueMembers = new Set(members);
  if (
    !uniqueMembers ||
    !uniqueMembers.size ||
    (uniqueMembers.has(req.user._id) && uniqueMembers.size === 1)
  )
    return res
      .status(400)
      .json({ error: "The group should contains at least two members" });
  const validSet = new Set(
    [...uniqueMembers].filter((element) =>
      mongoose.Types.ObjectId.isValid(element)
    )
  );
  if (validSet.size !== uniqueMembers.size)
    return res
      .status(400)
      .json({ error: "The group should contains at least two members" });
  if (!uniqueMembers.has(req.user._id)) uniqueMembers.add(req.user._id);
  return uniqueMembers;
};

module.exports = { moreThanTwoMembers };
