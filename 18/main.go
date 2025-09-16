package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Atomic
type AtomicCounter struct {
	value atomic.Int64
}

func (c *AtomicCounter) inc() {
	c.value.Add(1)
}

func (c *AtomicCounter) printValue() {
	fmt.Println(c.value.Load())
}

// Mutex
type MutexCounter struct {
	value int
	mu    sync.Mutex
}

func (c *MutexCounter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *MutexCounter) printValue() {
	fmt.Println(c.value)
}

// go run main.go -race
func main() {
	wg := sync.WaitGroup{}

	// atomic
	ac := &AtomicCounter{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				ac.inc()
			}
		}()
	}

	// mutex
	mc := &MutexCounter{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mc.inc()
			}
		}()
	}

	wg.Wait()
	ac.printValue()
	mc.printValue()
}
