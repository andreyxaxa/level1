package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	t := time.NewTimer(duration)

	// небуф. канал - блокирующее чтение
	<-t.C
}

func main() {
	now := time.Now()

	sleep(3 * time.Second)

	fmt.Println(time.Since(now))
}
