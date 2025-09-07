package main

import (
	"sync"
)

type Storage struct {
	m  map[int]int
	mu sync.Mutex
}

func NewStorage() *Storage {
	s := &Storage{
		m: make(map[int]int),
	}

	return s
}

func main() {
	s := NewStorage()
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				s.mu.Lock()
				s.m[n]++
				s.mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
}
