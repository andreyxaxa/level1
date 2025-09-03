package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("chan: goroutine stopped")
				return
			default:
				fmt.Println("goroutine working")
			}
		}
	}()

	// Дадим поработать 3 секунды
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * time.Duration(i))
	}
	// Закрываем канал
	close(stop)
	// Чтобы 'main' не завершился до созданной горутины
	time.Sleep(1 * time.Second)
}
