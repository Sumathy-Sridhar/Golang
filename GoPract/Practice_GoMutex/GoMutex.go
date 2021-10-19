package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 200; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("Final X value: ", x)

}

// without the use of syncmutex, the final value differes for each excution like 194,200, 196 because of race condition.
// with the help of syncmutex, the final value of x remains the same for all exceutions.
