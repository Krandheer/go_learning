package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("testing few things")
	websitelist := []string{
		"https://lco.dev", 
		"https://go.dev", 
		"https://google.com", 
	}
	for _, web := range websitelist {
		getStatusCode(web)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("error occured", err)
	} else {
		fmt.Printf("%d of the this %s website \n", res.StatusCode, endpoint)
	}
}