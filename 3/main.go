package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func Worker(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)

	workersNum, err := strconv.Atoi(os.Args[1])
	if err != nil || workersNum < 1 {
		return
	}

	// запуск воркеров
	for i := 0; i < workersNum; i++ {
		go Worker(ch)
	}

	// запись
	for {
		ch <- rand.Int()
	}
}
