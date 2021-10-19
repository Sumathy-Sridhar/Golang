package main

import (
	"fmt"
	"time"
)

func display(done chan bool) {
	fmt.Println("From Display method")
	time.Sleep(2 * time.Second)
	fmt.Println("From Display method after sleep")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("From Main method")
	go display(done)
	<-done
	fmt.Println("Data received from main")
}
