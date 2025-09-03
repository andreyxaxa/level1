package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context: goroutine stopped")
				return
			default:
				fmt.Println("goroutine working")
			}
		}
	}(ctx)

	// Дадим поработать 3 секунды
	time.Sleep(3 * time.Second)
	// Отмена контекста
	cancel()
	// Чтобы 'main' не завершился до созданной горутины
	time.Sleep(1 * time.Second)
}
