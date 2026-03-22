require("dotenv").config();
const { ethers } = require("ethers");

const provider = new ethers.WebSocketProvider(process.env.RPC_WS);

const contractAddress = process.env.CONTRACT;

const abi = [
  "event Transfer(address indexed from, address indexed to, uint256 value)"
];

const contract = new ethers.Contract(contractAddress, abi, provider);

console.log("Listening for Transfer events...");

contract.on("Transfer", (from, to, value, event) => {

    console.log("Transfer Event");

    console.log({
        from,
        to,
        value: value.toString(),
        txHash: event.log.transactionHash,
        block: event.log.blockNumber
    });

});