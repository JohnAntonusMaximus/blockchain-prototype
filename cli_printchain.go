package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {
	bc := NewBlockchain()
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============== Block %x ============\n", block.Hash)
		fmt.Println()
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("Proof Of Work: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
