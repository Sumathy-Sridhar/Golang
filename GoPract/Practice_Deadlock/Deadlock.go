package main

import "fmt"

func send(unichan chan<- int) {
	unichan <- 100
}

func closechannel(closech chan int) {
	for i := 1; i <= 5; i++ {
		closech <- i
	}
	close(closech) // closing the channel
}
func main() {
	// ch := make(chan int) // created a channel
	// ch <- 5              // written value 5 to the channel ch
	//Goroutine is sending data on a channel,
	// then it is expected that some other Goroutine should be receiving the data.
	//If this does not happen, then the program will panic at runtime with Deadlock.

	unichan := make(chan<- int)
	go send(unichan)
	fmt.Println(unichan)
	//	fmt.Println(<-unichan)  invalid operation: <-unichan (receive from send-only type chan<- int)

	// to convert the unidirectional channel to bidirectional channel
	bichan := make(chan int)
	go send(bichan)
	fmt.Println(<-bichan) // since it is converted as bidrectional channel now it can read the data that is written(100) to it.

	closech := make(chan int)
	go closechannel(closech)
	for {
		v, ok := <-closech
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
