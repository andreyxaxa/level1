package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	// Добавляем к счетчику
	wg.Add(1)
	go func() {
		// Уменьшаем счетчик
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine working")
		}
	}()

	// Ждем заверешения горутин(ы)
	wg.Wait()
	fmt.Println("goroutine stopped")
}
