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
			ch <- "some string"
			time.Sleep(700 * time.Millisecond)
		}
	}()

	for {
		select {
		// Читаем
		case v := <-ch:
			fmt.Println(v)
		case <-timeout:
			fmt.Println("timeout exceeded")
			return
		}
	}
}
