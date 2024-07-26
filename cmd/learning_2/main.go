package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		// fmt.Println("error occured", err)
		log.Fatal("error in url: ", endpoint)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d of the this %s website \n", res.StatusCode, endpoint)
	}
}

func understandingMutex() {
	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func userInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("input your value: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("you entered %v \n", input)
	fmt.Printf("type of input is %T \n", input)
}

func main() {
	// understandingMutex()
	userInput()
}
