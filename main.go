package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	bc := NewBlockChain()

	for i := 1; i < 3; i++ {
		timestamp := time.Now().Unix()
		h := sha1.New()
		h.Write([]byte(string(timestamp)))
		wallet := hex.EncodeToString(h.Sum(nil))
		rand.Seed(time.Now().UTC().UnixNano())
		bc.AddBlock(strconv.Itoa(i*rand.Intn(100)) + " BYT sent to wallet " + wallet)
	}

	for i, block := range bc.blocks {
		pow := NewProofOfWork(block)
		fmt.Println()
		if i > 0 {
			fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		} else {
			fmt.Println("----------- Current Blockchain ----------")
			fmt.Println()
		}
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
