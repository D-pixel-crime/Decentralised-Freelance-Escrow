export const factoryAbi = [
  {
    "type": "function",
    "name": "createEscrow",
    "inputs": [
      {
        "name": "_jobId",
        "type": "uint256",
        "internalType": "uint256"
      },
      {
        "name": "_client",
        "type": "address",
        "internalType": "address"
      },
      {
        "name": "_freelancer",
        "type": "address",
        "internalType": "address"
      },
      {
        "name": "_arbitrator",
        "type": "address",
        "internalType": "address"
      },
      {
        "name": "_confirmationPeriod",
        "type": "uint256",
        "internalType": "uint256"
      }
    ],
    "outputs": [
      {
        "name": "",
        "type": "address",
        "internalType": "address"
      }
    ],
    "stateMutability": "nonpayable"
  },
  {
    "type": "event",
    "name": "EscrowCreated",
    "inputs": [
      {
        "name": "jobId",
        "type": "uint256",
        "indexed": true,
        "internalType": "uint256"
      },
      {
        "name": "escrowAddress",
        "type": "address",
        "indexed": false,
        "internalType": "address"
      },
      {
        "name": "client",
        "type": "address",
        "indexed": true,
        "internalType": "address"
      },
      {
        "name": "freelancer",
        "type": "address",
        "indexed": true,
        "internalType": "address"
      }
    ],
    "anonymous": false
  }
] as const;
