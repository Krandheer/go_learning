package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
	defer wg.Done()

}

func main() {
	wg.Add(2)
	go say("world") // create a new goroutine
	say("hello")    // current goroutine
	wg.Wait()
}
