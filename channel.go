package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getData(channel chan<- string) string {
	//time.Sleep(2 * time.Second)
	channel <- "ini dataaa " + strconv.Itoa(rand.Int())
	fmt.Println("Selesai kirim ke channel")
	return "ini dataaaa"
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go getData(channel)
	go getData(channel)
	go getData(channel)
	func() {
		fmt.Println("Wow", <-channel)
	}()

	fmt.Println("Fantastic")

	time.Sleep(5 * time.Second)

}
