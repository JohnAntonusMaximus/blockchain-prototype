package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BYT to JohnAntonus")
	bc.AddBlock("Send 2 more BYT to JohnAntonus")

	for _, block := range bc.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
