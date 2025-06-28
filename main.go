package main

import (
	"fmt"
	"sync"
)

var count int = 0
var mu sync.Mutex
var wg sync.WaitGroup

func increment() {
	mu.Lock()
	defer mu.Unlock()
	count += 1
	wg.Done()

}

func main() {
	for range 1000000 {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println(count)
}
