package blockchain

import (
	"encoding/json"
	"errors"
	"sync"
)

type Chain struct {
	// Blockchain ID
	ID uint64 `json:"id"`
	// Name of the Blockchain
	Name string `json:"name"`
	// List of blocks that belong to the Blockchain
	Blocks []*Block `json:"blocks"`

	mu sync.RWMutex
}

// AddBlock add a new block to the chain
func (c *Chain) AddBlock()  {
	c.mu.Lock()
	defer c.mu.Unlock()

	block := NewBlock()
	c.Blocks = append(c.Blocks, block)
}

// UnmarshalJSON custom json unmarshaller and validator
func (c *Chain) UnmarshalJSON(b []byte) error {
	type plain Chain
	if err := json.Unmarshal(b, (*plain)(c)); err != nil {
		return err
	}

	if c.Name == "" {
		return errors.New("chain's name must be specified")
	}

	return nil
}
