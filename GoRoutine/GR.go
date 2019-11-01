package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 0
func main() {
	//This turns the function into a go routine (Spins off a green thread and runs this on the green thread).
	//This alone wont print Say Hello as it runs and finishes before it has a time to complete the func call
	//BEWARE OF THE RACE CONDITION
	runtime.GOMAXPROCS(100)
	for  i := 1; i <= 10 ;i++  {
	wg.Add(2)
		m.RLock()
		go showCount()
		m.Lock()
		go updateCount()
	}
	wg.Wait()
}

func showCount() {
	//Is a read lock
	fmt.Printf("#count: %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func updateCount() {
	//Is a write lock
	counter++
	m.Unlock()
	wg.Done()
}

