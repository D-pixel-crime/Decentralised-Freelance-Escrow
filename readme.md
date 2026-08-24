# NoTrust

A decentralized freelance platform that replaces traditional intermediary services with trustless, on-chain escrow contracts. Clients lock funds into a smart contract at the start of an engagement; funds are released automatically on approval or held for arbitration if a dispute arises.

## How It Works

1. **Agreement** -- A client and freelancer agree on terms. The client deposits ETH into an on-chain escrow contract.
2. **Execution** -- The freelancer completes the work and submits deliverables. File metadata is anchored to IPFS.
3. **Settlement** -- The client approves the deliverables and the contract releases funds directly to the freelancer's wallet.
4. **Disputes** -- If either party raises a dispute, funds are frozen until a resolution is executed programmatically.


## Project Structure

```
contracts/
├── src/
│   ├── FreelanceEscrow.sol    # Core escrow logic (deposits, releases, disputes)
│   └── EscrowFactory.sol      # Factory for deploying per-job escrow instances
├── script/                    # Foundry deploy scripts
└── foundry.toml

backend/
├── main.go                    # Entrypoint (bindings generation, server setup)
├── contracts/                 # Auto-generated Go bindings
├── handlers/                  # HTTP request handlers
├── routers/                   # Route definitions (auth, GET, POST)
├── models/                    # Data models
├── indexer/                   # On-chain event indexer (currently JIT)
└── utils/                     # DB, Redis, Web3 connection helpers

frontend/
├── app/                       # Next.js App Router pages
├── components/                # React components (shadcn/ui based)
├── constants/                 # ABI and contract addresses
├── contexts/                  # React context providers
├── lib/                       # Utility functions
└── types/                     # TypeScript type definitions
```

| Layer | Key Technologies |
|-------|-----------------|
| **Contracts** | Solidity, Foundry, OpenZeppelin |
| **Backend** | Go 1.25, Gin, MongoDB, Redis, go-ethereum (abigen) |
| **Frontend** | Next.js 16, TypeScript, Tailwind CSS, shadcn/ui, RainbowKit, wagmi, viem, SIWE |

## Prerequisites

- [Go](https://go.dev/dl/) >= 1.25
- [Node.js](https://nodejs.org/) >= 20 and [pnpm](https://pnpm.io/)
- [Foundry](https://book.getfoundry.sh/getting-started/installation) (`forge`, `anvil`, `cast`)
- [abigen](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings) (ships with `geth` tools)
- A running MongoDB instance
- A running Redis instance

## Getting Started

### 1. Clone the repository

```bash
git clone --recurse-submodules https://github.com/D-pixel-crime/Freelance_Escrow.git
cd Freelance_Escrow
```

If you already cloned without `--recurse-submodules`:

```bash
git submodule update --init --recursive
```

### 2. Smart Contracts

```bash
cd contracts
forge install        # fetch dependencies (forge-std, openzeppelin)
forge build          # compile contracts
forge test -vvv      # run the test suite
```

To deploy locally with Anvil:

```bash
anvil                                       # starts a local chain on http://127.0.0.1:8545
forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --broadcast
```

### 3. Backend

```bash
cd backend
cp .env.example .env   # then fill in your values (see Environment Variables below)
go mod download
air                    # starts the server with hot-reload on :8080
```

Or without Air:

```bash
go run .
```

> The backend automatically runs `forge build` and `abigen` on startup to regenerate Go bindings from the Solidity source.

### 4. Frontend

```bash
cd frontend
cp .env.example .env   # then fill in your values (see Environment Variables below)
pnpm install
pnpm dev               # starts Next.js dev server on http://localhost:3000
```

## Environment Variables

### Backend (`backend/.env`)

| Variable | Description |
|----------|-------------|
| `PORT` | Server port (default `8080`) |
| `MONGO_CONNECT_URI` | MongoDB connection string |
| `DATABASE_NAME` | MongoDB database name |
| `REDIS_CONNECT_URI` | Redis connection string |
| `JWT_SECRET` | Secret used to sign JWT tokens |
| `WEB3_RPC_URL` | Ethereum JSON-RPC endpoint (e.g. Anvil, Alchemy, Infura) |
| `ESCROW_CONTRACT_ADDRESS` | Deployed `FreelanceEscrow` contract address |
| `FACTORY_ADDRESS` | Deployed `EscrowFactory` contract address |
| `PRIVATE_KEY` | Private key of the deployer/admin wallet |

### Frontend (`frontend/.env`)

| Variable | Description |
|----------|-------------|
| `NEXT_PUBLIC_API_URL` | Backend API base URL |
| `NEXT_PUBLIC_RPC_URL` | Ethereum JSON-RPC endpoint |
| `NEXT_PUBLIC_CONTRACT_ADDRESS` | Deployed contract address |
| `PINATA_JWT` | Pinata JWT for IPFS uploads |

## Makefile Targets

Run these from the repository root:

```bash
make check-contracts   # forge build, fmt --check, test
make check-backend     # go build, go test
make check-frontend    # pnpm lint, pnpm build
make check-all         # all of the above
```
