package main

import (
	"fmt"
	"math/rand"
	"time"
)

func giveMeResponse(channel chan string) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	channel <- "Wow"
}
func main() {

	counter := 0

	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go giveMeResponse(channel1)
	go giveMeResponse(channel2)

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
