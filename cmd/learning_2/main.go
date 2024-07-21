package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}
func main() {
// learning wait group and mutex 
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

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("error occured", err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d of the this %s website \n", res.StatusCode, endpoint)
	}
}