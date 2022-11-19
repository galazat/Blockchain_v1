//package Blockchain

package main

import (
	"fmt"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficalty   int
}

func main() {
	fmt.Println("Hello")
}
