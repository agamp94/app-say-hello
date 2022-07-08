package main

import (
	"fmt"
	"sync"
	"time"
)

func runAsync(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func main() {
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		runAsync(&group)
	}

	group.Wait()
}
