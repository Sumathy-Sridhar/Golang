package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	ch <- "Brown"
	ch <- "Cony"
	//ch <- "Sumathy" // deadlock
	fmt.Println(<-ch)
	time.Sleep(2 * time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch) //
}
