require("dotenv").config();
const { ethers } = require("ethers");

const provider = new ethers.WebSocketProvider(process.env.RPC_WS);

provider.getBlockNumber().then((block) => {
  console.log("Connected to Ethereum");
  console.log("Current block:", block);
});