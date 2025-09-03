package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println("goroutine working")
		// Принудительно завершаем горутину
		runtime.Goexit()
	}()

	// Чтобы 'main' не завершился до созданной горутины
	time.Sleep(3 * time.Second)
}
