package storage

import (
	"sync"
)

var _ Storage = (*golangMap)(nil)

// Storage implementation.
type golangMap struct {
	data map[string]any
	mu   sync.Mutex
}

func NewMap() Storage {
	return &golangMap{
		data: make(map[string]any),
	}
}

// Delete sets key-value pair to storage.
func (g *golangMap) Set(key string, value any) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.data[key] = value
}

// Delete gets value from storage by key.
func (g *golangMap) Get(key string) any {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.data[key]
}

// Delete removes key value from storage by key.
func (g *golangMap) Delete(key string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.data, key)
}
