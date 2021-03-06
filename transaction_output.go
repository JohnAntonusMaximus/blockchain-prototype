package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

// TxOutput represents a transaction output
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

// Lock signs the output
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks if the output can be used by the owner of thep ubkey
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTxOutput creates a new TxOutput
func NewTxOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

// TxOutputs collects TxOutput
type TxOutputs struct {
	Outputs []TxOutput
}

// Serialize serializes TxOutputs
func (outs TxOutputs) Serialize() []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DeserializeOutputs deserializes TxOutputs
func DeserializeOutputs(data []byte) TxOutputs {
	var outputs TxOutputs

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
