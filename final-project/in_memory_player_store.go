package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		&sync.Mutex{},
	}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    *sync.Mutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	score := i.store[name]
	return score
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	i.store[name]++
	i.mu.Unlock()
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}
