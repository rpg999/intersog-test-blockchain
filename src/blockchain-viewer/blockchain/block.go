package blockchain

import (
	"time"
	"fmt"
	"crypto/rand"
	"log"
)

func NewBlock() *Block {
	return &Block{
		ID: generateHash(),
		Timestamp: time.Now().Unix(),
	}
}

type Block struct {
	// Unique hash for the block
	ID string `json:"id"`

	// The date when the block was created
	Timestamp int64 `json:"timestamp"`
}

// generateHash generate a new hash code for a block
func generateHash() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("error while generating key: %s", err)
	}

	return fmt.Sprintf("%x", key)
}
