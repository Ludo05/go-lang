package main

import (
	"fmt"
	"sync"
	"time"
)

 var wg sync.WaitGroup
func main() {

	wg.Add(1)
	go DoSomething("Hello")
	wg.Add(1)
	go DoSomething("There")
	wg.Wait()

	fmt.Println("jhjh")
}

func DoSomething(str string) {
	for i := 0; i < 3; i++  {
		fmt.Println(str)
		time.Sleep(time.Second * 2)
	}
	wg.Done()
}
