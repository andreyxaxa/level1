package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func Generator(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("generator stopped")
			return
		case ch <- rand.Int():
		}
	}
}

func Worker(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("context done: worker %d stopped\n", id)
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Printf("worker's %d channel closed\n", id)
				return
			}
			fmt.Printf("worker %d: %d\n", id, v)
		}
	}
}

func main() {
	ch := make(chan int)

	workersNum, err := strconv.Atoi(os.Args[1])
	if err != nil || workersNum < 1 {
		return
	}

	// контекст для отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go Worker(ctx, i+1, ch, &wg)
	}

	// Запуск генератора
	{
		wg.Add(1)
		go Generator(ctx, ch, &wg)
	}

	// Ожидание сигнала
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	fmt.Println("interrupt signal received")
	cancel()

	// Завершение всех горутин
	wg.Wait()
}
