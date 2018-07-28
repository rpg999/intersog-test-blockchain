package blockchain

import (
	"sort"
	"sync"
)

var DB = &database{
	table: make(map[uint64]*Chain),
	mu:    sync.RWMutex{},
}

// In-memory imitation of a database
type database struct {
	table  map[uint64]*Chain
	mu     sync.RWMutex
	lastId uint64
}

// Get returns a blockchain with a specific id,
// If the blockchain is not exists return nil, false
func (db *database) Get(id uint64) (*Chain, bool) {
	db.mu.RLock()
	bl, ok := db.table[id]
	db.mu.RUnlock()

	return bl, ok
}

// GetAll returns all entries from the database
func (db *database) GetAll() []*Chain {
	var chainOrder []uint64
	var chains []*Chain
	db.mu.RLock()
	defer db.mu.RUnlock()

	for i := range db.table {
		chainOrder = append(chainOrder, i)
	}

	sort.Slice(chainOrder, func(i, j int) bool {
		return chainOrder[i] > chainOrder[j]
	})

	for _, id := range chainOrder {
		chains = append(chains, db.table[id])
	}

	return chains
}

// Add add a new blockchain to the database and returns an ID of it
func (db *database) Add(c *Chain) uint64 {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.lastId += 1
	c.ID = db.lastId
	db.table[db.lastId] = c

	return db.lastId
}

// Delete delete an entry with the specific id from the database
func (db *database) Delete(id uint64) {
	db.mu.Lock()
	delete(db.table, id)
	db.mu.Unlock()
}
