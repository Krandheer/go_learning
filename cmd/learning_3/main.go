package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	defer wg.Done()

}

type User struct {
	username string
	id       int
}

func main() {
	/*wg.Add(2)
	go say("world") // create a new goroutine
	say("hello")    // current goroutine
	wg.Wait()*/
	u, name := User{"ravi", 1234}, "randheer"
	fmt.Println(u)
	fmt.Println(name)
	fmt.Printf("%v\n", u)
}
