package blockchain

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
	"encoding/json"
)

func NewBlock() *Block {
	return &Block{
		ID:        generateHash(),
		Timestamp: time.Now(),
	}
}

type Block struct {
	// Unique hash for the block
	ID string `json:"id"`

	// The date when the block was created
	Timestamp time.Time `json:"timestamp"`
}

func (b *Block) MarshalJSON() ([]byte, error) {
	type Plain Block
	return json.Marshal(&struct {
		*Plain
		Timestamp string `json:"timestamp"`
	}{
		Plain: (*Plain)(b),
		Timestamp:   b.Timestamp.UTC().Format(time.ANSIC),
	})
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
