package main

import (
	"fmt"
	"time"
)

//Every time you will run the program it will give different outputs since the goroutines will be run concurrently
func execute(id int) {
	fmt.Printf("id: %d\n", id)
}

func main() {
	fmt.Println("Started")
	for i := 0; i < 10; i++ {
		go execute(i)
	}
	time.Sleep(time.Second * 4)
	fmt.Println("Finished")
}
