package main

import (
	"fmt"
	"time"
)

func main() {
	// 5 секунд
	N := 5
	timeout := time.After(time.Duration(N) * time.Second)

	ch := make(chan string)

	// Пишем
	go func() {
		for {
			select {
			// Закрываем канал и прекращаем писать, как только срабатывает таймаут
			case <-timeout:
				fmt.Println("timeout exceeded")
				close(ch)
				return
			default:
				ch <- "some string"
				time.Sleep(700 * time.Millisecond)
			}
		}
	}()

	// Читаем
	for v := range ch {
		fmt.Println(v)
	}
}
