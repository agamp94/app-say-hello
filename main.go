package main

import (
	"fmt"
	go_say_hello "github.com/agamp94/go-say-hello/v2"
	"time"
)

func runHelloWorld() {
	fmt.Println("Hello World!")
}

func main() {
	go runHelloWorld()
	fmt.Println("Ups")
	fmt.Println(go_say_hello.SayHello("Agam"))
	time.Sleep(1 * time.Second)
}
