package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts an Int64 to a byte slice
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
