
# 🔗 Blockchain Projects 

This repository showcases my foundational and applied work in Blockchain Development, including a custom blockchain written in **Go (Golang)** and two full-stack **Decentralized Applications (DApps)**: one for **pet adoption** and another for **voting/elections** using smart contracts on Ethereum.

---

## 📁 Contents

- `golang-blockchain/` – Custom blockchain implementation in Go
- `petshop-dapp/` – DApp for decentralized pet adoption
- `election-dapp/` – DApp for conducting secure online elections
- `README.md` – Project overview and documentation

---

## ⚙️ Golang Blockchain Implementation

This section includes a self-built blockchain system developed in **Go** to understand core blockchain mechanisms.

### 🧱 Features

- **Create Blocks**: New blocks are appended with hash pointers.
- **Insert Blocks**: Adds a block dynamically to the chain.
- **Delete Blocks**: Ability to remove a block (for learning, not used in real-world chains).
- **Change Block Data**: Edits simulate tampering to observe hash inconsistency.
- **Hash Calculation**: Each block’s hash is calculated based on its data + previous hash.
- **Blockchain Integrity**: Any data change reflects in broken hashes, simulating tamper detection.

### 🛠 Technologies
- Language: **Go (Golang)**
- Crypto: `sha256` for hashing
- Data Structures: Slices for chain and structs for blocks

---

## 🐶 Pet Adoption DApp

A full-stack **Ethereum-based DApp** built to allow users to adopt pets and store adoption records on the blockchain.

### 🔍 Features

- Displays a list of adoptable pets
- Users can “adopt” a pet using MetaMask wallet
- Prevents double adoption
- Records are stored in smart contracts on the Ethereum testnet

### 📦 Tech Stack

- **Frontend**: React.js
- **Blockchain**: Solidity smart contract
- **Tools**: Truffle, Ganache, MetaMask, Web3.js

---

## 🗳️ Election Voting DApp

A secure blockchain-based voting system to ensure transparency and eliminate vote tampering.

### 🔍 Features

- Register candidates via smart contract
- Voters can vote once only
- Shows live voting results
- Ensures vote privacy and immutability

### 📦 Tech Stack

- **Frontend**: HTML/CSS + JavaScript (or React, if integrated)
- **Smart Contract**: Solidity
- **Blockchain Tools**: Ganache (local blockchain), Web3.js, Truffle

---

## 🚀 How to Run the DApps

### Prerequisites

- Node.js and npm
- Truffle and Ganache
- MetaMask extension
- Go (for Golang blockchain)

### Running the DApps

```bash
# Clone the repo
git clone https://github.com/your-username/blockchain-projects.git
cd blockchain-projects

# For PetShop DApp
cd petshop-dapp
truffle compile
truffle migrate
npm start

# For Election DApp
cd election-dapp
truffle compile
truffle migrate
npm start

# For Go Blockchain
cd golang-blockchain
go run main.go
