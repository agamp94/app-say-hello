package main

import (
	"fmt"
	"sync"
)

func main() {
	var x = 0
	var mutex sync.RWMutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	fmt.Println(x)
}
