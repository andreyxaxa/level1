package main

import (
	"fmt"
)

// 1. Пишем в переданный канал все элементы массива
func generate(ch chan<- int, numbers []int) {
	go func() {
		for _, v := range numbers {
			ch <- v
		}
		close(ch)
	}()
}

// 2. В 'chWrite' будет записан каждый 'v*2' из 'chRead'
func multiply(chRead <-chan int, chWrite chan<- int) {
	go func() {
		for v := range chRead {
			chWrite <- v * 2
		}
		close(chWrite)
	}()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	generate(ch1, numbers)
	multiply(ch1, ch2)

	// Выводим данные. 'ch2' закрывается в 'multiply', поэтому цикл корректно завершится, как только будут прочитаны все данные.
	for v := range ch2 {
		fmt.Println(v)
	}
}
