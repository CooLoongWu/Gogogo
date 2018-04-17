package main

import (
	"blockChain/chain"
	"fmt"
)

func main() {
	blockChain := chain.NewChain()
	blockChain.AddBlock("this is A, pay 20 to B")
	blockChain.AddBlock("this is C, pay 10 to A")

	for _, block := range blockChain.Blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}
