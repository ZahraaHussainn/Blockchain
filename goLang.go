package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block
type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {
	//Calculate Hash of a Block
	hashInput := fmt.Sprintf("%v%v", inputBlock.transactions, inputBlock.prevHash)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {

	//insert new block and return head pointer
	newBlock := &Block{
		transactions: transactionsToInsert,
		prevPointer:  chainHead,
		prevHash:     "",
		currentHash:  "",
	}
	if chainHead != nil {
		newBlock.prevHash = chainHead.currentHash
	}
	newBlock.currentHash = CalculateHash(newBlock)
	return newBlock

}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	//change transaction data inside block
	current := chainHead
	for current != nil {
		for i, transaction := range current.transactions {
			if transaction == oldTrans {
				current.transactions[i] = newTrans
				current.currentHash = CalculateHash(current) // Recalculate hash
				return
			}
		}
		current = current.prevPointer
	}
}

func ListBlocks(chainHead *Block) {

	//dispaly the data(transaction) inside all blocks
	current := chainHead
	for current != nil {
		fmt.Println("Transactions:", current.transactions)
		fmt.Println("Current Hash:", current.currentHash)
		fmt.Println("Previous Hash:", current.prevHash)
		fmt.Println("-----------------------------")
		current = current.prevPointer
	}

}

func VerifyChain(chainHead *Block) {

	//check whether "Block chain is compromised" or "Block chain is unchanged"
	current := chainHead
	for current != nil && current.prevPointer != nil {
		if current.prevHash != current.prevPointer.currentHash {
			fmt.Println("Blockchain is compromised")
			return
		}
		if CalculateHash(current.prevPointer) != current.prevHash {
			fmt.Println("Blockchain is compromised")
			return
		}
		current = current.prevPointer
	}
	fmt.Println("Blockchain is unchanged")

}

func main() {
	var blockchainHead *Block

	// block1
	blockchainHead = InsertBlock([]string{"Transaction A"}, blockchainHead)

	// block2
	blockchainHead = InsertBlock([]string{"Transaction B"}, blockchainHead)

	// block3
	blockchainHead = InsertBlock([]string{"Transaction C"}, blockchainHead)

	fmt.Println("--- Blockchain Initial State ---")
	ListBlocks(blockchainHead)

	fmt.Println("--- Verifying Blockchain ---")
	VerifyChain(blockchainHead)

	fmt.Println("--- Tampering with Blockchain ---")
	ChangeBlock("Transaction B", "Transaction B - MODIFIED", blockchainHead)

	fmt.Println("--- Blockchain After Tampering ---")
	ListBlocks(blockchainHead)

	fmt.Println("--- Verifying Blockchain After Tampering ---")
	VerifyChain(blockchainHead)
}
