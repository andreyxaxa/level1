package main

import (
	"fmt"
	"time"
)

func main() {
	stop := false

	go func() {
		// Условие
		for !stop {
			fmt.Println("goroutine working")
		}
		fmt.Println("goroutine stopped")
	}()

	// Дадим пару секунд поработать
	time.Sleep(2 * time.Second)
	// Условие
	stop = true
	// Чтобы 'main' не завершился до созданной горутины
	time.Sleep(1 * time.Second)
}
