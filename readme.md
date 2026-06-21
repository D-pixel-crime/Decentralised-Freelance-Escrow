# Technical Specification & Project Vision: Decentralised Freelance Platform

## 1. The Project Vision (The Whole Idea)
The project is a **Decentralised Freelance Hub & Escrow Platform** designed to replace traditional, high-fee intermediary platforms (like Upwork or Fiverr) with a trustless, cryptographic alternative. 

At its core, the platform eliminates the need for a centralized middleman to hold funds or resolve arguments. Instead, peer-to-peer freelance agreements are governed by smart contracts and a specialized escrow system. 

### Core Workflow:
1. **Agreement & Locking:** A client and freelancer agree on project terms. The client initiates the contract by depositing the project funds (in ETH or stablecoins) directly into a secure, on-chain escrow smart contract.
2. **Execution & Proof:** The freelancer completes the work and submits the deliverables. The associated metadata and proof-of-work files are cryptographically anchored to decentralized storage.
3. **Settlement:** Once the client approves the deliverables, the smart contract automatically releases the locked funds directly to the freelancer's wallet. 
4. **Dispute Resolution:** If a conflict arises, a decentralized or multi-party arbitration layer is triggered, halting the funds until evidence is evaluated and a fair resolution is executed programmatically.

### Visual Identity & Theme:
To counter the psychological anxiety of users locking up substantial funds, the platform rejects playful web aesthetics in favor of an **"Institutional-Grade DeFi / Digital Vault"** theme. Inspired by the ChainFund layout, it utilizes a sleek, dark navy palette with a sharp, vibrant blue central spotlight, clean geometric typography, and high-contrast status states to project ultimate precision, safety, and security.

---

## 2. Total Tech Stack (Current & Planned)

The architecture is split into a modern, high-performance web frontend, a robust systems-level off-chain backend, and a decentralized blockchain infrastructure layer.

### Frontend Layer
*   **Core Framework:** Next.js 15+ (utilizing the App Router paradigm for optimal Server/Client component distribution).
*   **Language:** TypeScript (ensuring strict type safety across contract interactions and API payloads).
*   **Package Manager:** `pnpm` (selected for fast, deterministic, space-efficient dependency management).
*   **Styling Engine:** Tailwind CSS (for rapid utility-first UI building).
*   **Component Foundation:** `shadcn/ui` (built on Radix UI primitives to establish a premium, accessible design system).
*   **Code Quality & Formatting:** ESLint (Flat Config) combined with Prettier (`prettier-plugin-tailwindcss` for automated class sorting).

### Web3 & Blockchain Integration Layer
*   **Wallet Orchestration:** `@rainbow-me/rainbowkit` (provides the premium, customizable "Connect Wallet" interface supporting multiple browser wallets and mobile connections).
*   **Blockchain Hooks:** `wagmi` (react hooks for reading/writing to smart contracts, tracking transaction states, and handling balance lookups).
*   **Low-Level Client:** `viem` (a lightweight, ultra-fast, strictly-typed alternative to legacy libraries like ethers.js).
*   **Asynchronous State:** `@tanstack/react-query` (handles caching, fetching, and data synchronization for on-chain events).

### Backend Layer (Go / Golang)
The off-chain architecture relies on a custom **Go backend** to act as a highly performant, type-safe service layer. It acts as the bridge between the live blockchain and the user interface, saving users massive amounts of gas fees by processing heavy logic off-chain.
*   **Event Indexing & Listening:** The Go service monitors the blockchain network for smart contract events (e.g., `FundsDeposited`, `DisputeRaised`, `FundsReleased`) to keep the application responsive.
*   **API Architecture:** A lightweight Go web framework (such as Gin or Chi) to handle fast off-chain queries like user profiles, job listings, and metadata management.
*   **Real-Time Sync:** WebSockets implemented in Go to push instant updates (like chat messages, status updates, or contract confirmations) to the Next.js frontend.
*   **Web3 Authentication:** Verification of cryptographic signatures on the backend (e.g., Sign-In with Ethereum / SIWE) to securely authenticate users without traditional passwords.

### Storage & Infrastructure Layer
*   **Smart Contracts:** Solidity-based smart contracts deployed to an EVM-compatible blockchain network to handle the trustless movement of funds.
*   **Decentralized Asset Storage:** IPFS (InterPlanetary File System) to securely and immutably store large file deliverables, project descriptions, and legal evidence without overloading the blockchain or relying on a centralized database server.