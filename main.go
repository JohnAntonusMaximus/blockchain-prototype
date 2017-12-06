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

	for i := 0; i < 1000; i++ {
		timestamp := time.Now().Unix()
		h := sha1.New()
		h.Write([]byte(string(timestamp)))
		wallet := hex.EncodeToString(h.Sum(nil))
		bc.AddBlock(strconv.Itoa(i*rand.Intn(100)) + " BYT sent to wallet " + wallet)
		time.Sleep(11 * time.Millisecond)
	}

	for _, block := range bc.blocks {
		fmt.Println()
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %d", block.Timestamp)
		fmt.Println()
	}

}
